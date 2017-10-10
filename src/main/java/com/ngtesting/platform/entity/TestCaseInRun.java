package com.ngtesting.platform.entity;

import javax.persistence.*;

@Entity
@Table(name = "tst_case_in_run")
public class TestCaseInRun extends BaseEntity {
    private static final long serialVersionUID = -2393416384079250976L;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "case_id", insertable = false, updatable = false)
    private TestCase testCase;

    @Column(name = "case_id")
    private Long caseId;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "run_id", insertable = false, updatable = false)
    private TestRun run;

    @Column(name = "run_id")
    private Long runId;

    @Enumerated(EnumType.STRING)
    private CaseStatusInRun status = CaseStatusInRun.untest;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "create_by_id", insertable = false, updatable = false)
    private TestUser createBy;

    @Column(name = "create_by_id")
    private Long createById;

    public TestCaseInRun() {
        super();
    }
    public TestCaseInRun(Long runId, Long caseId) {
        super();
        this.runId = runId;
        this.caseId = caseId;
    }

    public static enum CaseStatusInRun {
        untest("untest"),
        pass("pass"),
        fail("fail"),
        block("block");

        CaseStatusInRun(String val) {
            this.val = val;
        }

        private String val;
        public String toString() {
            return val;
        }
    }

    public CaseStatusInRun getStatus() {
        return status;
    }

    public void setStatus(CaseStatusInRun status) {
        this.status = status;
    }

    public TestRun getRun() {
        return run;
    }

    public void setRun(TestRun run) {
        this.run = run;
    }

    public Long getRunId() {
        return runId;
    }

    public void setRunId(Long runId) {
        this.runId = runId;
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

}
