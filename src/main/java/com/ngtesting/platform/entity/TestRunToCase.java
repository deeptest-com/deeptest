package com.ngtesting.platform.entity;

import javax.persistence.*;

@Entity
@Table(name = "tst_run_to_case")
public class TestRunToCase extends BaseEntity {

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "run_id", insertable = false, updatable = false)
    private TestRun run;

    @Column(name = "run_id")
    private Long runId;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "case_id", insertable = false, updatable = false)
    private TestCase testCase;

    @Column(name = "case_id")
    private Long caseId;

    @Enumerated(EnumType.STRING)
    private CaseStatusInRun status = CaseStatusInRun.untest;

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
}
