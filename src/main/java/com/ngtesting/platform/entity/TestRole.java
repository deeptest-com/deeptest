package com.ngtesting.platform.entity;

import javax.persistence.*;
import java.util.HashSet;
import java.util.Set;

@Entity
@Table(name = "tst_role")
public class TestRole extends BaseEntity {
    private static final long serialVersionUID = 4490780384999462762L;

    private String name;
    private String descr;

    @ManyToMany(fetch = FetchType.LAZY, cascade = CascadeType.ALL)
	@JoinTable(name = "tst_r_role_user", joinColumns = {
			@JoinColumn(name = "role_id", nullable = false, updatable = false) },
			inverseJoinColumns = { @JoinColumn(name = "user_id",
					nullable = false, updatable = false) })
    private Set<TestUser> userSet = new HashSet<TestUser>(0);
    
    @ManyToMany(fetch = FetchType.LAZY, cascade = CascadeType.ALL)
	@JoinTable(name = "tst_r_role_privilege", joinColumns = { 
			@JoinColumn(name = "role_id", nullable = false, updatable = false) }, 
			inverseJoinColumns = { @JoinColumn(name = "privilege_id", 
					nullable = false, updatable = false) })
    private Set<TestPrivilege> privilegeSet = new HashSet<TestPrivilege>(0);

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

	public Set<TestUser> getUserSet() {
		return userSet;
	}

	public void setUserSet(Set<TestUser> userSet) {
		this.userSet = userSet;
	}

	public Set<TestPrivilege> getPrivilegeSet() {
		return privilegeSet;
	}

	public void setPrivilegeSet(Set<TestPrivilege> privilegeSet) {
		this.privilegeSet = privilegeSet;
	}
    
}
