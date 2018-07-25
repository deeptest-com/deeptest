package com.ngtesting.platform.model;

public class TstCaseType extends BaseModel {
	private static final long serialVersionUID = -2071266644244632484L;

	private String code;
	private String name;
	private String descr;
    private Boolean isDefault = false;
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
