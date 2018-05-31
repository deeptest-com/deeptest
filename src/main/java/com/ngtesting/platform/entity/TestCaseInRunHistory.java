package com.ngtesting.platform.entity;

import javax.persistence.*;

@Entity
@Table(name = "tst_case_in_run_history")
public class TestCaseInRunHistory extends BaseEntity {
    private static final long serialVersionUID = 729187057688302743L;

    private String title;
    @Column(name = "descr", length = 1000)
    private String descr;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "case_id", insertable = false, updatable = false)
    private TestCaseInRun testCaseInRun;

    @Column(name = "case_in_run_id")
    private Long testCaseInRunId;

    public TestCaseInRun getTestCaseInRun() {
        return testCaseInRun;
    }

    public void setTestCaseInRun(TestCaseInRun testCaseInRun) {
        this.testCaseInRun = testCaseInRun;
    }

    public Long getTestCaseInRunId() {
        return testCaseInRunId;
    }

    public void setTestCaseInRunId(Long testCaseInRunId) {
        this.testCaseInRunId = testCaseInRunId;
    }

    public String getDescr() {
        return descr;
    }

    public void setDescr(String descr) {
        this.descr = descr;
    }

    public String getTitle() {
        return title;
    }

    public void setTitle(String title) {
        this.title = title;
    }
}
