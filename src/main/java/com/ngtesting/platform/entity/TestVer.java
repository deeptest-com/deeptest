package com.ngtesting.platform.entity;

import javax.persistence.*;
import java.util.Date;

@Entity
@Table(name = "tst_ver")
public class TestVer extends BaseEntity {

    private static final long serialVersionUID = 7260005873110268288L;
    private String name;

    @Enumerated(EnumType.STRING)
    private VerStatus status = VerStatus.in_progress;

    protected Date startTime;
    protected Date endTime;

	@Column(name = "descr", length = 1000)
    private String descr;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "project_id", insertable = false, updatable = false)
    private TestProject project;

    @Column(name = "project_id")
    private Long projectId;

    public static enum VerStatus {
        in_progress("in_progress"),
        end("end");

        VerStatus(String val) {
            this.val = val;
        }

        private String val;
        public String toString() {
            return val;
        }
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public VerStatus getStatus() {
        return status;
    }

    public void setStatus(VerStatus status) {
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

    public TestProject getProject() {
        return project;
    }

    public void setProject(TestProject project) {
        this.project = project;
    }

    public Long getProjectId() {
        return projectId;
    }

    public void setProjectId(Long projectId) {
        this.projectId = projectId;
    }
}
