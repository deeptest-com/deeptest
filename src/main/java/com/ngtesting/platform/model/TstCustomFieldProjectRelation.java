package com.ngtesting.platform.model;

public class TstCustomFieldProjectRelation extends BaseModel {
	private static final long serialVersionUID = -5497414330205266303L;

    private Integer orgId;

	private Integer customFieldId;

	private Integer projectId;

	private String projectName;
    private String projectType;

	private Boolean selected;
	private Boolean selecting;

    public Integer getOrgId() {
        return orgId;
    }

    public void setOrgId(Integer orgId) {
        this.orgId = orgId;
    }

    public Integer getCustomFieldId() {
        return customFieldId;
    }

    public void setCustomFieldId(Integer customFieldId) {
        this.customFieldId = customFieldId;
    }

    public Integer getProjectId() {
        return projectId;
    }

    public void setProjectId(Integer projectId) {
        this.projectId = projectId;
    }

    public String getProjectName() {
        return projectName;
    }

    public void setProjectName(String projectName) {
        this.projectName = projectName;
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

    public String getProjectType() {
        return projectType;
    }

    public void setProjectType(String projectType) {
        this.projectType = projectType;
    }
}
