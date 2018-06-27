package com.ngtesting.platform.vo;


public class RelationOrgGroupUserVo extends BaseVo {
	private static final long serialVersionUID = -7361457861754458408L;

	private Long orgId;
	
	private Long orgGroupId;
    
    private String orgGroupName;
    
    private Long userId;
    
    private String userName;
    
    private Boolean selected;
    private Boolean selecting;

	public Long getOrgGroupId() {
		return orgGroupId;
	}

	public void setOrgGroupId(Long orgGroupId) {
		this.orgGroupId = orgGroupId;
	}

	public String getOrgGroupName() {
		return orgGroupName;
	}

	public void setOrgGroupName(String orgGroupName) {
		this.orgGroupName = orgGroupName;
	}

	public Long getUserId() {
		return userId;
	}

	public void setUserId(Long userId) {
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

	public Long getOrgId() {
		return orgId;
	}

	public void setOrgId(Long orgId) {
		this.orgId = orgId;
	}

}
