package com.ngtesting.platform.entity;

import javax.persistence.*;
import java.util.HashSet;
import java.util.Set;

@Entity
@Table(name = "tst_org_role")
public class TestOrgRole extends BaseEntity {
	private static final long serialVersionUID = -3556080851163371948L;

	@Enumerated(EnumType.STRING)
    private OrgRoleCode code;

	private String name;
    private String descr;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "org_id", insertable = false, updatable = false)
    private TestOrg org;

    @Column(name = "org_id")
    private Long orgId;

    @ManyToMany(fetch = FetchType.LAZY, cascade = CascadeType.ALL)
	@JoinTable(name = "tst_r_org_role_user", joinColumns = {
			@JoinColumn(name = "org_role_id", nullable = false, updatable = false) },
			inverseJoinColumns = { @JoinColumn(name = "user_id",
					nullable = false, updatable = false) })
    private Set<TestUser> userSet = new HashSet<>(0);

    @ManyToMany(fetch = FetchType.LAZY, cascade = CascadeType.ALL)
	@JoinTable(name = "tst_r_org_role_privilege", joinColumns = {
			@JoinColumn(name = "org_role_id", nullable = false, updatable = false) },
			inverseJoinColumns = { @JoinColumn(name = "org_privilege_id",
					nullable = false, updatable = false) })
    private Set<TestOrgPrivilege> orgPrivilegeSet = new HashSet<>(0);

    public static enum OrgRoleCode {
        org_admin("org_admin", "组织管理员"),
        site_admin("site_admin", "站点管理员"),
        project_admin("project_admin", "项目管理员");

        private OrgRoleCode(String code, String name) {
            this.code = code;
            this.name = name;
        }

        public String code;
        public String name;
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
	public Set<TestOrgPrivilege> getOrgPrivilegeSet() {
		return orgPrivilegeSet;
	}
	public void setOrgPrivilegeSet(
			Set<TestOrgPrivilege> orgPrivilegeSet) {
		this.orgPrivilegeSet = orgPrivilegeSet;
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
	public Set<TestUser> getUserSet() {
		return userSet;
	}
	public void setUserSet(Set<TestUser> userSet) {
		this.userSet = userSet;
	}

	public OrgRoleCode getCode() {
		return code;
	}

	public void setCode(OrgRoleCode code) {
		this.code = code;
	}
}
