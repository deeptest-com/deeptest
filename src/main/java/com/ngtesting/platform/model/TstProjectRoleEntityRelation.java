package com.ngtesting.platform.model;


public class TstProjectRoleEntityRelation extends BaseModel {
	private static final long serialVersionUID = -2782548788965037290L;

    private Integer projectId;
    private Integer projectRoleId;

    private String projectRoleName;

    private Integer entityId;

    private String entityName;

    private Boolean selected;
    private Boolean selecting;

	private String type;

    public TstProjectRoleEntityRelation() {

    }

    public TstProjectRoleEntityRelation(Integer id, Integer projectId, Integer entityId, Integer projectRoleId,
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

	public Integer getProjectId() {
        return projectId;
    }

    public void setProjectId(Integer projectId) {
        this.projectId = projectId;
    }

    public Integer getProjectRoleId() {
		return projectRoleId;
	}

	public void setProjectRoleId(Integer projectRoleId) {
		this.projectRoleId = projectRoleId;
	}

	public String getProjectRoleName() {
		return projectRoleName;
	}

	public void setProjectRoleName(String projectRoleName) {
		this.projectRoleName = projectRoleName;
	}

	public Integer getEntityId() {
		return entityId;
	}

	public void setEntityId(Integer entityId) {
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
