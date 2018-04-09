package com.ngtesting.platform.entity;

import javax.persistence.*;

@Entity
@Table(name = "tst_case_history")
public class IssueHistory extends BaseEntity {

    private static final long serialVersionUID = 1545026595607666706L;
    private String title;
    @Column(name = "descr", length = 1000)
    private String descr;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "case_id", insertable = false, updatable = false)
    private TestCase testCase;

    @Column(name = "case_id")
    private Long testCaseId;

    public TestCase getTestCase() {
        return testCase;
    }

    public void setTestCase(TestCase testCase) {
        this.testCase = testCase;
    }

    public Long getTestCaseId() {
        return testCaseId;
    }

    public void setTestCaseId(Long testCaseId) {
        this.testCaseId = testCaseId;
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
