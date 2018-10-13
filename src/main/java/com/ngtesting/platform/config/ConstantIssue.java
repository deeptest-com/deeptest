package com.ngtesting.platform.config;

import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class ConstantIssue {

    public static Map<String, Boolean> Filters = new HashMap() {{
        put("project", true);
        put("type", true);
        put("status", true);
        put("priority", true);
        put("assignee", true);

        put("creator", false);
        put("reporter", false);
    }};

    public static List<String> Operators =
            Arrays.asList("equal", "not_equal", "in", "not_in", "is_null", "is_not_null");
}
