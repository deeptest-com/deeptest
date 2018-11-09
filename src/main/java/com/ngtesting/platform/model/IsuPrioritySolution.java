package com.ngtesting.platform.model;

public class IsuPrioritySolution extends BaseModel {

	private static final long serialVersionUID = 1038763093908219933L;
	private String name;

	private String descr;
	private Boolean isDefault = false;
	private Integer orgId;

	public String getName() {
		return name;
	}

	public void setName(String name) {
		this.name = name;
	}

	public Boolean getDefault() {
		return isDefault;
	}

	public void setDefault(Boolean aDefault) {
		isDefault = aDefault;
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
