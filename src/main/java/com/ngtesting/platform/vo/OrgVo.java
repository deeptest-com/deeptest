package com.ngtesting.platform.vo;


import java.util.Map;

public class OrgVo extends BaseVo {
	private static final long serialVersionUID = -7115478651798848319L;
	private String name;
    private String website;
    private Boolean defaultOrg;
	Map<String, Boolean> orgPrivileges;

	public String getName() {
		return name;
	}
	public void setName(String name) {
		this.name = name;
	}
	public String getWebsite() {
		return website;
	}
	public void setWebsite(String website) {
		this.website = website;
	}
	public Boolean getDefaultOrg() {
		return defaultOrg;
	}
	public void setDefaultOrg(Boolean defaultOrg) {
		this.defaultOrg = defaultOrg;
	}

	public Map<String, Boolean> getOrgPrivileges() {
		return orgPrivileges;
	}

	public void setOrgPrivileges(Map<String, Boolean> orgPrivileges) {
		this.orgPrivileges = orgPrivileges;
	}
}
