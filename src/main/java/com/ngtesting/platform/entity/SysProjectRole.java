package com.ngtesting.platform.entity;

import java.util.HashSet;
import java.util.Set;

import javax.persistence.CascadeType;
import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.FetchType;
import javax.persistence.JoinColumn;
import javax.persistence.JoinTable;
import javax.persistence.ManyToMany;
import javax.persistence.ManyToOne;
import javax.persistence.Table;

@Entity
@Table(name = "sys_project_role")
public class SysProjectRole extends BaseEntity {
	private static final long serialVersionUID = -3556080851163371948L;
	
	private String name;
    private String descr;
    
    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "org_id", insertable = false, updatable = false)
    private SysOrg org;

    @Column(name = "org_id")
    private Long orgId;
    
    @ManyToMany(fetch = FetchType.LAZY, cascade = CascadeType.ALL)
	@JoinTable(name = "r_project_role_priviledge", joinColumns = { 
			@JoinColumn(name = "project_role_id", nullable = false, updatable = false) }, 
			inverseJoinColumns = { @JoinColumn(name = "project_priviledge_id", 
					nullable = false, updatable = false) })
    private Set<SysProjectPriviledge> projectPriviledgeSet = new HashSet<SysProjectPriviledge>(0);
    
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
	public Set<SysProjectPriviledge> getProjectPriviledgeSet() {
		return projectPriviledgeSet;
	}
	public void setProjectPriviledgeSet(
			Set<SysProjectPriviledge> projectPriviledgeSet) {
		this.projectPriviledgeSet = projectPriviledgeSet;
	}
	public SysOrg getOrg() {
		return org;
	}
	public void setOrg(SysOrg org) {
		this.org = org;
	}
	public Long getOrgId() {
		return orgId;
	}
	public void setOrgId(Long orgId) {
		this.orgId = orgId;
	}
}
