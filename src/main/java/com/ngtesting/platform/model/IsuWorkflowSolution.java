package com.ngtesting.platform.model;

public class IsuWorkflowSolution extends BaseModel {

    private String name;
    private Integer orgId;

    private Boolean defaultVal = false;
    private Boolean buildIn = false;

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
