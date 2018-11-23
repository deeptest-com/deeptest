package com.ngtesting.platform.model;

public class TstEnv extends BaseModel {
    private static final long serialVersionUID = -8999964583887292845L;
    private String label;
    private String descr;

    private Integer projectId;
    private Integer orgId;

    private Integer ordr;

    public String getLabel() {
        return label;
    }

    public void setLabel(String label) {
        this.label = label;
    }

    public String getDescr() {
        return descr;
    }

    public void setDescr(String descr) {
        this.descr = descr;
    }

    public Integer getProjectId() {
        return projectId;
    }

    public void setProjectId(Integer projectId) {
        this.projectId = projectId;
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
}
