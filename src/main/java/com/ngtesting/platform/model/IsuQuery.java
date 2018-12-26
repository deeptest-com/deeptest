package com.ngtesting.platform.model;

import java.util.Date;

public class IsuQuery extends BaseModel {

	private static final long serialVersionUID = 5819657933820222077L;
	private String name;
	private String rule;
	private String orderBy;
	private String descr;
    private Date useTime;

//    private Integer orgId;
	private Integer projectId;
	private Integer userId;

	public String getRule() {
		return rule;
	}

	public void setRule(String rule) {
		this.rule = rule;
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

//	public Integer getOrgId() {
//		return orgId;
//	}
//
//	public void setOrgId(Integer orgId) {
//		this.orgId = orgId;
//	}

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

    public Date getUseTime() {
        return useTime;
    }

    public void setUseTime(Date useTime) {
        this.useTime = useTime;
    }

	public String getOrderBy() {
		return orderBy;
	}
	public void setOrderBy(String orderBy) {
		this.orderBy = orderBy;
	}
}
