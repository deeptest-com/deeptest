package com.ngtesting.platform.model;

import java.util.LinkedList;
import java.util.List;

public class IsuPage extends BaseModel {
    private static final long serialVersionUID = 8715160448820762150L;

    private String name;

    private Integer orgId;

    List<IsuPageTab> tabs = new LinkedList();

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

    public List<IsuPageTab> getTabs() {
        return tabs;
    }

    public void setTabs(List<IsuPageTab> tabs) {
        this.tabs = tabs;
    }
}
