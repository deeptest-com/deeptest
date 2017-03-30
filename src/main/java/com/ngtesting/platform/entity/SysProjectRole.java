package com.ngtesting.platform.entity;

import java.util.HashSet;
import java.util.Set;

import javax.persistence.CascadeType;
import javax.persistence.Entity;
import javax.persistence.FetchType;
import javax.persistence.JoinColumn;
import javax.persistence.JoinTable;
import javax.persistence.ManyToMany;
import javax.persistence.Table;

import org.hibernate.annotations.Cache;
import org.hibernate.annotations.CacheConcurrencyStrategy;
import org.hibernate.annotations.DynamicInsert;
import org.hibernate.annotations.DynamicUpdate;

@Entity
@Table(name = "sys_project_role")
@DynamicInsert @DynamicUpdate
@Cache(usage = CacheConcurrencyStrategy.READ_WRITE)
public class SysProjectRole extends BaseEntity {
	private static final long serialVersionUID = -3556080851163371948L;
	
	private String name;
    private String descr;
    
    @ManyToMany(fetch = FetchType.LAZY, cascade = CascadeType.ALL)
	@JoinTable(name = "r_project_role_priviledge", joinColumns = { 
			@JoinColumn(name = "role_id", nullable = false, updatable = false) }, 
			inverseJoinColumns = { @JoinColumn(name = "priviledge_id", 
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
}
