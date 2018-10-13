package com.ngtesting.platform.config;

import java.util.HashMap;
import java.util.Map;

public class ConstantIssue {

    public static Map<String, Boolean> IssueFilters = new HashMap() {{
        put("project", true);
        put("type", true);
        put("status", true);
        put("priority", true);
        put("assignee", true);

        put("creator", false);
        put("reporter", false);
    }};

}
