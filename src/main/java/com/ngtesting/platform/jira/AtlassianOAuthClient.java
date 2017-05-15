package com.ngtesting.platform.jira;

import com.google.common.collect.ImmutableList;
import net.oauth.OAuth;
import net.oauth.OAuthAccessor;
import net.oauth.OAuthConsumer;
import net.oauth.OAuthMessage;
import net.oauth.OAuthServiceProvider;
import net.oauth.client.OAuthClient;
import net.oauth.client.httpclient4.HttpClient4;
import net.oauth.signature.RSA_SHA1;

import java.util.ArrayList;
import java.util.Collection;
import java.util.Collections;
import java.util.List;
import java.util.Map;

import static net.oauth.OAuth.OAUTH_VERIFIER;

/**
 * @since v1.0
 */ 
public class AtlassianOAuthClient
{
    protected static final String SERVLET_BASE_URL = "/plugins/servlet";

    private final String consumerKey;
    private final String privateKey;
    private final String baseUrl;
    private final String callback;
    private OAuthAccessor accessor;


    public AtlassianOAuthClient(String consumerKey, String privateKey, String baseUrl, String callback)
    {
        this.consumerKey = consumerKey;
        this.privateKey = privateKey;
        this.baseUrl = baseUrl;
        this.callback = callback;
    }

    public TokenSecretVerifierHolder getRequestToken()
    {
        try
        {
            OAuthAccessor accessor = getAccessor();
            OAuthClient oAuthClient = new OAuthClient(new HttpClient4());
            List<OAuth.Parameter> callBack;
            if (callback == null || "".equals(callback))
            {
                callBack = Collections.<OAuth.Parameter>emptyList();
            }
            else
            {
                callBack = ImmutableList.of(new OAuth.Parameter(OAuth.OAUTH_CALLBACK, callback));
            }

            OAuthMessage message = oAuthClient.getRequestTokenResponse(accessor, "POST", callBack);
            TokenSecretVerifierHolder tokenSecretVerifier = new TokenSecretVerifierHolder();
            tokenSecretVerifier.token = accessor.requestToken;
            tokenSecretVerifier.secret = accessor.tokenSecret;
            tokenSecretVerifier.verifier = message.getParameter(OAUTH_VERIFIER);
            return tokenSecretVerifier;
        }
        catch (Exception e) {
            throw new RuntimeException("Failed to obtain request token", e);
        }
    }

    public String swapRequestTokenForAccessToken(String requestToken, String tokenSecret, String oauthVerifier)
    {
        try
        {
            OAuthAccessor accessor = getAccessor();
            OAuthClient client = new OAuthClient(new HttpClient4());
            accessor.requestToken = requestToken;
            accessor.tokenSecret = tokenSecret;
            OAuthMessage message = client.getAccessToken(accessor, "POST",
                    ImmutableList.of(new OAuth.Parameter(OAuth.OAUTH_VERIFIER, oauthVerifier)));
            return message.getToken();
        }
        catch (Exception e)
        {
            throw new RuntimeException("Failed to swap request token with access token", e);
        }
    }

    public String getAuthenticatedRequest(String url, String accessToken, List<OAuth.Parameter> params) {
        try
        {
            OAuthAccessor accessor = getAccessor();
            OAuthClient client = new OAuthClient(new HttpClient4());
            accessor.accessToken = accessToken;
            OAuthMessage response = client.invoke(accessor, url, params);
            return response.readBodyAsString();
        }
        catch (Exception e)
        {
            throw new RuntimeException("Failed to make an authenticated request.", e);
        }
    }

    private final OAuthAccessor getAccessor()
    {
        if (accessor == null)
        {
            OAuthServiceProvider serviceProvider = new OAuthServiceProvider(getRequestTokenUrl(), getAuthorizeUrl(), getAccessTokenUrl());
            OAuthConsumer consumer = new OAuthConsumer(callback, consumerKey, null, serviceProvider);
            consumer.setProperty(RSA_SHA1.PRIVATE_KEY, privateKey);
            consumer.setProperty(OAuth.OAUTH_SIGNATURE_METHOD, OAuth.RSA_SHA1);
            accessor = new OAuthAccessor(consumer);
        }
        return accessor;
    }

    private String getAccessTokenUrl()
    {
        return baseUrl + SERVLET_BASE_URL + "/oauth/access-token";
    }

    private String getRequestTokenUrl()
    {
        return  baseUrl + SERVLET_BASE_URL + "/oauth/request-token";
    }

    public String getAuthorizeUrlForToken(String token)
    {
        return getAuthorizeUrl() + "?oauth_token=" + token;
    }

    private String getAuthorizeUrl() {return baseUrl + SERVLET_BASE_URL + "/oauth/authorize";}
}
