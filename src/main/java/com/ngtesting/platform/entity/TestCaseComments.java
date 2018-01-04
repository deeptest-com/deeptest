package com.ngtesting.platform.entity;

import javax.persistence.*;

@Entity
@Table(name = "tst_case_comments")
public class TestCaseComments extends BaseEntity {
	private static final long serialVersionUID = -5934497865835276588L;

	private String content;

	private Integer ordr;

	@ManyToOne(fetch = FetchType.LAZY)
	@JoinColumn(name = "test_case_id", insertable = false, updatable = false)
	private TestCase testCase;

	@Column(name = "test_case_id")
	private Long testCaseId;

	public TestCaseComments() {
		super();
	}
	public TestCaseComments(Long caseId, String content, Integer ordr) {
		super();
		this.testCaseId = caseId;
		this.content = content;
		this.ordr = ordr;
	}

    public String getContent() {
        return content;
    }

    public void setContent(String content) {
        this.content = content;
    }

    public Integer getOrdr() {
        return ordr;
    }

    public void setOrdr(Integer ordr) {
        this.ordr = ordr;
    }

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
}
