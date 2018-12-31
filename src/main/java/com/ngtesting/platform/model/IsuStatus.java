package com.ngtesting.platform.model;

public class IsuStatus extends BaseModel {

	private static final long serialVersionUID = 1657004478821957505L;
	private String label;
	private String value;
    private String descr;
    private Integer ordr;

    private Boolean defaultVal = false;
    private Boolean finalVal;
    private Boolean buildIn = false;

	private Integer categoryId;
	private String categoryName;
    private Integer orgId;

    private Boolean selected;

    @Override
    public boolean equals(Object obj) {
        IsuStatus s = (IsuStatus)obj;
        return this.getId().intValue() == s.getId().intValue();
    }
    @Override
    public int hashCode() {
        return this.id;
    }

    public Boolean getDefaultVal() {
        return defaultVal;
    }

    public void setDefaultVal(Boolean aDefault) {
        defaultVal = aDefault;
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

    public Boolean getFinalVal() {
        return finalVal;
    }

    public void setFinalVal(Boolean finalVal) {
        this.finalVal = finalVal;
    }

    public Boolean getBuildIn() {
		return buildIn;
	}

	public void setBuildIn(Boolean buildIn) {
    this.buildIn = buildIn;
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

	public Integer getOrgId() {
		return orgId;
	}
	public void setOrgId(Integer orgId) {
		this.orgId = orgId;
	}

    public Integer getCategoryId() {
        return categoryId;
    }

    public void setCategoryId(Integer categoryId) {
        this.categoryId = categoryId;
    }

    public String getCategoryName() {
        return categoryName;
    }

    public void setCategoryName(String categoryName) {
        this.categoryName = categoryName;
    }

    public Boolean getSelected() {
        return selected;
    }

    public void setSelected(Boolean selected) {
        this.selected = selected;
    }

}
