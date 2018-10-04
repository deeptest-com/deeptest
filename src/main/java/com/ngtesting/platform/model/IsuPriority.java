package com.ngtesting.platform.model;

public class IsuPriority extends BaseModel {

	private static final long serialVersionUID = 4118180732729567467L;

	private String label;
	private String code;

	private String descr;
    private Boolean isDefault;
    private Integer displayOrder;
    private Integer orgId;

	public String getCode() {
		return code;
	}

	public void setCode(String code) {
		this.code = code;
	}

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

	public Boolean getIsDefault() {
		return isDefault;
	}
	public void setIsDefault(Boolean isDefault) {
		this.isDefault = isDefault;
	}
	public Integer getDisplayOrder() {
		return displayOrder;
	}
	public void setDisplayOrder(Integer displayOrder) {
		this.displayOrder = displayOrder;
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
