package com.ngtesting.platform.vo;

import java.util.Date;

public class TestCaseCommentsVo extends BaseVo {
    private static final long serialVersionUID = -624520281526573819L;

    private String summary;
    private String content;
    private String userName;
    private String userAvatar;
    private Long testCaseId;

    public TestCaseCommentsVo() {
    }

    public TestCaseCommentsVo(Long id, String summary, String content, String userName, String userAvatar,
                              Long testCaseId, Date updateTime) {
        this.id = id;
        this.summary = summary;
        this.userName = userName;
        this.userAvatar = userAvatar;
        this.content = content;
        this.testCaseId = testCaseId;
        this.updateTime = updateTime;
    }

    public String getContent() {
        return content;
    }

    public void setContent(String content) {
        this.content = content;
    }

    public Long getTestCaseId() {
        return testCaseId;
    }

    public void setTestCaseId(Long testCaseId) {
        this.testCaseId = testCaseId;
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
