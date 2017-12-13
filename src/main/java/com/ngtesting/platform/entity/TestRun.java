package com.ngtesting.platform.entity;

import javax.persistence.*;
import java.util.LinkedList;
import java.util.List;

@Entity
@Table(name = "tst_run")
public class TestRun extends BaseEntity {

    private String name;
	private Integer estimate;

    @Enumerated(EnumType.STRING)
    private RunStatus status;

	@Column(name = "descr", length = 1000)
    private String descr;

    private Integer ordr;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "project_id", insertable = false, updatable = false)
    private TestProject project;

    @Column(name = "project_id")
    private Long projectId;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "user_id", insertable = false, updatable = false)
    private TestUser user;

    @Column(name = "user_id")
    private Long userId;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "plan_id", insertable = false, updatable = false)
    private TestPlan plan;

    @Column(name = "plan_id")
    private Long planId;

    @OneToMany(mappedBy="run", cascade={CascadeType.ALL}, fetch=FetchType.LAZY)
    private List<TestCaseInRun> testcases = new LinkedList<>();

    public static enum RunStatus {
        not_start("not_start"),
        in_progress("in_progress"),
        end("end");

        RunStatus(String val) {
            this.val = val;
        }

        private String val;
        public String toString() {
            return val;
        }
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

    public List<TestCaseInRun> getTestcases() {
        return testcases;
    }

    public void setTestcases(List<TestCaseInRun> testcases) {
        this.testcases = testcases;
    }

    public Integer getOrdr() {
        return ordr;
    }

    public void setOrdr(Integer ordr) {
        this.ordr = ordr;
    }

    public TestPlan getPlan() {
        return plan;
    }

    public void setPlan(TestPlan plan) {
        this.plan = plan;
    }

    public Long getPlanId() {
        return planId;
    }

    public void setPlanId(Long planId) {
        this.planId = planId;
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

    public RunStatus getStatus() {
        return status;
    }

    public void setStatus(RunStatus status) {
        this.status = status;
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
