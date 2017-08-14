package com.ngtesting.platform.entity;

import javax.persistence.*;
import java.util.HashSet;
import java.util.Set;

@Entity
@Table(name = "tst_project_role")
public class TestProjectRole extends BaseEntity {
	private static final long serialVersionUID = -3556080851163371948L;
	
//	@Enumerated(EnumType.STRING)
    private String code;
	
	private String name;
    private String descr;
    
    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "org_id", insertable = false, updatable = false)
    private TestOrg org;

    @Column(name = "org_id")
    private Long orgId;
    
    @ManyToMany(fetch = FetchType.LAZY, cascade = CascadeType.ALL)
	@JoinTable(name = "tst_r_project_role_privilege", joinColumns = { 
			@JoinColumn(name = "project_role_id", nullable = false, updatable = false) }, 
			inverseJoinColumns = { @JoinColumn(name = "project_privilege_id", 
					nullable = false, updatable = false) })
    private Set<TestProjectPrivilege> projectPrivilegeSet = new HashSet<TestProjectPrivilege>(0);
    
//    public static enum ProjectRoleCode {
//		test_leader("test_leader"),
//        test_designer("test_designer"),
//        tester("tester"),
//		readonly("readonly");
//
//        private ProjectRoleCode(String textVal) {
//            this.textVal = textVal;
//        }
//
//        private String textVal;
//        public String toString() {
//            return textVal;
//        }
//
//    }
    
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
	public Set<TestProjectPrivilege> getProjectPrivilegeSet() {
		return projectPrivilegeSet;
	}
	public void setProjectPrivilegeSet(
			Set<TestProjectPrivilege> projectPrivilegeSet) {
		this.projectPrivilegeSet = projectPrivilegeSet;
	}
	public TestOrg getOrg() {
		return org;
	}
	public void setOrg(TestOrg org) {
		this.org = org;
	}
	public Long getOrgId() {
		return orgId;
	}
	public void setOrgId(Long orgId) {
		this.orgId = orgId;
	}

	public String getCode() {
		return code;
	}

	public void setCode(String code) {
		this.code = code;
	}
}
