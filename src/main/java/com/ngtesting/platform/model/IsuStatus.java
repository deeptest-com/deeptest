package com.ngtesting.platform.model;

public class IsuStatus extends BaseModel {

	private static final long serialVersionUID = 1657004478821957505L;
	private String label;
	private String value;
    private String descr;
    private Integer ordr;

    private Boolean isDefault = false;
    private Boolean isFinal;
    private Boolean isBuildIn = false;
    private Integer orgId;

    public Boolean getDefault() {
        return isDefault;
    }

    public void setDefault(Boolean aDefault) {
        isDefault = aDefault;
    }

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
