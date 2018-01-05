package com.ngtesting.platform.entity;

import javax.persistence.*;
import java.util.Date;

@Entity
@Table(name = "tst_case_comments")
public class TestCaseComments extends BaseEntity {
	private static final long serialVersionUID = -5934497865835276588L;

    private String summary;
	private String content;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "update_by_id", insertable = false, updatable = false)
    private TestUser updateBy;

    @Column(name = "update_by_id")
    private Long updateById;

	@ManyToOne(fetch = FetchType.LAZY)
	@JoinColumn(name = "test_case_id", insertable = false, updatable = false)
	private TestCase testCase;

	@Column(name = "test_case_id")
	private Long testCaseId;

	public TestCaseComments() {
		super();
	}
	public TestCaseComments(Long caseId, String content) {
		super();
		this.testCaseId = caseId;
		this.content = content;
        this.updateTime = new Date();
	}

    public String getSummary() {
        return summary;
    }

    public void setSummary(String summary) {
        this.summary = summary;
    }

    public TestUser getUpdateBy() {
        return updateBy;
    }

    public void setUpdateBy(TestUser updateBy) {
        this.updateBy = updateBy;
    }

    public Long getUpdateById() {
        return updateById;
    }

    public void setUpdateById(Long updateById) {
        this.updateById = updateById;
    }

    public String getContent() {
        return content;
    }

    public void setContent(String content) {
        this.content = content;
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
