package com.ngtesting.platform.model;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;

import java.util.Date;
import java.util.HashSet;
import java.util.Set;

@JsonIgnoreProperties(value={"salt", "password"})
public class TstUser extends BaseModel {
    private static final long serialVersionUID = 8137369995938797198L;

    private String email;
    private String password;
    private String nickname;
    private String phone;
    private String avatar = "upload/sample/user/avatar.png";

    private String temp;
    private String token;
    private Date lastLoginTime;
    private Integer leftSizeDesign = 300;
    private Integer leftSizeExe = 200;
    private Integer leftSizeIssue = 300;

    private String issueView = "table";
    private String issueColumns = "";
    private String issueFileds = "";

    private Integer defaultOrgId;
    private String defaultOrgName;
    private Integer defaultPrjId;
    private String defaultPrjName;

    private Boolean selected;
    private Boolean selecting;

    private Boolean locked = Boolean.FALSE;

    private String type = "user";

    private String authCacheKey;
    private transient String salt;
    private Set<String> roles = new HashSet<>();
    private Set<String> perms = new HashSet<>();

    public String getAuthCacheKey() {
        return id.toString();
    }

    public String getSalt() {
        return salt;
    }

    public void setSalt(String salt) {
        this.salt = salt;
    }

    public Set<String> getRoles() {
        return roles;
    }

    public void setRoles(Set<String> roles) {
        this.roles = roles;
    }

    public Set<String> getPerms() {
        return perms;
    }

    public void setPerms(Set<String> perms) {
        this.perms = perms;
    }

    public Integer getLeftSizeIssue() {
        return leftSizeIssue;
    }

    public void setLeftSizeIssue(Integer leftSizeIssue) {
        this.leftSizeIssue = leftSizeIssue;
    }

    public String getTemp() {
        return temp;
    }

    public void setTemp(String temp) {
        this.temp = temp;
    }

    public Integer getDefaultOrgId() {
        return defaultOrgId;
    }

    public void setDefaultOrgId(Integer defaultOrgId) {
        this.defaultOrgId = defaultOrgId;
    }

    public String getDefaultOrgName() {
        return defaultOrgName;
    }

    public void setDefaultOrgName(String defaultOrgName) {
        this.defaultOrgName = defaultOrgName;
    }

    public Integer getDefaultPrjId() {
        return defaultPrjId;
    }

    public void setDefaultPrjId(Integer defaultPrjId) {
        this.defaultPrjId = defaultPrjId;
    }

    public String getDefaultPrjName() {
        return defaultPrjName;
    }

    public void setDefaultPrjName(String defaultPrjName) {
        this.defaultPrjName = defaultPrjName;
    }

    public Boolean getSelected() {
        return selected;
    }

    public void setSelected(Boolean selected) {
        this.selected = selected;
    }

    public Boolean getSelecting() {
        return selecting;
    }

    public void setSelecting(Boolean selecting) {
        this.selecting = selecting;
    }

    public String getType() {
        return type;
    }

    public void setType(String type) {
        this.type = type;
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

    public Integer getLeftSizeDesign() {
        return leftSizeDesign;
    }

    public void setLeftSizeDesign(Integer leftSizeDesign) {
        this.leftSizeDesign = leftSizeDesign;
    }

    public Integer getLeftSizeExe() {
        return leftSizeExe;
    }

    public void setLeftSizeExe(Integer leftSizeExe) {
        this.leftSizeExe = leftSizeExe;
    }

    public String getIssueView() {
        return issueView;
    }

    public void setIssueView(String issueView) {
        this.issueView = issueView;
    }

    public String getIssueColumns() {
        return issueColumns;
    }

    public void setIssueColumns(String issueColumns) {
        this.issueColumns = issueColumns;
    }

    public String getIssueFileds() {
        return issueFileds;
    }

    public void setIssueFileds(String issueFileds) {
        this.issueFileds = issueFileds;
    }

    public Boolean getLocked() {
        return locked;
    }

    public void setLocked(Boolean locked) {
        this.locked = locked;
    }
}
