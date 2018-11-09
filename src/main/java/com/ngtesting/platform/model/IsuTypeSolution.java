package com.ngtesting.platform.model;

public class IsuTypeSolution extends BaseModel {
    private static final long serialVersionUID = 7219897125483066915L;

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
