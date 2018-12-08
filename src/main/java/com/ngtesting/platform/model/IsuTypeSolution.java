package com.ngtesting.platform.model;

import java.util.LinkedList;
import java.util.List;

public class IsuTypeSolution extends BaseModel {
    private static final long serialVersionUID = 7219897125483066915L;

    private String name;

    private String descr;
    private Boolean defaultVal = false;
    private Boolean buildIn = false;
    private Integer orgId;

    private List<IsuType> items = new LinkedList<>();

    public IsuTypeSolution() {

    }
    public IsuTypeSolution(String name) {
        this.name = name;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public Boolean getDefaultVal() {
        return defaultVal;
    }

    public void setDefaultVal(Boolean aDefault) {
        defaultVal = aDefault;
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

    public List<IsuType> getItems() {
        return items;
    }

    public void setItems(List<IsuType> items) {
        this.items = items;
    }

    public Boolean getBuildIn() {
        return buildIn;
    }

    public void setBuildIn(Boolean buildIn) {
      this.buildIn = buildIn;
    }
}
