package com.ngtesting.platform.model;


public class TstOrgGroup extends BaseModel {
	private static final long serialVersionUID = 9011842501304591654L;
	private String name;
    private String descr;

    private Integer orgId;
	private Boolean buildIn = false;

    private Boolean selecting;
    private Boolean selected;

	private String type = "group";

	public String getType() {
		return type;
	}

	public void setType(String type) {
		this.type = type;
	}

	public String getName() {
		return name;
	}

	public void setName(String name) {
		this.name = name;
	}

	public String getDescr() {
		return descr;
	}

	public void setDescr(String descr) {
		this.descr = descr;
	}

	public Integer getOrgId() {
		return orgId;
	}

	public void setOrgId(Integer orgId) {
		this.orgId = orgId;
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

	public Boolean getBuildIn() {
		return buildIn;
	}

	public void setBuildIn(Boolean buildIn) {
    this.buildIn = buildIn;
	}
}
