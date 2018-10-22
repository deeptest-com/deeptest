package com.ngtesting.platform.model;

public class TstCaseExeStatus extends BaseModel {
	private static final long serialVersionUID = -2071266644244632484L;

	private String label;
	private String value;

    private String descr;
    private Integer ordr;
    private Boolean isFinal;
    private Boolean isBuildIn = false;
    private Integer orgId;

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

	public String getDescr() {
		return descr;
	}
	public void setDescr(String descr) {
		this.descr = descr;
	}

	public Integer getOrdr() {
		return ordr;
	}

	public void setOrdr(Integer ordr) {
		this.ordr = ordr;
	}

	public Boolean getFinal() {
		return isFinal;
	}

	public void setFinal(Boolean aFinal) {
		isFinal = aFinal;
	}

	public Boolean getBuildIn() {
		return isBuildIn;
	}

	public void setBuildIn(Boolean buildIn) {
		isBuildIn = buildIn;
	}

	public Boolean getIsFinal() {
		return isFinal;
	}
	public void setIsFinal(Boolean isFinal) {
		this.isFinal = isFinal;
	}
	public Boolean getIsBuildIn() {
		return isBuildIn;
	}
	public void setIsBuildIn(Boolean isBuildIn) {
		this.isBuildIn = isBuildIn;
	}
	public Integer getOrgId() {
		return orgId;
	}
	public void setOrgId(Integer orgId) {
		this.orgId = orgId;
	}

}
