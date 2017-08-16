package com.ngtesting.platform.vo;


public class RelationProjectRoleEntityVo extends BaseVo {
	private static final long serialVersionUID = -2782548788965037290L;

    private Long projectId;
    private Long projectRoleId;
    
    private String projectRoleName;

    private Long entityId;

    private String entityName;
    
    private Boolean selected;
    private Boolean selecting;

	private String type;

    public RelationProjectRoleEntityVo() {

    }

    public RelationProjectRoleEntityVo(Long id, Long projectId, Long entityId, Long projectRoleId,
                                       String projectRoleName, String entityName, String type) {
        this.id = id;
        this.projectId = projectId;
        this.entityId = entityId;
        this.projectRoleId = projectRoleId;
        this.projectRoleName = projectRoleName;
        this.entityName = entityName;
        this.type = type;
    }


	public String getType() {
		return type;
	}

	public void setType(String type) {
		this.type = type;
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

	public Long getEntityId() {
		return entityId;
	}

	public void setEntityId(Long entityId) {
		this.entityId = entityId;
	}

	public String getEntityName() {
		return entityName;
	}

	public void setEntityName(String entityName) {
		this.entityName = entityName;
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
