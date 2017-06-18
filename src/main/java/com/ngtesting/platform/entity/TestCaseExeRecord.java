package com.ngtesting.platform.entity;

import com.ngtesting.platform.util.Constant.TreeNodeType;

import javax.persistence.*;

@Entity
@Table(name = "tst_case_exe_record")
public class TestCaseExeRecord extends BaseEntity {

    private String title;
    private Integer priority;
    private Integer estimate;

    @Column(name = "descr", length = 1000)
    private String descr;

    @Enumerated(EnumType.STRING)
    private TreeNodeType type;
    private String path;
    @Transient
    private Integer level;
    private Integer orderInParent;

    private String objective;

    @OneToOne(cascade = CascadeType.ALL)
    @JoinColumn(name = "case_extend_id")
    private TestCaseExtend caseExtend;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "parent_id", insertable = false, updatable = false)
    private TestCase parent;

    @Column(name = "parent_id")
    private Long parentId;

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

    private String statusCode;
    private String statusName;

    @Column(name = "run_id")
    private Long runId;

    @Column(name = "case_id")
    private Long caseId;

    public String getStatusCode() {
        return statusCode;
    }

    public void setStatusCode(String statusCode) {
        this.statusCode = statusCode;
    }

    public String getStatusName() {
        return statusName;
    }

    public void setStatusName(String statusName) {
        this.statusName = statusName;
    }

    public Long getRunId() {
        return runId;
    }

    public void setRunId(Long runId) {
        this.runId = runId;
    }

    public Long getCaseId() {
        return caseId;
    }

    public void setCaseId(Long caseId) {
        this.caseId = caseId;
    }

    public String getTitle() {
        return title;
    }

    public void setTitle(String title) {
        this.title = title;
    }

    public Integer getPriority() {
        return priority;
    }

    public void setPriority(Integer priority) {
        this.priority = priority;
    }

    public Integer getEstimate() {
        return estimate;
    }

    public void setEstimate(Integer estimate) {
        this.estimate = estimate;
    }

    public String getDescr() {
        return descr;
    }

    public void setDescr(String descr) {
        this.descr = descr;
    }

    public TreeNodeType getType() {
        return type;
    }

    public void setType(TreeNodeType type) {
        this.type = type;
    }

    public String getPath() {
        return path;
    }

    public void setPath(String path) {
        this.path = path;
    }

    public Integer getLevel() {
        return level;
    }

    public void setLevel(Integer level) {
        this.level = level;
    }

    public Integer getOrderInParent() {
        return orderInParent;
    }

    public void setOrderInParent(Integer orderInParent) {
        this.orderInParent = orderInParent;
    }

    public TestCaseExtend getCaseExtend() {
        return caseExtend;
    }

    public void setCaseExtend(TestCaseExtend caseExtend) {
        this.caseExtend = caseExtend;
    }

    public TestCase getParent() {
        return parent;
    }

    public void setParent(TestCase parent) {
        this.parent = parent;
    }

    public Long getParentId() {
        return parentId;
    }

    public void setParentId(Long parentId) {
        this.parentId = parentId;
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

    public String getObjective() {
        return objective;
    }

    public void setObjective(String objective) {
        this.objective = objective;
    }
}
