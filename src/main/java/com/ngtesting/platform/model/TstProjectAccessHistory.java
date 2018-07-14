package com.ngtesting.platform.model;

import java.util.Date;


public class TstProjectAccessHistory extends BaseModel {
	private static final Long serialVersionUID = -5963995629184890598L;

	private Date lastAccessTime;

	private Integer userId;

    private Integer projectId;
    private String projectName;


    public Date getLastAccessTime() {
		return lastAccessTime;
	}

	public void setLastAccessTime(Date lastAccessTime) {
		this.lastAccessTime = lastAccessTime;
	}

	public Integer getUserId() {
		return userId;
	}

	public void setUserId(Integer userId) {
		this.userId = userId;
	}

	public Integer getProjectId() {
		return projectId;
	}

	public void setProjectId(Integer projectId) {
		this.projectId = projectId;
	}

	public String getProjectName() {
		return projectName;
	}

	public void setProjectName(String projectName) {
		this.projectName = projectName;
	}

}
