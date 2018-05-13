package com.ngtesting.platform.vo;


public class UserVo extends BaseVo {
	private static final long serialVersionUID = 1988353599211843484L;
	private String name;
	private String email;
//	private String password;
    private String phone;
    private String avatar;

    private String token;
    private Long defaultOrgId;
	private String defaultOrgName;
    private Long defaultPrjId;
	private String defaultPrjName;

    private Boolean selected;
    private Boolean selecting;

	private Integer leftSizeCase;
	private Integer leftSizeIssue;

	private String type = "user";

	public UserVo() {

	}
	public UserVo(Long id, String name) {
		this.id = id;
		this.name = name;
	}

	public String getDefaultPrjName() {
		return defaultPrjName;
	}

	public void setDefaultPrjName(String defaultPrjName) {
		this.defaultPrjName = defaultPrjName;
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

	public String getType() {
		return type;
	}

	public void setType(String type) {
		this.type = type;
	}

	public String getDefaultOrgName() {
		return defaultOrgName;
	}

	public void setDefaultOrgName(String defaultOrgName) {
		this.defaultOrgName = defaultOrgName;
	}
}
