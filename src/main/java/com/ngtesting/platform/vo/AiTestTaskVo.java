package com.ngtesting.platform.vo;


import java.util.LinkedList;
import java.util.List;

public class AiTestTaskVo extends BaseVo {
    private static final long serialVersionUID = 1853368828266553643L;

    private String name;
	private String descr;

	private Integer displayOrder;

	private Long testProductId;
	private String testType;
	private String testEnv;
	private Integer testDuration;
    private Integer testConcurrent;
	private String productBranch;
	private String asrLangModel;
    private String audioType;
    private Boolean fuse;
    private String regexInput;
    private Integer startIndex;
    private Integer numbToRun;

    private Integer ordr;
    private Long pId;
    private Boolean isLeaf;

    private String key;

    private String testsetSrc;
    private Long testsetId;
    private String testsetName;
    private String testsetPath;

    private Long projectId;
    private Long createById;

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

    public void setTestsetSrc(String testsetSrc) {
        this.testsetSrc = testsetSrc;
    }

    public Integer getTestConcurrent() {
        return testConcurrent;
    }

    public void setTestConcurrent(Integer testConcurrent) {
        this.testConcurrent = testConcurrent;
    }

    public String getAudioType() {
        return audioType;
    }

    public void setAudioType(String audioType) {
        this.audioType = audioType;
    }

    public Boolean getFuse() {
        return fuse;
    }

    public void setFuse(Boolean fuse) {
        this.fuse = fuse;
    }

    public String getTestsetPath() {
        return testsetPath;
    }

    public void setTestsetPath(String testsetPath) {
        this.testsetPath = testsetPath;
    }

    private List<AiTestTaskVo> children = new LinkedList<>();

    public Long getTestsetId() {
        return testsetId;
    }

    public void setTestsetId(Long testsetId) {
        this.testsetId = testsetId;
    }

    public String getTestsetName() {
        return testsetName;
    }

    public void setTestsetName(String testsetName) {
        this.testsetName = testsetName;
    }

    public Long getProjectId() {
        return projectId;
    }

    public void setProjectId(Long projectId) {
        this.projectId = projectId;
    }

    public Long getCreateById() {
        return createById;
    }

    public void setCreateById(Long createById) {
        this.createById = createById;
    }

    public List<AiTestTaskVo> getChildren() {
        return children;
    }

    public void setChildren(List<AiTestTaskVo> children) {
        this.children = children;
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

	public Integer getDisplayOrder() {
		return displayOrder;
	}

	public void setDisplayOrder(Integer displayOrder) {
		this.displayOrder = displayOrder;
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

}
