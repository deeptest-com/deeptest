package com.ngtesting.platform.entity;

import org.hibernate.annotations.Where;

import javax.persistence.*;
import java.util.Date;
import java.util.LinkedList;
import java.util.List;

@Entity
@Table(name = "isu_workflow")
public class IssueWorkflow extends BaseEntity {

    private static final long serialVersionUID = 7260005873110268288L;
    private String name;

    protected Date startTime;
    protected Date endTime;

	@Column(name = "descr", length = 1000)
    private String descr;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "project_id", insertable = false, updatable = false)
    private TestProject project;

    @Column(name = "project_id")
    private Long projectId;

    @OneToMany(mappedBy="workflow", fetch=FetchType.LAZY)
    @Where(clause="!deleted")
    @OrderBy("ordr")
    private List<IssueStatusTransition> statusTansitions = new LinkedList<>();

    public List<IssueStatusTransition> getStatusTansitions() {
        return statusTansitions;
    }

    public void setStatusTansitions(List<IssueStatusTransition> statusTansitions) {
        this.statusTansitions = statusTansitions;
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
