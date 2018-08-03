package com.ngtesting.platform.model;

public class TstCaseType extends BaseModel {
	private static final long serialVersionUID = -2071266644244632484L;

	private String label;
	private String value;

	private String descr;
    private Boolean isDefault = false;
    private Integer ordr;
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

	public Integer getOrgId() {
		return orgId;
	}
	public void setOrgId(Integer orgId) {
		this.orgId = orgId;
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

	public Boolean getIsDefault() {
		return this.isDefault;
	}

	public void setIsDefault(Boolean isDefault) {
        this.isDefault = isDefault;
	}
}
