package com.ngtesting.platform.vo;


public class OrgPriviledgeVo extends BaseVo {
	private static final long serialVersionUID = -6981838223153872057L;
	
	private String name;
    private String descr;
    
    private Long orgId;
    
    private Boolean selecting;
    private Boolean selected;
    
    public OrgPriviledgeVo() {
	}

	public OrgPriviledgeVo(Long id, String name, String descr, Long orgId) {
		this.id = id;
		this.name = name;
		this.descr = descr;
		this.orgId = orgId;
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
    
}
