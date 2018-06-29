package com.ngtesting.platform.model;

import java.util.LinkedList;
import java.util.List;

public class TstSuite extends BaseModel {
	private static final long serialVersionUID = 2718812662455394471L;

	private String name;
	private Integer estimate;
	private String descr;

	private Long projectId;
	private String projectName;
    private Long caseProjectId;
    private String caseProjectName;
	private Long userId;
    private String userName;
    private Boolean selecting;

    Integer count;

	private List<TstCaseInSuite> testcases = new LinkedList();
    public TstSuite(){}
    public TstSuite(Integer id, String name, Integer estimate, String descr, Long projectId, Long userId) {
        super();

		this.id = id;
        this.name = name;
        this.estimate = estimate;
        this.descr = descr;
        this.projectId = projectId;
        this.userId = userId;
    }

    public Long getCaseProjectId() {
        return caseProjectId;
    }

    public void setCaseProjectId(Long caseProjectId) {
        this.caseProjectId = caseProjectId;
    }


    public String getCaseProjectName() {
        return caseProjectName;
    }

    public void setCaseProjectName(String caseProjectName) {
        this.caseProjectName = caseProjectName;
    }

    public String getProjectName() {
		return projectName;
	}
	public void setProjectName(String projectName) {
		this.projectName = projectName;
	}

	public Boolean getSelecting() {
        return selecting;
    }
    public void setSelecting(Boolean selecting) {
        this.selecting = selecting;
    }

    public Integer getCount() {
        return count;
    }

    public void setCount(Integer count) {
        this.count = count;
    }

    public String getUserName() {
        return userName;
    }

    public void setUserName(String userName) {
        this.userName = userName;
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

	public String getDescr() {
		return descr;
	}

	public void setDescr(String descr) {
		this.descr = descr;
	}

	public Long getProjectId() {
		return projectId;
	}

	public void setProjectId(Long projectId) {
		this.projectId = projectId;
	}

	public Long getUserId() {
		return userId;
	}

	public void setUserId(Long userId) {
		this.userId = userId;
	}

	public List<TstCaseInSuite> getTestcases() {
		return testcases;
	}

	public void setTestcases(List<TstCaseInSuite> testcases) {
		this.testcases = testcases;
	}
}
