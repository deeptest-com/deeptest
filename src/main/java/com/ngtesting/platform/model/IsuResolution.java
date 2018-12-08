package com.ngtesting.platform.model;

public class IsuResolution extends BaseModel {
	private static final long serialVersionUID = 4118180732729567467L;

	private String label;
	private String value;

	private String descr;
	private Boolean defaultVal = false;
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

	public Boolean getIsDefault() {
		return this.defaultVal;
	}

	public void setIsDefault(Boolean defaultVal) {
		this.defaultVal = defaultVal;
	}

	public Integer getOrdr() {
		return ordr;
	}

	public void setOrdr(Integer ordr) {
		this.ordr = ordr;
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

}
