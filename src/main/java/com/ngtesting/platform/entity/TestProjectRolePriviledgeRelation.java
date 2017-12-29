package com.ngtesting.platform.entity;

import javax.persistence.*;

@Entity
@Table(name = "tst_project_role_priviledge_relation")
public class TestProjectRolePriviledgeRelation extends BaseEntity {
    private static final long serialVersionUID = -737513767576675486L;

    private Boolean isBuildIn = false;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "project_privilege_define_id", insertable = false, updatable = false)
    private TestProjectPrivilegeDefine projectPrivilegeDefine;
    @Column(name = "project_privilege_define_id")
    private Long projectPrivilegeDefineId;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "project_role_id", insertable = false, updatable = false)
    private TestProjectRoleForOrg projectRole;
    @Column(name = "project_role_id")
    private Long projectRoleId;

    public TestProjectRolePriviledgeRelation(){};
    public TestProjectRolePriviledgeRelation(Long projectPrivilegeDefineId, Long projectRoleId) {
        super();
        this.projectPrivilegeDefineId = projectPrivilegeDefineId;
        this.projectRoleId = projectRoleId;
    }

    public TestProjectPrivilegeDefine getProjectPrivilegeDefine() {
        return projectPrivilegeDefine;
    }

    public void setProjectPrivilegeDefine(TestProjectPrivilegeDefine projectPrivilegeDefine) {
        this.projectPrivilegeDefine = projectPrivilegeDefine;
    }

    public Long getProjectPrivilegeDefineId() {
        return projectPrivilegeDefineId;
    }

    public void setProjectPrivilegeDefineId(Long projectPrivilegeDefineId) {
        this.projectPrivilegeDefineId = projectPrivilegeDefineId;
    }

    public TestProjectRoleForOrg getProjectRole() {
        return projectRole;
    }

    public void setProjectRole(TestProjectRoleForOrg projectRole) {
        this.projectRole = projectRole;
    }

    public Long getProjectRoleId() {
        return projectRoleId;
    }

    public void setProjectRoleId(Long projectRoleId) {
        this.projectRoleId = projectRoleId;
    }

    public Boolean getBuildIn() {
        return isBuildIn;
    }

    public void setBuildIn(Boolean buildIn) {
        isBuildIn = buildIn;
    }
}
