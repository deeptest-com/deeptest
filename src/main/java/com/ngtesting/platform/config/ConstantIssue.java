package com.ngtesting.platform.config;

public class ConstantIssue {
    public static enum IssueOpt {
        create("create"),
        edit("edit"),
        view("view");

        IssueOpt(String textVal) {
            this.textVal = textVal;
        }

        private String textVal;

        public String toString() {
            return textVal;
        }
    }

//    public static List<String> OperatorsForSelect =
//            Arrays.asList("equal", "not_equal", "in", "not_in", "is_null", "is_not_null");
//
//    public static List<String> OperatorsForText =
//            Arrays.asList("equal", "not_equal", "in", "not_in", "begins_with", "not_begins_with",
//                    "contains", "not_contains", "ends_with", "not_ends_with",
//                    "is_empty", "is_not_empty", "is_null", "is_not_null");
//
//    public static List<String> OperatorsForDate =
//            Arrays.asList("equal", "not_equal", "in", "not_in",
//                    "less", "less_or_equal", "equal", "greater", "greater_or_equal",
//                    "between", "not_between", "is_null", "is_not_null");

    public static enum TextFormat {
        plain_text("plain_text");
//        rich_text("rich_text");

        private TextFormat(String textVal) {
            this.textVal = textVal;
        }

        private String textVal;

        public String toString() {
            return textVal;
        }
    }
}
