package com.ngtesting.platform.model;


public class IsuLink extends BaseModel {
	private static final long serialVersionUID = 7948735298148360241L;

	private String reason;

	private Integer srcIssueId;
	private Integer dictIssueId;
	private Integer userId;

	public IsuLink() {
	}
	public IsuLink(String reason, Integer srcIssueId, Integer dictIssueId, Integer userId) {
		this.reason = reason;
		this.srcIssueId = srcIssueId;
		this.dictIssueId = dictIssueId;
		this.userId = userId;
	}

	public String getReason() {
		return reason;
	}

	public void setReason(String reason) {
		this.reason = reason;
	}

	public Integer getSrcIssueId() {
		return srcIssueId;
	}

	public void setSrcIssueId(Integer srcIssueId) {
		this.srcIssueId = srcIssueId;
	}

	public Integer getDictIssueId() {
		return dictIssueId;
	}

	public void setDictIssueId(Integer dictIssueId) {
		this.dictIssueId = dictIssueId;
	}

	public Integer getUserId() {
		return userId;
	}

	public void setUserId(Integer userId) {
		this.userId = userId;
	}
}
