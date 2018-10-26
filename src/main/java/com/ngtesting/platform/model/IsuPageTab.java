package com.ngtesting.platform.model;

import java.util.LinkedList;
import java.util.List;

public class IsuPageTab extends BaseModel {
    private static final long serialVersionUID = -3446493515026238121L;

    private String name;
    private String descr;
    private Integer pageId;

    private Integer ordr;
    private Integer orgId;

    List<IsuPageElement> elements = new LinkedList();

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
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

    public List<IsuPageElement> getElements() {
        return elements;
    }

    public void setElements(List<IsuPageElement> elements) {
        this.elements = elements;
    }

    public Integer getPageId() {
        return pageId;
    }

    public void setPageId(Integer pageId) {
        this.pageId = pageId;
    }
}
