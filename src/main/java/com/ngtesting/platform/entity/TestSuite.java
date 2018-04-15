package com.ngtesting.platform.entity;

import javax.persistence.*;
import java.util.LinkedList;
import java.util.List;

@Entity
@Table(name = "tst_suite")
public class TestSuite extends BaseEntity {
    private static final long serialVersionUID = -8163006531081824433L;

    private String name;
    private Integer estimate;

    @Column(name = "descr", length = 1000)
    private String descr;

    private Integer ordr;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "case_project_id", insertable = false, updatable = false)
    private TestProject caseProject;
    @Column(name = "case_project_id")
    private Long caseProjectId;

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

    @OneToMany(mappedBy="suite", cascade={CascadeType.ALL}, fetch=FetchType.LAZY)
    private List<TestCaseInSuite> testcases = new LinkedList<>();

    public TestProject getCaseProject() {
        return caseProject;
    }

    public void setCaseProject(TestProject caseProject) {
        this.caseProject = caseProject;
    }

    public Long getCaseProjectId() {
        return caseProjectId;
    }

    public void setCaseProjectId(Long caseProjectId) {
        this.caseProjectId = caseProjectId;
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

    public String getDescr() {
        return descr;
    }

    public void setDescr(String descr) {
        this.descr = descr;
    }

    public Integer getOrdr() {
        return ordr;
    }

    public void setOrdr(Integer ordr) {
        this.ordr = ordr;
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

    public List<TestCaseInSuite> getTestcases() {
        return testcases;
    }

    public void setTestcases(List<TestCaseInSuite> testcases) {
        this.testcases = testcases;
    }
}
