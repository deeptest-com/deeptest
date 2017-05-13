package com.ngtesting.platform.jira;

import com.google.common.collect.ImmutableList;
import com.google.common.collect.Lists;

import net.oauth.OAuth;

import java.util.ArrayList;
import java.util.List;
import java.util.Map;

public class JiraApiTest {
	
    public static void main(String[] args) {
    	String baseUrl = "http://localhost:9000";
//    	String requestUrl = "http://localhost:9000/rest/api/2/issue/NGTES-7";
    	
    	String requestUrl = "http://localhost:9000/rest/api/2/search?jql=assignee=aaron";
    	
        AtlassianOAuthClient jiraoAuthClient = new AtlassianOAuthClient(
        		JIRAOAuthClient.CONSUMER_KEY, JIRAOAuthClient.CONSUMER_PRIVATE_KEY, 
        		baseUrl, null);
        
//        TokenSecretVerifierHolder requestToken = jiraoAuthClient.getRequestToken();
//        String authorizeUrl = jiraoAuthClient.getAuthorizeUrlForToken(requestToken.token);
//        System.out.println("Token is " + requestToken.token);
//        System.out.println("Token secret is " + requestToken.secret);
//        System.out.println("Retrieved request token. go to " + authorizeUrl);
//        
//        String accessToken = jiraoAuthClient.swapRequestTokenForAccessToken(requestToken.token, requestToken.secret, null);
//        System.out.println("Access token is : " + accessToken);

        String accessToken = "4OklL7x2niOEg06DX0YMSSV0UrtrNM6W";
        
        List<OAuth.Parameter> params = new ArrayList<OAuth.Parameter>(1);
//        params.add(new OAuth.Parameter("jql", "assignee=tester"));
        String responseAsString = jiraoAuthClient.getAuthenticatedRequest(requestUrl, accessToken, params);
        System.out.println("RESPONSE IS" + responseAsString);
        
//        List<OAuth.Parameter> params = ImmutableList.of(new OAuth.Parameter("aa", "1"));
//        responseAsString = jiraoAuthClient.postAuthenticatedRequest(requestUrl, accessToken, null);
//        System.out.println("RESPONSE IS" + responseAsString);
        
    }

}
