package com.ngtesting.platform.entity;

import java.util.HashSet;
import java.util.Set;

import javax.persistence.CascadeType;
import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.EnumType;
import javax.persistence.Enumerated;
import javax.persistence.FetchType;
import javax.persistence.JoinColumn;
import javax.persistence.JoinTable;
import javax.persistence.ManyToMany;
import javax.persistence.ManyToOne;
import javax.persistence.Table;

import com.ngtesting.platform.entity.SysOrgRole.OrgRoleCode;

@Entity
@Table(name = "sys_project_role")
public class SysProjectRole extends BaseEntity {
	private static final long serialVersionUID = -3556080851163371948L;
	
	@Enumerated(EnumType.STRING)
    private ProjectRoleCode code;
	
	private String name;
    private String descr;
    
    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "org_id", insertable = false, updatable = false)
    private SysOrg org;

    @Column(name = "org_id")
    private Long orgId;
    
    @ManyToMany(fetch = FetchType.LAZY, cascade = CascadeType.ALL)
	@JoinTable(name = "r_project_role_privilege", joinColumns = { 
			@JoinColumn(name = "project_role_id", nullable = false, updatable = false) }, 
			inverseJoinColumns = { @JoinColumn(name = "project_privilege_id", 
					nullable = false, updatable = false) })
    private Set<SysProjectPrivilege> projectPrivilegeSet = new HashSet<SysProjectPrivilege>(0);
    
    public static enum ProjectRoleCode {
    	project_manager("project_manager"),
        test_designer("test_designer"),
        tester("tester"),
        developer("developer");

        private ProjectRoleCode(String textVal) {
            this.textVal = textVal;
        }

        private String textVal;
        public String toString() {
            return textVal;
        }
        
        public static ProjectRoleCode getValue(String str) {
        	ProjectRoleCode status = null;
        	switch(str) { 
            	case "project_manager": status = ProjectRoleCode.project_manager; break;
            	case "test_designer": status = ProjectRoleCode.test_designer; break;
            	case "tester": status = ProjectRoleCode.tester; break;
            	case "developer": status = ProjectRoleCode.developer; break;
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
	public Set<SysProjectPrivilege> getProjectPrivilegeSet() {
		return projectPrivilegeSet;
	}
	public void setProjectPrivilegeSet(
			Set<SysProjectPrivilege> projectPrivilegeSet) {
		this.projectPrivilegeSet = projectPrivilegeSet;
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
