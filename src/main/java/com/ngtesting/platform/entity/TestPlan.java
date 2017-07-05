package com.ngtesting.platform.entity;

import javax.persistence.*;
import java.util.Date;
import java.util.LinkedList;
import java.util.List;

@Entity
@Table(name = "tst_plan")
public class TestPlan extends BaseEntity {

    private static final long serialVersionUID = -2388027442087410471L;

    private String name;
	private Integer estimate;

    @Enumerated(EnumType.STRING)
    private PlanStatus status;

    @Column(insertable = true, updatable = false)
    protected Date startTime = new Date();

    @Column(insertable = true, updatable = false)
    protected Date endTime = new Date();
    
	@Column(name = "descr", length = 1000)
    private String descr;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "project_id", insertable = false, updatable = false)
    private TestProject project;

    @Column(name = "project_id")
    private Long projectId;

    @OneToMany(mappedBy="plan", fetch=FetchType.LAZY, cascade=CascadeType.ALL)
    @OrderBy("ordr")
    private List<TestRun> runs = new LinkedList<>();

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

    public List<TestRun> getRuns() {
        return runs;
    }

    public void setRuns(List<TestRun> runs) {
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
