package com.ngtesting.platform.model;

import java.util.Date;
import java.util.LinkedList;
import java.util.List;

public class TstPlan extends BaseModel {

    private static final long serialVersionUID = -209226082263881616L;

    private String name;
    private Integer estimate;

    private PlanStatus status;

    protected Date startTime;

    protected Date endTime;

    private String descr;

    private Long projectId;
    private String projectName;

    private Long verId;
    private String verName;

    private List<TstTask> runs = new LinkedList<>();

    public static enum PlanStatus {
        not_start("not_start"),
        in_progress("in_progress"),
        end("end");

        PlanStatus(String val) {
            this.val = val;
        }

        private String val;
        public String toString() {
            return val;
        }
    }

    public String getProjectName() {
        return projectName;
    }

    public void setProjectName(String projectName) {
        this.projectName = projectName;
    }

    public String getVerName() {
        return verName;
    }

    public void setVerName(String verName) {
        this.verName = verName;
    }

    public Long getVerId() {
        return verId;
    }

    public void setVerId(Long verId) {
        this.verId = verId;
    }

    public List<TstTask> getRuns() {
        return runs;
    }

    public void setRuns(List<TstTask> runs) {
        this.runs = runs;
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

    public PlanStatus getStatus() {
        return status;
    }

    public void setStatus(PlanStatus status) {
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
