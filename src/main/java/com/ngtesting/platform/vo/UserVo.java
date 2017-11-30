package com.ngtesting.platform.vo;


import java.util.List;
import java.util.Map;

public class UserVo extends BaseVo {
	private static final long serialVersionUID = 1988353599211843484L;
	private String name;
	private String email;
//	private String password;
    private String phone;
    private String avatar;

    private String token;
    private Long defaultOrgId;
    private Long defaultProjectId;

    private Boolean selected;
    private Boolean selecting;

	private Integer caseBoardLeftSize;
	private Integer caseBoardRightSize;

	private String type = "user";

	List<OrgVo> orgs;
	List<TestProjectAccessHistoryVo> recentProjects;
	Map<String, Map<String,String>> casePropertyMap;

	Map<String,Boolean> sysPrivilege;
	Map<String, Map<String,Boolean>> orgPrivilege;
	Map<String, Map<String,Boolean>> projectPrivilege;

	public UserVo() {

	}
	public UserVo(Long id, String name) {
		this.id = id;
		this.name = name;
	}

	public List<OrgVo> getOrgs() {
		return orgs;
	}

	public void setOrgs(List<OrgVo> orgs) {
		this.orgs = orgs;
	}

	public List<TestProjectAccessHistoryVo> getRecentProjects() {
		return recentProjects;
	}

	public void setRecentProjects(List<TestProjectAccessHistoryVo> recentProjects) {
		this.recentProjects = recentProjects;
	}

	public Map<String, Map<String, String>> getCasePropertyMap() {
		return casePropertyMap;
	}

	public void setCasePropertyMap(Map<String, Map<String, String>> casePropertyMap) {
		this.casePropertyMap = casePropertyMap;
	}

	public Map<String, Boolean> getSysPrivilege() {
		return sysPrivilege;
	}

	public void setSysPrivilege(Map<String, Boolean> sysPrivilege) {
		this.sysPrivilege = sysPrivilege;
	}

	public Map<String, Map<String, Boolean>> getOrgPrivilege() {
		return orgPrivilege;
	}
	public void setOrgPrivilege(Map<String, Map<String, Boolean>> orgPrivilege) {
		this.orgPrivilege = orgPrivilege;
	}

	public Map<String, Map<String, Boolean>> getProjectPrivilege() {
		return projectPrivilege;
	}

	public void setProjectPrivilege(Map<String, Map<String, Boolean>> projectPrivilege) {
		this.projectPrivilege = projectPrivilege;
	}

	public Integer getCaseBoardLeftSize() {
		return caseBoardLeftSize;
	}

	public void setCaseBoardLeftSize(Integer caseBoardLeftSize) {
		this.caseBoardLeftSize = caseBoardLeftSize;
	}

	public Integer getCaseBoardRightSize() {
		return caseBoardRightSize;
	}

	public void setCaseBoardRightSize(Integer caseBoardRightSize) {
		this.caseBoardRightSize = caseBoardRightSize;
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
//	public String getPassword() {
//		return password;
//	}
//	public void setPassword(String password) {
//		this.password = password;
//	}
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
	public Long getDefaultProjectId() {
		return defaultProjectId;
	}
	public void setDefaultProjectId(Long defaultProjectId) {
		this.defaultProjectId = defaultProjectId;
	}

	public String getType() {
		return type;
	}

	public void setType(String type) {
		this.type = type;
	}
}
