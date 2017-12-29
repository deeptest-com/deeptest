package com.ngtesting.platform.entity;

import javax.persistence.*;

@Entity
@Table(name = "tst_project_role_for_org")
public class TestProjectRoleForOrg extends BaseEntity {
	private static final long serialVersionUID = -3556080851163371948L;

    private String code;

	private String name;
    private String descr;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "org_id", insertable = false, updatable = false)
    private TestOrg org;

    @Column(name = "org_id")
    private Long orgId;

	private Boolean isBuildIn = false;

//    @ManyToMany(fetch = FetchType.LAZY, cascade = CascadeType.ALL)
//	@JoinTable(name = "tst_r_project_role_privilege", joinColumns = {
//			@JoinColumn(name = "project_role_id", nullable = false, updatable = false) },
//			inverseJoinColumns = { @JoinColumn(name = "project_privilege_id",
//					nullable = false, updatable = false) })
//    private Set<TestProjectPrivilegeDefine> projectPrivilegeSet = new HashSet<TestProjectPrivilegeDefine>(0);

	public Boolean getBuildIn() {
		return isBuildIn;
	}
	public void setBuildIn(Boolean buildIn) {
		isBuildIn = buildIn;
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
//	public Set<TestProjectPrivilegeDefine> getProjectPrivilegeSet() {
//		return projectPrivilegeSet;
//	}
//	public void setProjectPrivilegeSet(
//			Set<TestProjectPrivilegeDefine> projectPrivilegeSet) {
//		this.projectPrivilegeSet = projectPrivilegeSet;
//	}
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
