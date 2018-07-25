package com.ngtesting.platform.model;

public class TstCasePriority extends BaseModel {
	private static final long serialVersionUID = -2071266644244632484L;

	private String code;
	private String name;
	private String descr;
    private String priority;
    private Boolean isDefault;
    private Integer ordr;
    private Integer orgId;

	public String getCode() {
		return code;
	}
	public void setCode(String code) {
		this.code = code;
	}
	public String getName() {
		return name;
	}
	public void setName(String name) {
		this.name = name;
	}
	public String getPriority() {
		return priority;
	}
	public void setPriority(String priority) {
		this.priority = priority;
	}

	public Boolean getDefault() {
		return isDefault;
	}

	public void setDefault(Boolean aDefault) {
		isDefault = aDefault;
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
