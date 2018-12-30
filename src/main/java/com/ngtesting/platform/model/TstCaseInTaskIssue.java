package com.ngtesting.platform.model;


public class TstCaseInTaskIssue extends BaseModel {

	private static final long serialVersionUID = 4897059624066098332L;

	private Integer caseInTaskId;
	private Integer issueId;
    private Integer userId;

    private String title;

    public TstCaseInTaskIssue() {
    }

    public TstCaseInTaskIssue(Integer issueId, Integer caseInTaskId, Integer userId) {
        this.caseInTaskId = caseInTaskId;
        this.issueId = issueId;
        this.userId = userId;
    }

    public String getTitle() {
        return title;
    }

    public void setTitle(String title) {
        this.title = title;
    }

    public Integer getCaseInTaskId() {
        return caseInTaskId;
    }

    public void setCaseInTaskId(Integer caseInTaskId) {
        this.caseInTaskId = caseInTaskId;
    }

    public Integer getIssueId() {
		return issueId;
	}

	public void setIssueId(Integer issueId) {
		this.issueId = issueId;
	}

    public Integer getUserId() {
        return userId;
    }

    public void setUserId(Integer userId) {
        this.userId = userId;
    }
}
