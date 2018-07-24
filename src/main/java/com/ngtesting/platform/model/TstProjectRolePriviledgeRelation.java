package com.ngtesting.platform.model;

import java.io.Serializable;

public class TstProjectRolePriviledgeRelation implements Serializable {
    private static final long serialVersionUID = -737513767576675486L;

    private Integer orgId;
    private Integer projectPrivilegeDefineId;
    private Integer projectRoleId;

    public Integer getOrgId() {
        return orgId;
    }

    public void setOrgId(Integer orgId) {
        this.orgId = orgId;
    }

    public Integer getProjectPrivilegeDefineId() {
        return projectPrivilegeDefineId;
    }

    public void setProjectPrivilegeDefineId(Integer projectPrivilegeDefineId) {
        this.projectPrivilegeDefineId = projectPrivilegeDefineId;
    }

    public Integer getProjectRoleId() {
        return projectRoleId;
    }

    public void setProjectRoleId(Integer projectRoleId) {
        this.projectRoleId = projectRoleId;
    }
}
