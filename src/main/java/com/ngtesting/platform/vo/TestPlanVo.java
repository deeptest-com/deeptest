package com.ngtesting.platform.vo;

import java.util.Date;
import java.util.LinkedList;
import java.util.List;

public class TestPlanVo extends BaseVo {

    private static final long serialVersionUID = -209226082263881616L;

    private String name;
    private Integer estimate;

    private String status;

    protected Date startTime;

    protected Date endTime;

    private String descr;

    private Long projectId;

    private List<TestRunVo> runVos = new LinkedList<>();

    public List<TestRunVo> getRunVos() {
        return runVos;
    }

    public void setRunVos(List<TestRunVo> runVos) {
        this.runVos = runVos;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public Integer getEstimate() {
        return estimate;
    }

    public void setEstimate(Integer estimate) {
        this.estimate = estimate;
    }

    public String getStatus() {
        return status;
    }

    public void setStatus(String status) {
        this.status = status;
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

    public Long getProjectId() {
        return projectId;
    }

    public void setProjectId(Long projectId) {
        this.projectId = projectId;
    }

}
