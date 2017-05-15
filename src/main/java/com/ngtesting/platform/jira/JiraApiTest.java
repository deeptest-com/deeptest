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
    	
    	String requestUrl = "http://localhost:9000/rest/api/2/search";
    	
        AtlassianOAuthClient jiraoAuthClient = new AtlassianOAuthClient(
        		JIRAOAuthClient.CONSUMER_KEY, JIRAOAuthClient.CONSUMER_PRIVATE_KEY, 
        		baseUrl, null);
//        
//        TokenSecretVerifierHolder requestToken = jiraoAuthClient.getRequestToken();
//        String authorizeUrl = jiraoAuthClient.getAuthorizeUrlForToken(requestToken.token);
//        System.out.println("Token is " + requestToken.token);
//        System.out.println("Token secret is " + requestToken.secret);
//        System.out.println("Retrieved request token. go to " + authorizeUrl);
//        
//        String accessToken = jiraoAuthClient.swapRequestTokenForAccessToken(requestToken.token, requestToken.secret, null);
//        System.out.println("Access token is : " + accessToken);

        String accessToken = "eOs8g3eL7ZUf6FljFGEf2hnqezetJfvi";
        
        List<OAuth.Parameter> params = ImmutableList.of(new OAuth.Parameter("jql", 
        		"project=TEST AND issuetype = Bug "
//        		+ " AND status in (\"In Progress\", \"To Do\")"
        		+ " ORDER BY created DESC"));
        
        String responseAsString1 = jiraoAuthClient.getAuthenticatedRequest(requestUrl, accessToken, params);
        System.out.println(responseAsString1);
        
    }

}
