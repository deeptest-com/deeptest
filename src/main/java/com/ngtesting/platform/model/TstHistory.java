package com.ngtesting.platform.model;

public class TstHistory extends BaseModel {

    private static final long serialVersionUID = -6418363700406245211L;

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

    private String title;
    private String descr;
    private String uri;

    private Long entityId;
    private String type;

    private Long userId;

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

    public Long getUserId() {
        return userId;
    }

    public void setUserId(Long userId) {
        this.userId = userId;
    }

    public Long getEntityId() {
        return entityId;
    }

    public void setEntityId(Long entityId) {
        this.entityId = entityId;
    }

    public String getType() {
        return type;
    }

    public void setType(String type) {
        this.type = type;
    }
}
