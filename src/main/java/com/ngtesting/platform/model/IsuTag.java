package com.ngtesting.platform.model;


public class IsuTag extends BaseModel {
	private static final long serialVersionUID = 6038925588573971801L;

	private String name;
    private Integer orgId;
    private Integer userId;
    private String type = "tag";

	public IsuTag() {
	}
	public IsuTag(String name, Integer orgId, Integer userId) {
		this.name = name;
		this.orgId = orgId;
        this.userId = userId;
	}

    public String getType() {
        return type;
    }

    public void setType(String type) {
        this.type = type;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public Integer getOrgId() {
        return orgId;
    }

    public void setOrgId(Integer orgId) {
        this.orgId = orgId;
    }

    public Integer getUserId() {
        return userId;
    }

    public void setUserId(Integer userId) {
        this.userId = userId;
    }
}
