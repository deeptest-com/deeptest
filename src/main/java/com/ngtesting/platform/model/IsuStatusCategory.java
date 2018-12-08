package com.ngtesting.platform.model;

public class IsuStatusCategory extends BaseModel {
	private static final long serialVersionUID = -3974051187772599821L;

	private String label;
	private String code;
    private String descr;
    private Integer displayOrder;
    private Boolean finalVal;
    private Boolean buildIn = false;
    private Integer orgId;

	public String getLabel() {
		return label;
	}

	public void setLabel(String label) {
		this.label = label;
	}

	public String getCode() {
		return code;
	}

	public void setCode(String code) {
		this.code = code;
	}

	public Boolean getFinal() {
		return finalVal;
	}

	public void setFinal(Boolean aFinal) {
		finalVal = aFinal;
	}

	public Boolean getBuildIn() {
		return buildIn;
	}

	public void setBuildIn(Boolean buildIn) {
    this.buildIn = buildIn;
	}

	public String getDescr() {
		return descr;
	}
	public void setDescr(String descr) {
		this.descr = descr;
	}
	public Integer getDisplayOrder() {
		return displayOrder;
	}
	public void setDisplayOrder(Integer displayOrder) {
		this.displayOrder = displayOrder;
	}
	public Integer getOrgId() {
		return orgId;
	}
	public void setOrgId(Integer orgId) {
		this.orgId = orgId;
	}

}
