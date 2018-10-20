package com.ngtesting.platform.config;

import java.util.Arrays;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;

public class ConstantIssue {

    public static Map<String, String> IssueColumns = new LinkedHashMap() {{
        put("title", "标题");
        put("typeId", "类型");
        put("statusId", "状态");
        put("priorityId", "优先级");
        put("assigneeId", "经办人");

        put("creatorId", "创家人");
        put("reporterId", "报告人");

        put("verId", "版本");
        put("envId", "环境");
        put("resolution", "解决结果");
        put("dueTime", "到期时间");
        put("resolveTime", "解决时间");
        put("projectId", "项目");
    }};

    public static Map<String, String> IssueFilters = new LinkedHashMap() {{
        put("projectId", "项目");
        put("typeId", "类型");
        put("statusId", "状态");
        put("priorityId", "优先级");
        put("assigneeId", "经办人");

        put("creatorId", "创建人");
        put("reporterId", "报告人");

        put("verId", "版本");
        put("envId", "环境");
        put("resolutionId", "解决结果");
        put("dueTime", "到期时间");
        put("resolveTime", "解决时间");
        put("comments", "注释");
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
