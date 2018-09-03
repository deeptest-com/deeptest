package com.ngtesting.platform.model;


import java.io.Serializable;

public class TstOrgRoleGroupRelation implements Serializable {

	private static final long serialVersionUID = -2031588520117561747L;
	private Integer orgId;

	private Integer orgRoleId;

    private String orgRoleName;

    private Integer groupId;

    private String groupName;

    private Boolean selected;
    private Boolean selecting;

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

	public Integer getGroupId() {
		return groupId;
	}

	public void setGroupId(Integer groupId) {
		this.groupId = groupId;
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
