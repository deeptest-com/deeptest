package com.ngtesting.platform.entity;

import javax.persistence.*;
import java.util.HashSet;
import java.util.Set;

@Entity
@Table(name = "sys_privilege")
public class SysPrivilege extends BaseEntity {
	private static final long serialVersionUID = -3213715765380446311L;

	@Enumerated(EnumType.STRING)
    private SysPrivilegeCode code;

    private String name;
    private String descr;

    @ManyToMany(cascade = {CascadeType.PERSIST, CascadeType.MERGE}, mappedBy = "sysPrivilegeSet")
    private Set<SysRole> sysRoleSet = new HashSet<>(0);

    public static enum SysPrivilegeCode {
    	sys_admin("sys_admin"),
    	data_backup("data_backup"),
    	account_admin("account_admin");

        private SysPrivilegeCode(String textVal) {
            this.textVal = textVal;
        }

        private String textVal;
        public String toString() {
            return textVal;
        }
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

	public SysPrivilegeCode getCode() {
		return code;
	}

	public void setCode(SysPrivilegeCode code) {
		this.code = code;
	}

	public Set<SysRole> getSysRoleSet() {
		return sysRoleSet;
	}

	public void setSysRoleSet(Set<SysRole> sysRoleSet) {
		this.sysRoleSet = sysRoleSet;
	}
}
