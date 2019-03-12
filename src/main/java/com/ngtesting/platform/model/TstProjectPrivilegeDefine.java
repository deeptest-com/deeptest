package com.ngtesting.platform.model;


public class TstProjectPrivilegeDefine extends BaseModel {
	private static final long serialVersionUID = -6981838223153872057L;

	private String code;
	private String action;
	private String name;
    private String descr;

	private Integer relationId;
    private Integer orgId;

    private Boolean selecting;
    private Boolean selected;

    public TstProjectPrivilegeDefine() {
	}

	public TstProjectPrivilegeDefine(Integer id, String code, String action, String name, String descr,
									 Integer orgId) {
		this.id = id;
		this.code = code;
		this.action = action;
		this.name = name;
		this.descr = descr;
		this.orgId = orgId;
	}

	public Integer getRelationId() {
		return relationId;
	}

	public void setRelationId(Integer relationId) {
		this.relationId = relationId;
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

	public Integer getOrgId() {
		return orgId;
	}

	public void setOrgId(Integer orgId) {
		this.orgId = orgId;
	}

	public Boolean getSelected() {
		return selected;
	}

	public void setSelected(Boolean selected) {
		this.selected = selected;
	}

	public Boolean getSelecting() {
		return selecting;
	}

	public void setSelecting(Boolean selecting) {
		this.selecting = selecting;
	}

	public String getCode() {
		return code;
	}

	public void setCode(String code) {
		this.code = code;
	}

	public String getAction() {
		return action;
	}

	public void setAction(String action) {
		this.action = action;
	}

}
