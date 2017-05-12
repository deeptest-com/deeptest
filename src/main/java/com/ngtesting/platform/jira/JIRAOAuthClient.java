package com.ngtesting.platform.jira;

import com.google.common.collect.Lists;

import java.util.ArrayList;

/**
 * @since v1.0
 */
public class JIRAOAuthClient
{
    private static final String CALLBACK_URI = "http://consumer/callback";
    protected static final String CONSUMER_KEY = "hardcoded-consumer";
    protected static final String CONSUMER_PRIVATE_KEY = "MIICeAIBADANBgkqhkiG9w0BAQEFAASCAmIwggJeAgEAAoGBALmL1I5DFFplVnoaMjINAlzz3xLXVpdfhbnONDZjPUsXLSyq7bHaI6unIZGr7u3NgqXo5mPdezhEePjCKm7NWZgxE+eBG48iWu/E2Tc4eBCZtsBa4oTnD49RQ0txAuGxApO+/AG1lSNIWka5XxCYzcD/fJraWPTwpMdnrO6Z3QHJAgMBAAECgYB6WJepztWG3bduAAQFOMrMHAqF0/RHEDePU0bebgWmk/u3rmXZlmta7nOWKHb92ztAxxfT1eFAMvZJoU51jXZ7YNohgBw97bgtudJlItAO7c5yuv7iYIleWDcw72uWZmZqWjL1jxmeM2kwT57hB3T8vlZkw3MnTOnz+RfKvvcMAQJBAPSd116jXRCB2dpMAZJkI0rC3IGMqWEbROM3wR6ho7Asuy2YqB/QMIR6csyAayeI59BHHvFBcT7qdhpfCadTHUECQQDCLkYZ1XaBCnu23q7DiWnU8g/wT5sAnkTaFwCb1Z1MAZKQ9CtswaKPUfYdJLAXgiAHFbaGB/ckEMd0vGDEOdqJAkEAgIN4hejv2N9PlAeAf+eKPxnW/VzoE/Neor1FAZHMTJ+DizX7hhM7mi42p8gEA9ZCa8Mht4A5PeOyPrKVgMoCwQJBAJvLDplJh9OOeqXE0gi2JkAgmiMfa2g6k1k7HVTqNVK27EX6cSDH1soQY1sMhW/HCjVE+XVdzK6V+8EwJp8pH6kCQQDpXdIUAnvri6rGGBBNGUh1/v9IRc6fBtny2Tz+JTPNGzn0ZSRUuSFPHSgu6YQAUzhKNxZYpQSMJ/8Nyv0jtIvs";
    public enum Command
    {
        REQUEST_TOKEN("requestToken"),
        ACCESS_TOKEN("accessToken"), REQUEST("request");

        private String name;

        Command(final String name)
        {
            this.name = name;
        }

        public String getName()
        {
            return name;
        }
    }

    public static void main(String[] args)
    {
        ArrayList<String> arguments = Lists.newArrayList(args);
        if (arguments.isEmpty())
        {
            throw new IllegalArgumentException("No command specified. Use one of " + getCommandNames() );
        }
        String action = arguments.get(0);
        if (Command.REQUEST_TOKEN.getName().equals(action))
        {
            String baseUrl = arguments.get(1);
            String callBack = "oob";
            if (arguments.size() == 3)
            {
                callBack = arguments.get(2);
            }
            AtlassianOAuthClient jiraoAuthClient = new AtlassianOAuthClient(CONSUMER_KEY, CONSUMER_PRIVATE_KEY, baseUrl, callBack);
            //STEP 1: Get request token
            TokenSecretVerifierHolder requestToken = jiraoAuthClient.getRequestToken();
            String authorizeUrl = jiraoAuthClient.getAuthorizeUrlForToken(requestToken.token);
            System.out.println("Token is " + requestToken.token);
            System.out.println("Token secret is " + requestToken.secret);
            System.out.println("Retrieved request token. go to " + authorizeUrl);
        }
        else if (Command.ACCESS_TOKEN.getName().equals(action))
        {
            String baseUrl = arguments.get(1);
            AtlassianOAuthClient jiraoAuthClient = new AtlassianOAuthClient(CONSUMER_KEY, CONSUMER_PRIVATE_KEY, baseUrl, CALLBACK_URI);
            String requestToken = arguments.get(2);
            String tokenSecret = arguments.get(3);
            String verifier = arguments.get(4);
            String accessToken = jiraoAuthClient.swapRequestTokenForAccessToken(requestToken, tokenSecret, verifier);
            System.out.println("Access token is : " + accessToken);
        }
        else if (Command.REQUEST.getName().equals(action))
        {
            AtlassianOAuthClient jiraoAuthClient = new AtlassianOAuthClient(CONSUMER_KEY, CONSUMER_PRIVATE_KEY, null, CALLBACK_URI);
            String accessToken = arguments.get(1);
            String url = arguments.get(2);
            String responseAsString = jiraoAuthClient.makeAuthenticatedRequest(url, accessToken);
            System.out.println("RESPONSE IS" + responseAsString);
        }
        else
        {
            System.out.println("Command " + action + " not supported. Only " + getCommandNames() + " are supported.");
        }
    }

    private static String getCommandNames()
    {
        String names = "";
        for (Command value : Command.values())
        {
            names += value.getName() + " ";
        }
        return names;
    }
}
