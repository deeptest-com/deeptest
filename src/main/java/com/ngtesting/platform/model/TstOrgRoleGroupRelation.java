package com.ngtesting.platform.model;


import java.io.Serializable;

public class TstOrgRoleGroupRelation implements Serializable {

	private static final long serialVersionUID = -2031588520117561747L;
	private Integer orgId;

	private Integer orgRoleId;

    private String orgRoleName;

	private Integer orgGroupId;

    private String groupName;

    private Boolean selected;
    private Boolean selecting;

	public Integer getOrgGroupId() {
		return orgGroupId;
	}

	public void setOrgGroupId(Integer orgGroupId) {
		this.orgGroupId = orgGroupId;
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

	public String getGroupName() {
		return groupName;
	}

	public void setGroupName(String groupName) {
		this.groupName = groupName;
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
