package com.ngtesting.platform.model;

import java.util.LinkedList;
import java.util.List;

public class IsuPage extends BaseModel {
    private static final long serialVersionUID = 8715160448820762150L;

    private String name;
    private Integer orgId;
    private Boolean defaultVal = false;
    private Boolean buildIn = false;

    List<IsuPageElement> elements = new LinkedList();

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

    public List<IsuPageElement> getElements() {
        return elements;
    }

    public void setElements(List<IsuPageElement> elements) {
        this.elements = elements;
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
