package com.ngtesting.platform.model;

public class TstCasePriority extends BaseModel {
	private static final long serialVersionUID = -2071266644244632484L;

	private String label;
	private String value;

	private String descr;
    private String priority;
    private Boolean defaultVal;
    private Integer ordr;
    private Integer orgId;

	public Boolean getDefaultVal() {
		return defaultVal;
	}

	public void setDefaultVal(Boolean defaultVal) {
		this.defaultVal = defaultVal;
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

	public String getPriority() {
		return priority;
	}
	public void setPriority(String priority) {
		this.priority = priority;
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
