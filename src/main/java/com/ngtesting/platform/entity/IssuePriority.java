package com.ngtesting.platform.entity;

import javax.persistence.*;
import java.util.Date;

@Entity
@Table(name = "isu_priority")
public class IssuePriority extends BaseEntity {

    private static final long serialVersionUID = 7260005873110268288L;
    private String name;

    protected Date startTime;
    protected Date endTime;

	@Column(name = "descr", length = 1000)
    private String descr;

    private Integer displayOrder;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "org_id", insertable = false, updatable = false)
    private TestOrg org;

    @Column(name = "org_id")
    private Long orgId;

    public Integer getDisplayOrder() {
        return displayOrder;
    }

    public void setDisplayOrder(Integer displayOrder) {
        this.displayOrder = displayOrder;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
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

    public TestOrg getOrg() {
        return org;
    }

    public void setOrg(TestOrg org) {
        this.org = org;
    }

    public Long getOrgId() {
        return orgId;
    }

    public void setOrgId(Long orgId) {
        this.orgId = orgId;
    }
}
