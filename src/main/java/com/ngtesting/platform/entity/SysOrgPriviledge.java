package com.ngtesting.platform.entity;

import java.util.HashSet;
import java.util.Set;

import javax.persistence.CascadeType;
import javax.persistence.Entity;
import javax.persistence.ManyToMany;
import javax.persistence.Table;

@Entity
@Table(name = "sys_org_priviledge")
public class SysOrgPriviledge extends BaseEntity {
	private static final long serialVersionUID = -5510206858644860272L;

    private String name;
    private String descr;
    
    @ManyToMany(cascade = {CascadeType.PERSIST, CascadeType.MERGE}, mappedBy = "orgPriviledgeSet")
    private Set<SysOrgRole> orgRoleSet = new HashSet<SysOrgRole>(0);

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

	public Set<SysOrgRole> getOrgRoleSet() {
		return orgRoleSet;
	}

	public void setOrgRoleSet(Set<SysOrgRole> orgRoleSet) {
		this.orgRoleSet = orgRoleSet;
	}
    
}
