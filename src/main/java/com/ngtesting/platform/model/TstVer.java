package com.ngtesting.platform.model;

import java.util.Date;

public class TstVer extends BaseModel {
    private static final long serialVersionUID = 6989572574245991239L;

    private String label;

    protected Date startTime;
    protected Date endTime;

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

    public Date getStartTime() {
        return startTime;
    }

    public void setStartTime(Date startTime) {
        this.startTime = startTime;
    }

    public Date getEndTime() {
        return endTime;
    }

    public void setEndTime(Date endTime) {
        this.endTime = endTime;
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
