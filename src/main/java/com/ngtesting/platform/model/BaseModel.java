package com.ngtesting.platform.model;

import java.util.Date;

public class BaseModel {
    private Long id;

    private String email;
    private String password;
    private String name;
    private String phone;
    private String avatar;

    private String token;
    private Date lastLoginTime;
    private Integer leftSizeCase = 300;
    private Integer leftSizeIssue = 200;

    private Long defaultOrgId;
    private Long defaultPrjId;

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
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

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
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

    public Long getDefaultOrgId() {
        return defaultOrgId;
    }

    public void setDefaultOrgId(Long defaultOrgId) {
        this.defaultOrgId = defaultOrgId;
    }

    public Long getDefaultPrjId() {
        return defaultPrjId;
    }

    public void setDefaultPrjId(Long defaultPrjId) {
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
