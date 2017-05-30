package com.ngtesting.platform.entity;

import java.util.HashSet;
import java.util.Set;

import javax.persistence.CascadeType;
import javax.persistence.Entity;
import javax.persistence.ManyToMany;
import javax.persistence.Table;

@Entity
@Table(name = "tst_privilege")
public class TestPrivilege extends BaseEntity {
	private static final long serialVersionUID = -5510206858644860272L;

    private String name;
    private String descr;
    
    @ManyToMany(cascade = {CascadeType.PERSIST, CascadeType.MERGE}, mappedBy = "privilegeSet")
    private Set<TestRole> roleSet = new HashSet<TestRole>(0);

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

	public Set<TestRole> getRoleSet() {
		return roleSet;
	}

	public void setRoleSet(Set<TestRole> roleSet) {
		this.roleSet = roleSet;
	}
    
}
