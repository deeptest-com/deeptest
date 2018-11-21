package com.ngtesting.platform.model;

public class IsuWorkflowSolution extends BaseModel {

    private String name;
    private Integer orgId;

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
}
