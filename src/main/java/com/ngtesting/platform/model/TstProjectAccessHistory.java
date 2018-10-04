package com.ngtesting.platform.model;

import java.util.Date;


public class TstProjectAccessHistory extends BaseModel {
	private static final long serialVersionUID = -5963995629184890598L;

	private Date lastAccessTime;

	private Integer userId;

    private Integer prjId;
    private String prjName;


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

	public Integer getPrjId() {
		return prjId;
	}

	public void setPrjId(Integer prjId) {
		this.prjId = prjId;
	}

	public String getPrjName() {
		return prjName;
	}

	public void setPrjName(String prjName) {
		this.prjName = prjName;
	}
}
