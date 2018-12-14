package com.ngtesting.platform.model;

import java.util.Date;

public class IsuComments extends BaseModel {

    private static final long serialVersionUID = 2999250634324982041L;
    private String summary;
    private String content;
    private Integer userId;
    private String userName;
    private String userAvatar;
    private Integer issueId;
    private Date changeTime;

    public IsuComments() {
    }

    public IsuComments(Integer issueId, String summary, String content) {
        this.issueId = issueId;
        this.summary = summary;
        this.content = content;
    }

    public Integer getUserId() {
        return userId;
    }

    public void setUserId(Integer userId) {
        this.userId = userId;
    }

    public Date getChangeTime() {
        return changeTime;
    }

    public void setChangeTime(Date changeTime) {
        this.changeTime = changeTime;
    }

    public String getContent() {
        return content;
    }

    public void setContent(String content) {
        this.content = content;
    }

    public Integer getIssueId() {
        return issueId;
    }

    public void setIssueId(Integer issueId) {
        this.issueId = issueId;
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
