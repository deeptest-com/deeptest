package com.ngtesting.platform.model;

public class TstMsg extends BaseModel {
    private static final long serialVersionUID = -6418363700406245211L;

    private String title;
    private String uri;
    private Boolean isRead;

    private Integer assigneeId;
    private String assigneeName;
    private String assigneeAvatar;

    private Integer userId;
    private String userName;
    private String userAvatar;

    public Integer getAssigneeId() {
        return assigneeId;
    }

    public void setAssigneeId(Integer assigneeId) {
        this.assigneeId = assigneeId;
    }

    public String getAssigneeName() {
        return assigneeName;
    }

    public void setAssigneeName(String assigneeName) {
        this.assigneeName = assigneeName;
    }

    public String getAssigneeAvatar() {
        return assigneeAvatar;
    }

    public void setAssigneeAvatar(String assigneeAvatar) {
        this.assigneeAvatar = assigneeAvatar;
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

    public Boolean getRead() {
        return isRead;
    }

    public void setRead(Boolean read) {
        isRead = read;
    }

    public String getTitle() {
        return title;
    }

    public void setTitle(String title) {
        this.title = title;
    }

    public String getUri() {
        return uri;
    }

    public void setUri(String uri) {
        this.uri = uri;
    }

    public Integer getUserId() {
        return userId;
    }

    public void setUserId(Integer userId) {
        this.userId = userId;
    }

}
