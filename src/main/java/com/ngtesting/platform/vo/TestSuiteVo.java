package com.ngtesting.platform.vo;

import java.util.LinkedList;
import java.util.List;

public class TestSuiteVo extends BaseVo {

	private static final long serialVersionUID = 2718812662455394471L;

	private String name;
	private Integer estimate;
	private String descr;

	private Long projectId;
	private Long userId;
    private String userName;

    Integer count;

	private List<TestCaseInSuiteVo> testcases = new LinkedList();
    public TestSuiteVo(){}
    public TestSuiteVo(Long id, String name, Integer estimate, String descr, Long projectId, Long userId) {
        super();

		this.id = id;
        this.name = name;
        this.estimate = estimate;
        this.descr = descr;
        this.projectId = projectId;
        this.userId = userId;
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

	public List<TestCaseInSuiteVo> getTestcases() {
		return testcases;
	}

	public void setTestcases(List<TestCaseInSuiteVo> testcases) {
		this.testcases = testcases;
	}
}
