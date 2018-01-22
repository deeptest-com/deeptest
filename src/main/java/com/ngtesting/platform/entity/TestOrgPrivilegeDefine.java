package com.ngtesting.platform.entity;

import javax.persistence.*;
import java.util.HashSet;
import java.util.Set;

@Entity
@Table(name = "tst_org_privilege_define")
public class TestOrgPrivilegeDefine extends BaseEntity {
	private static final long serialVersionUID = -5510206858644860272L;

	@Enumerated(EnumType.STRING)
    private OrgPrivilegeCode code;

    private String name;
    private String descr;

    @ManyToMany(cascade = {CascadeType.PERSIST, CascadeType.MERGE}, mappedBy = "orgPrivilegeSet")
    private Set<TestOrgRole> orgRoleSet = new HashSet<TestOrgRole>(0);

    public static enum OrgPrivilegeCode {
        org_admin("org_admin", "组织管理员"),
        site_admin("site_admin", "站点管理员"),
        project_admin("project_admin", "项目管理员");

        private OrgPrivilegeCode(String code, String name) {
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

	public Set<TestOrgRole> getOrgRoleSet() {
		return orgRoleSet;
	}

	public void setOrgRoleSet(Set<TestOrgRole> orgRoleSet) {
		this.orgRoleSet = orgRoleSet;
	}

	public OrgPrivilegeCode getCode() {
		return code;
	}

	public void setCode(OrgPrivilegeCode code) {
		this.code = code;
	}
}
