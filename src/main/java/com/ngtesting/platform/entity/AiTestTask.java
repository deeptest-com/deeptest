package com.ngtesting.platform.entity;

import com.ngtesting.platform.config.Constant;

import javax.persistence.*;

@Entity
@Table(name = "ai_test_task")
public class AiTestTask extends BaseEntity {
    private static final long serialVersionUID = -4545071056199720955L;
    private String name;

	@Column(name = "descr", length = 1000)
    private String descr;

    private Integer displayOrder;

    private Long testProductId;
    private String testType = "asr";
    private String testEnv = "test";
    private Integer testDuration;
    private Integer testConcurrent = 1;
    private String productBranch = "autoTest";
    private String asrLangModel = "comm";
    private String audioType = "wav";
    private Boolean fuse = false;

    private String testsetName;
    private String testsetPath;

    private String regexInput;
    private Integer startIndex;
    private Integer numbToRun;
    @Column(name = "mlfs", length = 10000)
    private String mlfs;

    private Integer ordr;
    private Long pId;
    private Boolean isLeaf;

    @Transient
    private String key;

    private String testsetSrc = "upload";

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "testset_id", insertable = false, updatable = false)
    private AiTestSet testset;

    @Column(name = "testset_id")
    private Long testsetId;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "project_id", insertable = false, updatable = false)
    private TestProject project;

    @Column(name = "project_id")
    private Long projectId;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "create_by_id", insertable = false, updatable = false)
    private TestUser createBy;

    @Column(name = "create_by_id")
    private Long createById;

    public String getMlfs() {
        return mlfs;
    }

    public void setMlfs(String mlfs) {
        this.mlfs = mlfs;
    }

    public String getRegexInput() {
        return regexInput;
    }

    public void setRegexInput(String regexInput) {
        this.regexInput = regexInput;
    }

    public Integer getStartIndex() {
        return startIndex;
    }

    public void setStartIndex(Integer startIndex) {
        this.startIndex = startIndex;
    }

    public Integer getNumbToRun() {
        return numbToRun;
    }

    public void setNumbToRun(Integer numbToRun) {
        this.numbToRun = numbToRun;
    }

    public String getTestsetSrc() {
        return testsetSrc;
    }

    public String getTestsetName() {
        return testsetName;
    }

    public void setTestsetName(String testsetName) {
        this.testsetName = testsetName;
    }

    public String getTestsetPath() {
        return testsetPath;
    }

    public void setTestsetPath(String testsetPath) {
        this.testsetPath = testsetPath;
    }

    public void setTestsetSrc(String testsetSrc) {
        this.testsetSrc = testsetSrc;
    }

    public Integer getTestConcurrent() {
        return testConcurrent;
    }

    public void setTestConcurrent(Integer testConcurrent) {
        this.testConcurrent = testConcurrent;
    }

    public Boolean getFuse() {
        return fuse;
    }

    public void setFuse(Boolean fuse) {
        this.fuse = fuse;
    }

    public String getAudioType() {
        return audioType;
    }

    public void setAudioType(String audioType) {
        this.audioType = audioType;
    }

    public String getKey() {
        return Constant.KEY_TESTCASE_DESIGN + getId();
    }

    public void setKey(String key) {
        this.key = key;
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

    public Integer getOrdr() {
        return ordr;
    }

    public void setOrdr(Integer ordr) {
        this.ordr = ordr;
    }

    public Long getpId() {
        return pId;
    }

    public void setpId(Long pId) {
        this.pId = pId;
    }

    public Boolean getLeaf() {
        return isLeaf;
    }

    public void setLeaf(Boolean leaf) {
        isLeaf = leaf;
    }

    public AiTestSet getTestset() {
        return testset;
    }

    public void setTestset(AiTestSet testset) {
        this.testset = testset;
    }

    public Long getTestsetId() {
        return testsetId;
    }

    public void setTestsetId(Long testsetId) {
        this.testsetId = testsetId;
    }

    public Long getTestProductId() {
        return testProductId;
    }

    public void setTestProductId(Long testProductId) {
        this.testProductId = testProductId;
    }

    public String getTestType() {
        return testType;
    }

    public void setTestType(String testType) {
        this.testType = testType;
    }

    public String getTestEnv() {
        return testEnv;
    }

    public void setTestEnv(String testEnv) {
        this.testEnv = testEnv;
    }

    public Integer getTestDuration() {
        return testDuration;
    }

    public void setTestDuration(Integer testDuration) {
        this.testDuration = testDuration;
    }

    public String getProductBranch() {
        return productBranch;
    }

    public void setProductBranch(String productBranch) {
        this.productBranch = productBranch;
    }

    public String getAsrLangModel() {
        return asrLangModel;
    }

    public void setAsrLangModel(String asrLangModel) {
        this.asrLangModel = asrLangModel;
    }

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

    public String getDescr() {
        return descr;
    }

    public void setDescr(String descr) {
        this.descr = descr;
    }

}
