package com.ngtesting.platform.entity;

import com.ngtesting.platform.config.Constant;

import javax.persistence.*;

@Entity
@Table(name = "tst_case_in_run")
public class TestCaseInRun extends BaseEntity {
    private static final long serialVersionUID = -2393416384079250976L;

//    @Enumerated(EnumType.STRING)
    private String status = "untest";

    private Integer ordr;
    private Long pId;

    private String result;
    @Transient
    private String key;

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

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "create_by_id", insertable = false, updatable = false)
    private TestUser createBy;

    @Column(name = "create_by_id")
    private Long createById;

    public TestCaseInRun() {
        super();
    }
    public TestCaseInRun(Long runId, Long caseId, Integer ordr, Long pid) {
        super();
        this.runId = runId;
        this.caseId = caseId;
        this.ordr = ordr;
        this.pId = pid;
    }

//    public static enum CaseStatusInRun {
//        untest("untest"),
//        pass("pass"),
//        fail("fail"),
//        block("block");
//
//        CaseStatusInRun(String val) {
//            this.val = val;
//        }
//
//        private String val;
//        public String toString() {
//            return val;
//        }
//    }


    public String getKey() {
        return Constant.KEY_TESTCASE_EXE + getId();
    }

    public String getResult() {
        return result;
    }

    public void setResult(String result) {
        this.result = result;
    }

    public Long getpId() {
        return pId;
    }

    public void setpId(Long pId) {
        this.pId = pId;
    }

    public Integer getOrdr() {
        return ordr;
    }

    public void setOrdr(Integer ordr) {
        this.ordr = ordr;
    }

    public String getStatus() {
        return status;
    }

    public void setStatus(String status) {
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
