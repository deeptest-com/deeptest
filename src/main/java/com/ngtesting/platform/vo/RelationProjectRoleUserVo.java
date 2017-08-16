package com.ngtesting.platform.vo;


public class RelationProjectRoleUserVo extends BaseVo {
	private static final long serialVersionUID = -2782548788965037290L;

    private Long projectId;
    private Long projectRoleId;
    
    private String projectRoleName;

    private Long userId;

    private String userName;
    
    private Boolean selected;
    private Boolean selecting;

    public RelationProjectRoleUserVo() {

    }

    public RelationProjectRoleUserVo(Long projectId, Long userId, Long projectRoleId, String projectRoleName, String userName) {
        this.projectId = projectId;
        this.userId = userId;
        this.projectRoleId = projectRoleId;
        this.projectRoleName = projectRoleName;
        this.userName = userName;
    }

    public Long getProjectId() {
        return projectId;
    }

    public void setProjectId(Long projectId) {
        this.projectId = projectId;
    }

    public Long getProjectRoleId() {
		return projectRoleId;
	}

	public void setProjectRoleId(Long projectRoleId) {
		this.projectRoleId = projectRoleId;
	}

	public String getProjectRoleName() {
		return projectRoleName;
	}

	public void setProjectRoleName(String projectRoleName) {
		this.projectRoleName = projectRoleName;
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

    
    
}
