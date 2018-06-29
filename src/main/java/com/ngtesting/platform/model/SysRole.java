package com.ngtesting.platform.model;

import java.util.HashSet;
import java.util.Set;

public class SysRole extends BaseModel {
    private static final long serialVersionUID = 4490780384999462762L;

    private String name;
    private String descr;

    private Set<TstUser> userSet = new HashSet(0);

	private Set<SysPrivilege> sysPrivilegeSet = new HashSet(0);

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

	public Set<TstUser> getUserSet() {
		return userSet;
	}

	public void setUserSet(Set<TstUser> userSet) {
		this.userSet = userSet;
	}
}
