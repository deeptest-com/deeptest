package com.ngtesting.platform.entity;

import javax.persistence.*;
import java.util.Date;

@Entity
@Table(name = "tst_case_comments")
public class TestCaseComments extends BaseEntity {
	private static final long serialVersionUID = -5934497865835276588L;

    private String summary;
	private String content;

    private Date changeTime = new Date();

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "user_id", insertable = false, updatable = false)
    private TestUser user;

    @Column(name = "user_id")
    private Long userId;

	@ManyToOne(fetch = FetchType.LAZY)
	@JoinColumn(name = "test_case_id", insertable = false, updatable = false)
	private TestCase testCase;

	@Column(name = "test_case_id")
	private Long testCaseId;

	public TestCaseComments() {
		super();
	}

    public Date getChangeTime() {
        return changeTime;
    }

    public void setChangeTime(Date changeTime) {
        this.changeTime = changeTime;
    }

    public String getSummary() {
        return summary;
    }

    public void setSummary(String summary) {
        this.summary = summary;
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
