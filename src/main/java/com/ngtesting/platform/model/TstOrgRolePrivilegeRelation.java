package com.ngtesting.platform.model;


public class TstOrgRolePrivilegeRelation extends BaseModel {
	private static final long serialVersionUID = -7584449391114803263L;

	private Integer orgId;

	private Integer orgRoleId;

    private String orgRoleName;

    private Integer orgPrivilegeId;

    private String orgPrivilegeName;

    private Boolean selected;
    private Boolean selecting;

    public Integer getOrgPrivilegeId() {
        return orgPrivilegeId;
    }

    public void setOrgPrivilegeId(Integer orgPrivilegeId) {
        this.orgPrivilegeId = orgPrivilegeId;
    }

    public String getOrgPrivilegeName() {
        return orgPrivilegeName;
    }

    public void setOrgPrivilegeName(String orgPrivilegeName) {
        this.orgPrivilegeName = orgPrivilegeName;
    }

    public Integer getOrgId() {
		return orgId;
	}

	public void setOrgId(Integer orgId) {
		this.orgId = orgId;
	}

	public Integer getOrgRoleId() {
		return orgRoleId;
	}

	public void setOrgRoleId(Integer orgRoleId) {
		this.orgRoleId = orgRoleId;
	}

	public String getOrgRoleName() {
		return orgRoleName;
	}

	public void setOrgRoleName(String orgRoleName) {
		this.orgRoleName = orgRoleName;
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

}
