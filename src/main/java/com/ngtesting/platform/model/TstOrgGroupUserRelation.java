package com.ngtesting.platform.model;


public class TstOrgGroupUserRelation extends BaseModel {
	private static final long serialVersionUID = -7361457861754458408L;

	private Integer orgId;

	private Integer orgGroupId;

    private String orgGroupName;

    private Integer userId;

    private String userName;

    private Boolean selected;
    private Boolean selecting;

	public Integer getOrgGroupId() {
		return orgGroupId;
	}

	public void setOrgGroupId(Integer orgGroupId) {
		this.orgGroupId = orgGroupId;
	}

	public String getOrgGroupName() {
		return orgGroupName;
	}

	public void setOrgGroupName(String orgGroupName) {
		this.orgGroupName = orgGroupName;
	}

	public Integer getUserId() {
		return userId;
	}

	public void setUserId(Integer userId) {
		this.userId = userId;
	}

	public String getUserName() {
		return userName;
	}

	public void setUserName(String userName) {
		this.userName = userName;
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

	public Integer getOrgId() {
		return orgId;
	}

	public void setOrgId(Integer orgId) {
		this.orgId = orgId;
	}

}
