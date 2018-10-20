package com.ngtesting.platform.config;

import java.util.Arrays;
import java.util.List;

public class ConstantIssue {

    public static String[][] IssueColumns =
            {
                    {"title", "标题", "string"},
                    {"typeId", "类型", "string"},
                    {"statusId", "状态", "string"},
                    {"priorityId", "优先级", "string"},
                    {"assigneeId", "经办人", "string"},

                    {"creatorId", "创家人", "string"},
                    {"reporterId", "报告人", "string"},

                    {"verId", "版本", "string"},
                    {"envId", "环境", "string"},
                    {"resolution", "解决结果", "string"},
                    {"dueTime", "到期时间", "date"},
                    {"resolveTime", "解决时间", "date"},
                    {"projectId", "项目", "string"}
            };

    public static String[][] IssueFilters =
            {
                    {"projectId", "项目"},
                    {"typeId", "类型"},
                    {"statusId", "状态"},
                    {"priorityId", "优先级"},
                    {"assigneeId", "经办人"},

                    {"creatorId", "创建人"},
                    {"reporterId", "报告人"},

                    {"verId", "版本"},
                    {"envId", "环境"},
                    {"resolutionId", "解决结果"},
                    {"dueTime", "到期时间"},
                    {"resolveTime", "解决时间"},
                    {"comments", "注释"}
            };

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
