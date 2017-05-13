package com.ngtesting.platform.jira;

import com.google.common.collect.Lists;

import java.util.ArrayList;

/**
 * @since v1.0
 */
public class JIRAOAuthClient
{
    private static final String CALLBACK_URI = "";
    protected static final String CONSUMER_KEY = "hardcoded-consumer";
    protected static final String CONSUMER_PRIVATE_KEY = "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDFkPMZQaTqsSXI+bSI65rSVaDzic6WFA3WCZMVMi7lYXJAUdkXo4DgdfvEBO21Bno3bXIoxqS411G8S53I39yhSp7z2vcB76uQQifi0LEaklZfbTnFUXcKCyfwgKPp0tQVA+JZei6hnscbSw8qEItdc69ReZ6SK+3LHhvFUUP1nLhJDsgdPHRXSllgZzqvWAXQupGYZVANpBJuK+KAfiaVXCgA71N9xx/5XTSFi5K+e1T4HVnKAzDasAUt7Mmad+1PE+56Gpa73FLk1Ww+xaAEvss6LehjyWHM5iNswoNYzrNS2k6ZYkDnZxUlbrPDELETbz/n3YgBHGUlyrXi2PBjAgMBAAECggEAAtMctqq6meRofuQbEa4Uq5cv0uuQeZLV086VPMNX6k2nXYYODYl36T2mmNndMC5khvBYpn6Ykk/5yjBmlB2nQOMZPLFPwMZVdJ2Nhm+naJLZC0o7fje49PrN2mFsdoZeI+LHVLIrgoILpLdBAz/zTiW+RvLvMnXQU4wdp4eO6i8J/Jwh0AY8rWsAGkk1mdZDwklPZZiwR3z+DDsDwPxFs8z6cE5rWJd2c/fhAQrHwOXyrQPsGyLHTOqS3BkjtEZrKRUlfdgV76VlThwrE5pAWuO0GPyfK/XCklwcNS1a5XxCOq3uUogWRhCsqUX6pYfAVS6xzX56MGDndQVlp7U5uQKBgQDyTDwhsNTWlmr++FyYrc6liSF9NEMBNDubrfLJH1kaOp590bE8fu3BG0UlkVcueUr05e33Kx1DMSFW72lR4dht1jruWsbFp6LlT3SUtyW2kcSet3fC8gySs2r6NncsZ2XFPoxTkalKpQ1atGoBe3XIKeT8RDZtgoLztQy7/7yANQKBgQDQvSHEKS5SttoFFf4YkUh2QmNX5m7XaDlTLB/3xjnlz8NWOweK1aVysb4t2Tct/SR4ZZ/qZDBlaaj4X9h9nlxxIMoXEyX6Ilc4tyCWBXxn6HFMSa/Rrq662Vzz228cPvW2XGOQWdj7IqwKO9cXgJkI5W84YtMtYrTPLDSjhfpxNwKBgGVCoPq/iSOpN0wZhbE1KiCaP8mwlrQhHSxBtS6CkF1a1DPm97g9n6VNfUdnB1Vf0YipsxrSBOe416MaaRyUUzwMBRLqExo1pelJnIIuTG+RWeeu6zkoqUKCAxpQuttu1uRo8IJYZLTSZ9NZhNfbveyKPa2D4G9B1PJ+3rSO+ztlAoGAZNRHQEMILkpHLBfAgsuC7iUJacdUmVauAiAZXQ1yoDDo0Xl4HjcvUSTMkccQIXXbLREh2w4EVqhgR4G8yIk7bCYDmHvWZ2o5KZtD8VO7EVI1kD0z4Zx4qKcggGbp2AINnMYqDetopX7NDbB0KNUklyiEvf72tUCtyDk5QBgSrqcCgYEAnlg3ByRd/qTFz/darZi9ehT68Cq0CS7/B9YvfnF7YKTAv6J2Hd/i9jGKcc27x6IMi0vf7zrqCyTMq56omiLdu941oWfsOnwffWRBInvrUWTj6yGHOYUtg2z4xESUoFYDeWwe/vX6TugL3oXSX3Sy3KWGlJhn/OmsN2fgajHRip0=";
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
            String responseAsString = jiraoAuthClient.getAuthenticatedRequest(url, accessToken);
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
