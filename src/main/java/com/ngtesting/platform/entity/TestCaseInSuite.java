package com.ngtesting.platform.entity;

import javax.persistence.*;

@Entity
@Table(name = "tst_case_in_suite")
public class TestCaseInSuite extends BaseEntity {
    private static final long serialVersionUID = 6886190863032525770L;

    private Long pId;
    private Boolean isLeaf;

    public TestCaseInSuite() {
        super();
    }
    public TestCaseInSuite(Long projectId, Long suiteId, Long caseId, Long pid, Boolean isLeaf) {
        super();
        this.projectId = projectId;
        this.suiteId = suiteId;

        this.caseId = caseId;
        this.pId = pid;
        this.isLeaf = isLeaf;
    }

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "case_id", insertable = false, updatable = false)
    private TestCase testCase;
    @Column(name = "case_id")
    private Long caseId;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "suite_id", insertable = false, updatable = false)
    private TestSuite suite;
    @Column(name = "suite_id")
    private Long suiteId;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "project_id", insertable = false, updatable = false)
    private TestProject project;
    @Column(name = "project_id")
    private Long projectId;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "create_by_id", insertable = false, updatable = false)
    private TestUser createBy;
    @Column(name = "create_by_id")
    private Long createById;

    public Boolean getLeaf() {
        return isLeaf;
    }

    public void setLeaf(Boolean leaf) {
        isLeaf = leaf;
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

    public Long getpId() {
        return pId;
    }

    public void setpId(Long pId) {
        this.pId = pId;
    }

    public TestCase getTestCase() {
        return testCase;
    }

    public void setTestCase(TestCase testCase) {
        this.testCase = testCase;
    }

    public Long getCaseId() {
        return caseId;
    }

    public void setCaseId(Long caseId) {
        this.caseId = caseId;
    }

    public TestUser getCreateBy() {
        return createBy;
    }

    public void setCreateBy(TestUser createBy) {
        this.createBy = createBy;
    }

    public Long getCreateById() {
        return createById;
    }

    public void setCreateById(Long createById) {
        this.createById = createById;
    }

    public TestSuite getSuite() {
        return suite;
    }

    public void setSuite(TestSuite suite) {
        this.suite = suite;
    }

    public Long getSuiteId() {
        return suiteId;
    }

    public void setSuiteId(Long suiteId) {
        this.suiteId = suiteId;
    }

}
