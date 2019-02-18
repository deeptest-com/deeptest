package com.ngtesting.platform.model;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;

import java.util.Date;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;

public class IsuIssue extends BaseModel {
	private static final long serialVersionUID = 690125164201934534L;

    private String title;
    private String descr;

    private Integer projectId;
    private String projectName;

    private Integer typeId;
    private String typeName;

    private Integer statusId;
    private String statusName;

    private Integer priorityId;
    private String priorityName;

    private Integer verId;
    private String verName;

    private Integer envId;
    private String envName;

    private Integer resolutionId;
    private String resolutionName;

    private String resolutionDescr;

    private Date dueTime;
    private Date resolveTime;
    private Date setFinalTime;

    private Integer creatorId;
    private String creatorName;
    private Integer reporterId;
    private String reporterName;
    private Integer assigneeId;
    private String assigneeName;

    private Boolean watched = false;

    private Integer orgId;
    private Integer prjId;

    private String extProp;
    private String uuid;

    private List<IsuComments> comments = new LinkedList<>();
    private List<IsuAttachment> attachments = new LinkedList<>();
    private List<IsuHistory> histories = new LinkedList<>();
    private List<IsuTag> tags = new LinkedList<>();
    private List<Map> watchList = new LinkedList<>();
    private List<IsuLink> links = new LinkedList<>();

    public JSONObject getJsonProp() {
        return JSON.parseObject(extProp);
    }
    public void setExtProp(String extProp) {
        this.extProp = extProp;
    }

    public String getDescr() {
        return descr;
    }

    public void setDescr(String descr) {
        this.descr = descr;
    }

    public List<Map> getWatchList() {
        return watchList;
    }

    public void setWatchList(List<Map> watchList) {
        this.watchList = watchList;
    }

    public Date getSetFinalTime() {
        return setFinalTime;
    }

    public void setSetFinalTime(Date setFinalTime) {
        this.setFinalTime = setFinalTime;
    }

    public List<IsuLink> getLinks() {
        return links;
    }

    public void setLinks(List<IsuLink> links) {
        this.links = links;
    }

    public List<IsuTag> getTags() {
        return tags;
    }

    public void setTags(List<IsuTag> tags) {
        this.tags = tags;
    }

    public String getTitle() {
        return title;
    }

    public void setTitle(String title) {
        this.title = title;
    }

    public Integer getProjectId() {
        return projectId;
    }

    public void setProjectId(Integer projectId) {
        this.projectId = projectId;
    }

    public Integer getTypeId() {
        return typeId;
    }

    public void setTypeId(Integer typeId) {
        this.typeId = typeId;
    }

    public String getTypeName() {
        return typeName;
    }

    public void setTypeName(String typeName) {
        this.typeName = typeName;
    }

    public Integer getStatusId() {
        return statusId;
    }

    public void setStatusId(Integer statusId) {
        this.statusId = statusId;
    }

    public String getStatusName() {
        return statusName;
    }

    public void setStatusName(String statusName) {
        this.statusName = statusName;
    }

    public Integer getPriorityId() {
        return priorityId;
    }

    public void setPriorityId(Integer priorityId) {
        this.priorityId = priorityId;
    }

    public String getPriorityName() {
        return priorityName;
    }

    public void setPriorityName(String priorityName) {
        this.priorityName = priorityName;
    }

    public Integer getVerId() {
        return verId;
    }

    public void setVerId(Integer verId) {
        this.verId = verId;
    }

    public String getVerName() {
        return verName;
    }

    public void setVerName(String verName) {
        this.verName = verName;
    }

    public Integer getEnvId() {
        return envId;
    }

    public void setEnvId(Integer envId) {
        this.envId = envId;
    }

    public String getEnvName() {
        return envName;
    }

    public void setEnvName(String envName) {
        this.envName = envName;
    }

    public Integer getResolutionId() {
        return resolutionId;
    }

    public void setResolutionId(Integer resolutionId) {
        this.resolutionId = resolutionId;
    }

    public String getResolutionName() {
        return resolutionName;
    }

    public void setResolutionName(String resolutionName) {
        this.resolutionName = resolutionName;
    }

    public Date getDueTime() {
        return dueTime;
    }

    public void setDueTime(Date dueTime) {
        this.dueTime = dueTime;
    }

    public Date getResolveTime() {
        return resolveTime;
    }

    public void setResolveTime(Date resolveTime) {
        this.resolveTime = resolveTime;
    }

    public Integer getCreatorId() {
        return creatorId;
    }

    public void setCreatorId(Integer creatorId) {
        this.creatorId = creatorId;
    }

    public String getCreatorName() {
        return creatorName;
    }

    public void setCreatorName(String creatorName) {
        this.creatorName = creatorName;
    }

    public Integer getReporterId() {
        return reporterId;
    }

    public void setReporterId(Integer reporterId) {
        this.reporterId = reporterId;
    }

    public String getReporterName() {
        return reporterName;
    }

    public void setReporterName(String reporterName) {
        this.reporterName = reporterName;
    }

    public Integer getAssigneeId() {
        return assigneeId;
    }

    public void setAssigneeId(Integer assigneeId) {
        this.assigneeId = assigneeId;
    }

    public String getAssigneeName() {
        return assigneeName;
    }

    public void setAssigneeName(String assigneeName) {
        this.assigneeName = assigneeName;
    }

    public List<IsuAttachment> getAttachments() {
        return attachments;
    }

    public void setAttachments(List<IsuAttachment> attachments) {
        this.attachments = attachments;
    }

    public List<IsuComments> getComments() {
        return comments;
    }

    public void setComments(List<IsuComments> comments) {
        this.comments = comments;
    }

    public List<IsuHistory> getHistories() {
        return histories;
    }

    public void setHistories(List<IsuHistory> histories) {
        this.histories = histories;
    }

    public String getProjectName() {
        return projectName;
    }

    public void setProjectName(String projectName) {
        this.projectName = projectName;
    }

    public Integer getOrgId() {
        return orgId;
    }

    public void setOrgId(Integer orgId) {
        this.orgId = orgId;
    }

    public Integer getPrjId() {
        return prjId;
    }

    public void setPrjId(Integer prjId) {
        this.prjId = prjId;
    }

    public String getUuid() {
        return uuid;
    }

    public void setUuid(String uuid) {
        this.uuid = uuid;
    }

    public String getResolutionDescr() {
        return resolutionDescr;
    }

    public void setResolutionDescr(String resolutionDescr) {
        this.resolutionDescr = resolutionDescr;
    }

    public Boolean getWatched() {
        return watched;
    }

    public void setWatched(Boolean watched) {
        this.watched = watched;
    }
}
