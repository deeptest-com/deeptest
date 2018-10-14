package com.ngtesting.platform.config;

import java.util.Arrays;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;

public class ConstantIssue {

    public static Map<String, Boolean> IssueFilters = new LinkedHashMap() {{
        put("project", true);
        put("type", true);
        put("status", true);
        put("priority", true);
        put("assignee", true);

        put("creator", false);
        put("reporter", false);

        put("ver", false);
        put("env", false);
        put("resolution", false);
        put("dueTime", false);
        put("resolveTime", false);
        put("comments", false);
    }};

    public static enum IssueFilterInput {
        string("string"),
        select("select"),
        radio("radio"),
        date("date");

        private IssueFilterInput(String textVal) {
            this.textVal = textVal;
        }

        private String textVal;
        public String toString() {
            return textVal;
        }
    }

    public static enum IssueFilterType {
        integer("integer"),
        doubl("doubl"),
        string("string"),
        date("date");

        private IssueFilterType(String textVal) {
            this.textVal = textVal;
        }

        private String textVal;
        public String toString() {
            return textVal;
        }
    }

    public static List<String> OperatorsForSelect =
            Arrays.asList("equal", "not_equal", "in", "not_in", "is_null", "is_not_null");

    public static List<String> OperatorsForString =
            Arrays.asList("equal", "not_equal", "in", "not_in", "begins_with", "not_begins_with",
                    "contains", "not_contains", "ends_with", "not_ends_with",
                    "is_empty", "is_not_empty", "is_null", "is_not_null");

    public static List<String> OperatorsForDate =
            Arrays.asList("equal", "not_equal", "in", "not_in",
                    "less", "less_or_equal", "equal", "greater", "greater_or_equal",
                    "between", "not_between", "is_null", "is_not_null");
}
