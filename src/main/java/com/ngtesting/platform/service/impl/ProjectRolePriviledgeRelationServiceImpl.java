package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.entity.TestProjectPrivilegeDefine;
import com.ngtesting.platform.entity.TestProjectRolePriviledgeRelation;
import com.ngtesting.platform.service.ProjectPrivilegeService;
import com.ngtesting.platform.service.ProjectRolePriviledgeRelationService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class ProjectRolePriviledgeRelationServiceImpl extends BaseServiceImpl implements ProjectRolePriviledgeRelationService {
    @Autowired
    private ProjectPrivilegeService projectPrivilegeService;

	@Override
	public void addPriviledgeForLeaderPers(Long projectRoleId) {
        List<TestProjectPrivilegeDefine> allProjectPrivileges =
                projectPrivilegeService.listAllProjectPrivileges();
        for (TestProjectPrivilegeDefine projectPrivilege: allProjectPrivileges) {
            addPriviledgePers(projectPrivilege.getId(), projectRoleId);
        }
	}
	@Override
	public void addPriviledgeForDesignerPers(Long projectRoleId) {
        List<TestProjectPrivilegeDefine> allProjectPrivileges =
                projectPrivilegeService.listAllProjectPrivileges();
        for (TestProjectPrivilegeDefine projectPrivilege: allProjectPrivileges) {
            addPriviledgePers(projectPrivilege.getId(), projectRoleId);
        }
	}
	@Override
	public void addPriviledgeForTesterPers(Long projectRoleId) {
        List<TestProjectPrivilegeDefine> allProjectPrivileges =
                projectPrivilegeService.listAllProjectPrivileges();
        for (TestProjectPrivilegeDefine projectPrivilege: allProjectPrivileges) {
            if (projectPrivilege.getCode().toString().indexOf("result") > -1) {
                addPriviledgePers(projectPrivilege.getId(), projectRoleId);
            }
        }
	}

    @Override
    public void addPriviledgePers(Long projectPrivilegeId, Long projectRoleId) {
        TestProjectRolePriviledgeRelation po =
                new TestProjectRolePriviledgeRelation(projectPrivilegeId, projectRoleId);
        saveOrUpdate(po);
    }

}
