package com.ngtesting.platform.entity;

import java.util.HashSet;
import java.util.Set;

import javax.persistence.CascadeType;
import javax.persistence.Entity;
import javax.persistence.EnumType;
import javax.persistence.Enumerated;
import javax.persistence.ManyToMany;
import javax.persistence.Table;

@Entity
@Table(name = "tst_org_privilege")
public class TestOrgPrivilege extends BaseEntity {
	private static final long serialVersionUID = -5510206858644860272L;
	
	@Enumerated(EnumType.STRING)
    private OrgPrivilegeCode code;

    private String name;
    private String descr;
    
    @ManyToMany(cascade = {CascadeType.PERSIST, CascadeType.MERGE}, mappedBy = "orgPrivilegeSet")
    private Set<TestOrgRole> orgRoleSet = new HashSet<TestOrgRole>(0);
    
    public static enum OrgPrivilegeCode {
    	org_admin("org_admin"),
    	site_admin("site_admin"),
    	project_admin("project_admin");

        private OrgPrivilegeCode(String textVal) {
            this.textVal = textVal;
        }

        private String textVal;
        public String toString() {
            return textVal;
        }
        
        public static OrgPrivilegeCode getValue(String str) {
        	OrgPrivilegeCode status = null;
        	switch(str) { 
            	case "org_admin": status = OrgPrivilegeCode.org_admin; break;
            	case "site_admin": status = OrgPrivilegeCode.site_admin; break;
            	case "project_admin": status = OrgPrivilegeCode.project_admin; break;
            }
        	
        	return status;
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

	public Set<TestOrgRole> getOrgRoleSet() {
		return orgRoleSet;
	}

	public void setOrgRoleSet(Set<TestOrgRole> orgRoleSet) {
		this.orgRoleSet = orgRoleSet;
	}
    
}
