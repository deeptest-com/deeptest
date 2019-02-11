package com.ngtesting.platform.model;

public class IsuStatusCategoryDefine extends BaseModel {
	private static final long serialVersionUID = -3974051187772599821L;

	private String label;
	private String value;
    private Integer ordr;
    private Boolean finalVal;
    private Boolean buildIn = false;

	public String getLabel() {
		return label;
	}

	public void setLabel(String label) {
		this.label = label;
	}

	public String getValue() {
		return value;
	}

	public void setValue(String value) {
		this.value = value;
	}

	public Integer getOrdr() {
		return ordr;
	}

	public void setOrdr(Integer ordr) {
		this.ordr = ordr;
	}

	public Boolean getFinalVal() {
		return finalVal;
	}

	public void setFinalVal(Boolean finalVal) {
		this.finalVal = finalVal;
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

}
