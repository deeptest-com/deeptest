package com.ngtesting.platform.model;

import java.util.LinkedList;
import java.util.List;

public class TstSuite extends BaseModel {
	private static final long serialVersionUID = 2718812662455394471L;

	private String name;
	private Integer estimate;
	private String descr;

	private Integer projectId;
	private String projectName;
    private Integer caseProjectId;
    private String caseProjectName;
	private Integer userId;
    private String userName;
    private Boolean selecting;

    Integer count;

	private List<TstCaseInSuite> testCases = new LinkedList();

    public Integer getCaseProjectId() {
        return caseProjectId;
    }

    public void setCaseProjectId(Integer caseProjectId) {
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

	public Integer getProjectId() {
		return projectId;
	}

	public void setProjectId(Integer projectId) {
		this.projectId = projectId;
	}

	public Integer getUserId() {
		return userId;
	}

	public void setUserId(Integer userId) {
		this.userId = userId;
	}

	public List<TstCaseInSuite> getTestCases() {
		return testCases;
	}

	public void setTestCases(List<TstCaseInSuite> testCases) {
		this.testCases = testCases;
	}
}
