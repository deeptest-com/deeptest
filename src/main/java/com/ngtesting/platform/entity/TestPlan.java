package com.ngtesting.platform.entity;

import org.hibernate.annotations.Where;

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
    private PlanStatus status = PlanStatus.not_start;

    protected Date startTime;

    protected Date endTime;

	@Column(name = "descr", length = 1000)
    private String descr;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "project_id", insertable = false, updatable = false)
    private TestProject project;
    @Column(name = "project_id")
    private Long projectId;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "ver_id", insertable = false, updatable = false)
    private TestVer ver;
    @Column(name = "ver_id")
    private Long verId;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "env_id", insertable = false, updatable = false)
    private TestVer env;
    @Column(name = "env_id")
    private Long envId;

    @OneToMany(mappedBy="plan", fetch=FetchType.LAZY, cascade=CascadeType.ALL)
    @OrderBy("id")
    @Where(clause="!deleted")
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

    public TestVer getVer() {
        return ver;
    }

    public void setVer(TestVer ver) {
        this.ver = ver;
    }

    public Long getVerId() {
        return verId;
    }

    public void setVerId(Long verId) {
        this.verId = verId;
    }

    public TestVer getEnv() {
        return env;
    }

    public void setEnv(TestVer env) {
        this.env = env;
    }

    public Long getEnvId() {
        return envId;
    }

    public void setEnvId(Long envId) {
        this.envId = envId;
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
