package com.ngtesting.platform.model;


public class TstOrgRole extends BaseModel {
	private static final long serialVersionUID = 2846494844575998128L;

	private String code;
    private String name;
    private String descr;

    private Long orgId;

	public static enum OrgRoleCode {
		org_admin("org_admin", "组织管理员"),
		site_admin("site_admin", "站点管理员"),
		project_admin("project_admin", "项目管理员");

		private OrgRoleCode(String code, String name) {
			this.code = code;
			this.name = name;
		}

		public String code;
		public String name;
	}

	public String getCode() {
		return code;
	}

	public void setCode(String code) {
		this.code = code;
	}

	public String getName() {
		return name;
	}

	public void setName(String name) {
		this.name = name;
	}

	public Long getOrgId() {
		return orgId;
	}

	public void setOrgId(Long orgId) {
		this.orgId = orgId;
	}

	public String getDescr() {
		return descr;
	}

	public void setDescr(String descr) {
		this.descr = descr;
	}

}
