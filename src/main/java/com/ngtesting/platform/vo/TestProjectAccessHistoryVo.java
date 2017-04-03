package com.ngtesting.platform.vo;

import java.util.Date;


public class TestProjectAccessHistoryVo extends BaseVo {
	private static final long serialVersionUID = -5963995629184890598L;

	private Date lastAccessTime;

	private Long userId;

    private Long projectId;
    private String projectName;
    

    public Date getLastAccessTime() {
		return lastAccessTime;
	}

	public void setLastAccessTime(Date lastAccessTime) {
		this.lastAccessTime = lastAccessTime;
	}

	public Long getUserId() {
		return userId;
	}

	public void setUserId(Long userId) {
		this.userId = userId;
	}

	public Long getProjectId() {
		return projectId;
	}

	public void setProjectId(Long projectId) {
		this.projectId = projectId;
	}

	public String getProjectName() {
		return projectName;
	}

	public void setProjectName(String projectName) {
		this.projectName = projectName;
	}

}
