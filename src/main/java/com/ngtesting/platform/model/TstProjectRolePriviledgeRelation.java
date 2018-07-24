package com.ngtesting.platform.model;

import java.io.Serializable;

public class TstProjectRolePriviledgeRelation implements Serializable {
    private static final long serialVersionUID = -737513767576675486L;

//    private Boolean isBuildIn = false;

    private Long projectPrivilegeDefineId;
    private Long projectRoleId;

//    public Boolean getBuildIn() {
//        return isBuildIn;
//    }
//
//    public void setBuildIn(Boolean buildIn) {
//        isBuildIn = buildIn;
//    }

    public Long getProjectPrivilegeDefineId() {
        return projectPrivilegeDefineId;
    }

    public void setProjectPrivilegeDefineId(Long projectPrivilegeDefineId) {
        this.projectPrivilegeDefineId = projectPrivilegeDefineId;
    }

    public Long getProjectRoleId() {
        return projectRoleId;
    }

    public void setProjectRoleId(Long projectRoleId) {
        this.projectRoleId = projectRoleId;
    }
}
