package com.ngtesting.platform.model;

import java.util.Date;

public class TstUser {
    private Integer id;

    private String email;
    private String password;
    private String nickname;
    private String phone;
    private String avatar;

    private String token;
    private Date lastLoginTime;
    private Integer leftSizeCase = 300;
    private Integer leftSizeIssue = 200;

    private Integer defaultOrgId;
    private Integer defaultPrjId;

    public Integer getId() {
        return id;
    }

    public void setId(Integer id) {
        this.id = id;
    }

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public String getPhone() {
        return phone;
    }

    public void setPhone(String phone) {
        this.phone = phone;
    }

    public String getNickname() {
        return nickname;
    }

    public void setNickname(String nickname) {
        this.nickname = nickname;
    }

    public String getPassword() {
        return password;
    }

    public void setPassword(String password) {
        this.password = password;
    }

    public String getToken() {
        return token;
    }

    public void setToken(String token) {
        this.token = token;
    }

    public String getAvatar() {
        return avatar;
    }

    public void setAvatar(String avatar) {
        this.avatar = avatar;
    }

    public Date getLastLoginTime() {
        return lastLoginTime;
    }

    public void setLastLoginTime(Date lastLoginTime) {
        this.lastLoginTime = lastLoginTime;
    }

    public Integer getDefaultOrgId() {
        return defaultOrgId;
    }

    public void setDefaultOrgId(Integer defaultOrgId) {
        this.defaultOrgId = defaultOrgId;
    }

    public Integer getDefaultPrjId() {
        return defaultPrjId;
    }

    public void setDefaultPrjId(Integer defaultPrjId) {
        this.defaultPrjId = defaultPrjId;
    }

    public Integer getLeftSizeCase() {
        return leftSizeCase;
    }

    public void setLeftSizeCase(Integer leftSizeCase) {
        this.leftSizeCase = leftSizeCase;
    }

    public Integer getLeftSizeIssue() {
        return leftSizeIssue;
    }

    public void setLeftSizeIssue(Integer leftSizeIssue) {
        this.leftSizeIssue = leftSizeIssue;
    }
}
