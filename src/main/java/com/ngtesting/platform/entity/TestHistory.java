package com.ngtesting.platform.entity;

import javax.persistence.*;

@Entity
@Table(name = "tst_history")
public class TestHistory extends BaseEntity {
    private static final long serialVersionUID = -6608023158199904153L;

    private String title;
    @Column(name = "descr", length = 1000)
    private String descr;
    private String uri;

    private Long entityId;
    @Enumerated(EnumType.STRING)
    private TargetType entityType;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "project_id", insertable = false, updatable = false)
    private TestProject project;

    @Column(name = "project_id")
    private Long projectId;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "user_id", insertable = false, updatable = false)
    private TestUser user;

    @Column(name = "user_id")
    private Long userId;

    public enum TargetType {
        project("project", "项目"),
        project_member("project_member", "项目成员"),

        plan("plan", "计划"),
        suite("suite", "测试集"),
        run("run", "测试任务");

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

    public TestProject getProject() {
        return project;
    }

    public void setProject(TestProject project) {
        this.project = project;
    }

    public Long getProjectId() {
        return projectId;
    }

    public void setProjectId(Long projectId) {
        this.projectId = projectId;
    }

    public TargetType getEntityType() {
        return entityType;
    }

    public void setEntityType(TargetType entityType) {
        this.entityType = entityType;
    }

    public Long getEntityId() {
        return entityId;
    }

    public void setEntityId(Long entityId) {
        this.entityId = entityId;
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

    public String getTitle() {
        return title;
    }

    public void setTitle(String title) {
        this.title = title;
    }

    public TestUser getUser() {
        return user;
    }

    public void setUser(TestUser user) {
        this.user = user;
    }

    public Long getUserId() {
        return userId;
    }

    public void setUserId(Long userId) {
        this.userId = userId;
    }
}
