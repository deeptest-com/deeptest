package com.ngtesting.platform.vo;


public class ProjectPrivilegeDefineVo extends BaseVo {
	private static final long serialVersionUID = -6981838223153872057L;
	
	private String code;
	private String name;
	private String action;
    private String descr;

	private Long relationId;
    private Long orgId;
    
    private Boolean selecting;
    private Boolean selected;
    
    public ProjectPrivilegeDefineVo() {
	}

	public ProjectPrivilegeDefineVo(Long id, String code, String action, String name, String descr,
                                    Long orgId) {
		this.id = id;
		this.code = code;
		this.action = action;
		this.name = name;
		this.descr = descr;
		this.orgId = orgId;
	}

	public Long getRelationId() {
		return relationId;
	}

	public void setRelationId(Long relationId) {
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

	public Long getOrgId() {
		return orgId;
	}

	public void setOrgId(Long orgId) {
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
