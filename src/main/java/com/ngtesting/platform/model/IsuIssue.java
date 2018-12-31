package com.ngtesting.platform.model;

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

    private String uuid;

    private List<IsuComments> comments = new LinkedList<>();
    private List<IsuAttachment> attachments = new LinkedList<>();
    private List<IsuHistory> histories = new LinkedList<>();
    private List<IsuTag> tags = new LinkedList<>();
    private List<Map> watchList = new LinkedList<>();
    private List<IsuLink> links = new LinkedList<>();

    private String prop01;
    private String prop02;
    private String prop03;
    private String prop04;
    private String prop05;

    private String prop06;
    private String prop07;
    private String prop08;
    private String prop09;
    private String prop10;

    private String prop11;
    private String prop12;
    private String prop13;
    private String prop14;
    private String prop15;

    private String prop16;
    private String prop17;
    private String prop18;
    private String prop19;
    private String prop20;

    private String prop21;
    private String prop22;
    private String prop23;
    private String prop24;
    private String prop25;

    private String prop26;
    private String prop27;
    private String prop28;
    private String prop29;
    private String prop30;


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

    public String getProp01() {
        return prop01;
    }

    public void setProp01(String prop01) {
        this.prop01 = prop01;
    }

    public String getProp02() {
        return prop02;
    }

    public void setProp02(String prop02) {
        this.prop02 = prop02;
    }

    public String getProp03() {
        return prop03;
    }

    public void setProp03(String prop03) {
        this.prop03 = prop03;
    }

    public String getProp04() {
        return prop04;
    }

    public void setProp04(String prop04) {
        this.prop04 = prop04;
    }

    public String getProp05() {
        return prop05;
    }

    public void setProp05(String prop05) {
        this.prop05 = prop05;
    }

    public String getProp06() {
        return prop06;
    }

    public void setProp06(String prop06) {
        this.prop06 = prop06;
    }

    public String getProp07() {
        return prop07;
    }

    public void setProp07(String prop07) {
        this.prop07 = prop07;
    }

    public String getProp08() {
        return prop08;
    }

    public void setProp08(String prop08) {
        this.prop08 = prop08;
    }

    public String getProp09() {
        return prop09;
    }

    public void setProp09(String prop09) {
        this.prop09 = prop09;
    }

    public String getProp10() {
        return prop10;
    }

    public void setProp10(String prop10) {
        this.prop10 = prop10;
    }

    public String getProp11() {
        return prop11;
    }

    public void setProp11(String prop11) {
        this.prop11 = prop11;
    }

    public String getProp12() {
        return prop12;
    }

    public void setProp12(String prop12) {
        this.prop12 = prop12;
    }

    public String getProp13() {
        return prop13;
    }

    public void setProp13(String prop13) {
        this.prop13 = prop13;
    }

    public String getProp14() {
        return prop14;
    }

    public void setProp14(String prop14) {
        this.prop14 = prop14;
    }

    public String getProp15() {
        return prop15;
    }

    public void setProp15(String prop15) {
        this.prop15 = prop15;
    }

    public String getProp16() {
        return prop16;
    }

    public void setProp16(String prop16) {
        this.prop16 = prop16;
    }

    public String getProp17() {
        return prop17;
    }

    public void setProp17(String prop17) {
        this.prop17 = prop17;
    }

    public String getProp18() {
        return prop18;
    }

    public void setProp18(String prop18) {
        this.prop18 = prop18;
    }

    public String getProp19() {
        return prop19;
    }

    public void setProp19(String prop19) {
        this.prop19 = prop19;
    }

    public String getProp20() {
        return prop20;
    }

    public void setProp20(String prop20) {
        this.prop20 = prop20;
    }

    public String getProp21() {
        return prop21;
    }

    public void setProp21(String prop21) {
        this.prop21 = prop21;
    }

    public String getProp22() {
        return prop22;
    }

    public void setProp22(String prop22) {
        this.prop22 = prop22;
    }

    public String getProp23() {
        return prop23;
    }

    public void setProp23(String prop23) {
        this.prop23 = prop23;
    }

    public String getProp24() {
        return prop24;
    }

    public void setProp24(String prop24) {
        this.prop24 = prop24;
    }

    public String getProp25() {
        return prop25;
    }

    public void setProp25(String prop25) {
        this.prop25 = prop25;
    }

    public String getProp26() {
        return prop26;
    }

    public void setProp26(String prop26) {
        this.prop26 = prop26;
    }

    public String getProp27() {
        return prop27;
    }

    public void setProp27(String prop27) {
        this.prop27 = prop27;
    }

    public String getProp28() {
        return prop28;
    }

    public void setProp28(String prop28) {
        this.prop28 = prop28;
    }

    public String getProp29() {
        return prop29;
    }

    public void setProp29(String prop29) {
        this.prop29 = prop29;
    }

    public String getProp30() {
        return prop30;
    }

    public void setProp30(String prop30) {
        this.prop30 = prop30;
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
