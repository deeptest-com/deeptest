package com.ngtesting.platform.entity;

import javax.persistence.*;
import java.util.LinkedList;
import java.util.List;

@Entity
@Table(name = "tst_case_in_run")
public class TestCaseInRun extends BaseEntity {
    private static final long serialVersionUID = -2393416384079250976L;

    private String name;

    @Enumerated(EnumType.STRING)
    private TestCase.CasePriority priority;

    @Enumerated(EnumType.STRING)
    private TestCase.CaseType type;

    private Integer estimate;

    @Column(name = "objective", length = 1000)
    private String objective;

    @Column(name = "descr", length = 1000)
    private String descr;

    private Integer ordr;

    @Column(name = "pId")
    private Long pId;

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

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "case_id", insertable = false, updatable = false)
    private TestCase testCase;

    @Column(name = "case_id")
    private Long caseId;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "run_id", insertable = false, updatable = false)
    private TestRun run;

    @Column(name = "run_id")
    private Long runId;

    @Enumerated(EnumType.STRING)
    private CaseStatusInRun status = CaseStatusInRun.untest;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "create_by_id", insertable = false, updatable = false)
    private TestUser createBy;

    @Column(name = "create_by_id")
    private Long createById;

    @OneToMany(mappedBy="testCaseInRun", fetch=FetchType.LAZY, cascade=CascadeType.ALL)
    @OrderBy("ordr")
    private List<TestCaseStepInRun> steps = new LinkedList<>();

    public static enum CaseStatusInRun {
        untest("untest"),
        pass("pass"),
        fail("fail"),
        block("block");

        CaseStatusInRun(String val) {
            this.val = val;
        }

        private String val;
        public String toString() {
            return val;
        }
    }

    public Long getpId() {
        return pId;
    }

    public void setpId(Long pId) {
        this.pId = pId;
    }

    public List<TestCaseStepInRun> getSteps() {
        return steps;
    }

    public void setSteps(List<TestCaseStepInRun> steps) {
        this.steps = steps;
    }

    public CaseStatusInRun getStatus() {
        return status;
    }

    public void setStatus(CaseStatusInRun status) {
        this.status = status;
    }

    public TestRun getRun() {
        return run;
    }

    public void setRun(TestRun run) {
        this.run = run;
    }

    public Long getRunId() {
        return runId;
    }

    public void setRunId(Long runId) {
        this.runId = runId;
    }

    public TestCase getTestCase() {
        return testCase;
    }

    public void setTestCase(TestCase testCase) {
        this.testCase = testCase;
    }

    public Long getCaseId() {
        return caseId;
    }

    public void setCaseId(Long caseId) {
        this.caseId = caseId;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public TestCase.CasePriority getPriority() {
        return priority;
    }

    public void setPriority(TestCase.CasePriority priority) {
        this.priority = priority;
    }

    public TestCase.CaseType getType() {
        return type;
    }

    public void setType(TestCase.CaseType type) {
        this.type = type;
    }

    public Integer getEstimate() {
        return estimate;
    }

    public void setEstimate(Integer estimate) {
        this.estimate = estimate;
    }

    public String getObjective() {
        return objective;
    }

    public void setObjective(String objective) {
        this.objective = objective;
    }

    public String getDescr() {
        return descr;
    }

    public void setDescr(String descr) {
        this.descr = descr;
    }

    public Integer getOrdr() {
        return ordr;
    }

    public void setOrdr(Integer ordr) {
        this.ordr = ordr;
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

    public TestUser getCreateBy() {
        return createBy;
    }

    public void setCreateBy(TestUser createBy) {
        this.createBy = createBy;
    }

    public Long getCreateById() {
        return createById;
    }

    public void setCreateById(Long createById) {
        this.createById = createById;
    }

}
