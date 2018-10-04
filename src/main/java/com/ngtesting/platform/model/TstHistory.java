package com.ngtesting.platform.model;

public class TstHistory extends BaseModel {
    private static final long serialVersionUID = -6418363700406245211L;

    private String title;
    private String descr;
    private String uri;

    private Integer projectId;
    private Integer entityId;
    private TargetType entityType;

    private Integer userId;

    public enum TargetType {
        project("project", "项目"),
        project_member("project_member", "项目成员"),

        plan("plan", "计划"),
        suite("suite", "测试集"),
        task("task", "测试任务");

        TargetType(String code, String name) {
            this.code = code;
            this.name = name;
        }

        public String code;
        public String name;
        public String toString() {
            return code;
        }
    }

    public String getTitle() {
        return title;
    }

    public void setTitle(String title) {
        this.title = title;
    }

    public String getDescr() {
        return descr;
    }

    public void setDescr(String descr) {
        this.descr = descr;
    }

    public String getUri() {
        return uri;
    }

    public void setUri(String uri) {
        this.uri = uri;
    }

    public Integer getUserId() {
        return userId;
    }

    public void setUserId(Integer userId) {
        this.userId = userId;
    }

    public Integer getEntityId() {
        return entityId;
    }

    public void setEntityId(Integer entityId) {
        this.entityId = entityId;
    }

    public Integer getProjectId() {
        return projectId;
    }

    public void setProjectId(Integer projectId) {
        this.projectId = projectId;
    }

    public TargetType getEntityType() {
        return entityType;
    }

    public void setEntityType(TargetType entityType) {
        this.entityType = entityType;
    }
}
