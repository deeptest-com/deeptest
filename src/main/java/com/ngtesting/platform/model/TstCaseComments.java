package com.ngtesting.platform.model;

public class TstCaseComments extends BaseModel {
    private static final long serialVersionUID = -624520281526573819L;

    private String summary;
    private String content;
    private Integer userId;
    private String userName;
    private String userAvatar;
    private Integer caseId;

    public TstCaseComments() {
    }
    public TstCaseComments(Integer caseId, String summary) {
        this.caseId = caseId;
        this.summary = summary;
    }

    public Integer getUserId() {
        return userId;
    }

    public void setUserId(Integer userId) {
        this.userId = userId;
    }

    public String getContent() {
        return content;
    }

    public void setContent(String content) {
        this.content = content;
    }

    public Integer getCaseId() {
        return caseId;
    }

    public void setCaseId(Integer caseId) {
        this.caseId = caseId;
    }

    public String getSummary() {
        return summary;
    }

    public void setSummary(String summary) {
        this.summary = summary;
    }

    public String getUserName() {
        return userName;
    }

    public void setUserName(String userName) {
        this.userName = userName;
    }

    public String getUserAvatar() {
        return userAvatar;
    }

    public void setUserAvatar(String userAvatar) {
        this.userAvatar = userAvatar;
    }
}
