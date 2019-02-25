package com.ngtesting.platform.model;

public class IsuType extends BaseModel {
	private static final long serialVersionUID = -2904250132388811594L;

    private String label;
    /*private String value;*/

    private String descr;
    private Boolean defaultVal = false;
    private Boolean buildIn = false;
    private Integer ordr;
    private Integer orgId;

    public String getLabel() {
        return label;
    }

    public void setLabel(String label) {
        this.label = label;
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
