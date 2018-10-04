package com.ngtesting.platform.model;

import java.util.*;

public class TstTask extends BaseModel {

    private static final long serialVersionUID = 3655131645148750323L;
    private String name;
    private Integer estimate;

    private TaskStatus status = TaskStatus.not_start;

    protected Date startTime;

    protected Date endTime;

    private String descr;

    private Integer projectId;
    private String projectName;
    private Integer caseProjectId;
    private String caseProjectName;

    private Integer planId;

    private Integer userId;
    private String userName;
    private Integer envId;
    private String envName;

    public TstTask() {

    }

    private List<TstUser> assignees = new LinkedList<>();

    private Map<String, Integer> countMap = new HashMap<String, Integer>(){{
        put("total", 0);
        put("pass", 0);
        put("fail", 0);
        put("block", 0);
        put("untest", 0);
    }};
    private Map<String, Integer> widthMap = new HashMap<String, Integer>(){{
        put("total", 0);
        put("pass", 0);
        put("fail", 0);
        put("block", 0);
        put("untest", 0);
    }};

    public static enum TaskStatus {
        not_start("not_start"),
        in_progress("in_progress"),
        end("end");

        TaskStatus(String val) {
            this.val = val;
        }

        private String val;
        public String toString() {
            return val;
        }
    }

    public String getProjectName() {
        return projectName;
    }

    public void setProjectName(String projectName) {
        this.projectName = projectName;
    }

    public Integer getCaseProjectId() {
        return caseProjectId;
    }

    public void setCaseProjectId(Integer caseProjectId) {
        this.caseProjectId = caseProjectId;
    }

    public String getCaseProjectName() {
        return caseProjectName;
    }

    public void setCaseProjectName(String caseProjectName) {
        this.caseProjectName = caseProjectName;
    }

    public Integer getEnvId() {
        return envId;
    }

    public void setEnvId(Integer envId) {
        this.envId = envId;
    }

    public String getEnvName() {
        return envName;
    }

    public void setEnvName(String envName) {
        this.envName = envName;
    }

    public List<TstUser> getAssignees() {
        return assignees;
    }
    public void setAssignees(List<TstUser> assignees) {
        this.assignees = assignees;
    }

    public Integer getUserId() {
        return userId;
    }

    public void setUserId(Integer userId) {
        this.userId = userId;
    }

    public String getUserName() {
        return userName;
    }

    public void setUserName(String userName) {
        this.userName = userName;
    }

    public Map<String, Integer> getWidthMap() {
        return widthMap;
    }

    public void setWidthMap(Map<String, Integer> widthMap) {
        this.widthMap = widthMap;
    }

    public Map<String, Integer> getCountMap() {
        return countMap;
    }

    public void setCountMap(Map<String, Integer> countMap) {
        this.countMap = countMap;
    }

    public Integer getPlanId() {
        return planId;
    }

    public void setPlanId(Integer planId) {
        this.planId = planId;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public Integer getEstimate() {
        return estimate;
    }

    public void setEstimate(Integer estimate) {
        this.estimate = estimate;
    }

    public TaskStatus getStatus() {
        return status;
    }

    public void setStatus(TaskStatus status) {
        this.status = status;
    }

    public Date getStartTime() {
        return startTime;
    }

    public void setStartTime(Date startTime) {
        this.startTime = startTime;
    }

    public Date getEndTime() {
        return endTime;
    }

    public void setEndTime(Date endTime) {
        this.endTime = endTime;
    }

    public String getDescr() {
        return descr;
    }

    public void setDescr(String descr) {
        this.descr = descr;
    }

    public Integer getProjectId() {
        return projectId;
    }

    public void setProjectId(Integer projectId) {
        this.projectId = projectId;
    }

}
