package com.ngtesting.platform.entity;

import javax.persistence.*;
import java.util.Date;

@Entity
@Table(name = "tst_alert")
public class TestAlert extends BaseEntity {
    private static final long serialVersionUID = 4639102366457159222L;

    private Long entityId;
    private String entityName;

    @Column(name = "descr", length = 10000)
    private String descr;
    private String uri;
    private String type;
    private String status;

    private Date startTime;
    private Date endTime;

    private Boolean isRead = false;
    private Boolean isSent = false;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "assignee_id", insertable = false, updatable = false)
    private TestUser assignee;

    @Column(name = "assignee_id")
    private Long assigneeId;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "user_id", insertable = false, updatable = false)
    private TestUser user;

    @Column(name = "user_id")
    private Long userId;

    public String getStatus() {
        return status;
    }

    public void setStatus(String status) {
        this.status = status;
    }

    public Long getEntityId() {
        return entityId;
    }

    public void setEntityId(Long entityId) {
        this.entityId = entityId;
    }

    public String getEntityName() {
        return entityName;
    }

    public void setEntityName(String entityName) {
        this.entityName = entityName;
    }

    public String getDescr() {
        return descr;
    }

    public void setDescr(String descr) {
        this.descr = descr;
    }

    public String getUri() {
        return uri;
    }

    public void setUri(String uri) {
        this.uri = uri;
    }

    public String getType() {
        return type;
    }

    public void setType(String type) {
        this.type = type;
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

    public Boolean getRead() {
        return isRead;
    }

    public void setRead(Boolean read) {
        isRead = read;
    }

    public Boolean getSent() {
        return isSent;
    }

    public void setSent(Boolean sent) {
        isSent = sent;
    }

    public TestUser getAssignee() {
        return assignee;
    }

    public void setAssignee(TestUser assignee) {
        this.assignee = assignee;
    }

    public Long getAssigneeId() {
        return assigneeId;
    }

    public void setAssigneeId(Long assigneeId) {
        this.assigneeId = assigneeId;
    }

    public TestUser getUser() {
        return user;
    }

    public void setUser(TestUser user) {
        this.user = user;
    }

    public Long getUserId() {
        return userId;
    }

    public void setUserId(Long userId) {
        this.userId = userId;
    }
}
