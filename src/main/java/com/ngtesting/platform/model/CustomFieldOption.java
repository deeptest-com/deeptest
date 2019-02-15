package com.ngtesting.platform.model;

public class CustomFieldOption extends BaseModel {

	private static final long serialVersionUID = 8057353932992599921L;
	private String label;
	private String descr;
	private Integer ordr;
	private Integer fieldId;
	private Integer orgId;
	private Boolean defaultVal = false;
    private Boolean buildIn = false;

	public Integer getOrgId() {
		return orgId;
	}

	public void setOrgId(Integer orgId) {
		this.orgId = orgId;
	}

	public Integer getFieldId() {
		return fieldId;
	}

	public void setFieldId(Integer fieldId) {
		this.fieldId = fieldId;
	}

    public String getLabel() {
		return label;
	}

	public void setLabel(String label) {
		this.label = label;
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

	public Boolean getDefaultVal() {
		return defaultVal;
	}

	public void setDefaultVal(Boolean aDefault) {
		defaultVal = aDefault;
	}

    public Boolean getBuildIn() {
        return buildIn;
    }

    public void setBuildIn(Boolean buildIn) {
      this.buildIn = buildIn;
    }
}
