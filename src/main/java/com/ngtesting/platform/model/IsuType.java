package com.ngtesting.platform.model;

public class IsuType extends BaseModel {
	private static final long serialVersionUID = -2904250132388811594L;

	private String label;
	private String code;
    private String descr;
    private Integer displayOrder;
    private Boolean isBuildIn = false;
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
	public Integer getDisplayOrder() {
		return displayOrder;
	}
	public void setDisplayOrder(Integer displayOrder) {
		this.displayOrder = displayOrder;
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
