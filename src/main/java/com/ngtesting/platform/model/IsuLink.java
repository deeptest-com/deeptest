package com.ngtesting.platform.model;


public class IsuLink extends BaseModel {
	private static final long serialVersionUID = 7948735298148360241L;

    private Integer reasonId;
	private String reasonName;

	private Integer srcIssueId;
	private Integer dictIssueId;
	private Integer userId;

    public Integer getReasonId() {
        return reasonId;
    }

    public void setReasonId(Integer reasonId) {
        this.reasonId = reasonId;
    }

    public String getReasonName() {
        return reasonName;
    }

    public void setReasonName(String reasonName) {
        this.reasonName = reasonName;
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
