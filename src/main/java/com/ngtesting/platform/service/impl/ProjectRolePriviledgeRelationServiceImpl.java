package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.model.TstProjectPrivilegeDefine;
import com.ngtesting.platform.service.ProjectRolePriviledgeRelationService;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class ProjectRolePriviledgeRelationServiceImpl extends BaseServiceImpl implements ProjectRolePriviledgeRelationService {
    static String SQL = "INSERT INTO tst_project_role_priviledge_relation " +
                "( project_privilege_define_id,   project_role_id, create_time, deleted, disabled, is_build_in ) " +
     "VALUES "+ "( $project_privilege_define_id$, $project_role_id$, now(),       false,   false,    false );";

	@Override
	public String addPriviledgeForLeaderPers(List<TstProjectPrivilegeDefine> allProjectPrivileges, Integer projectRoleId) {
        String temp = "";
//	    for (TestProjectPrivilegeDefine projectPrivilege: allProjectPrivileges) {
//            temp += addPriviledgePers(projectPrivilege.getId(), projectRoleId);
//        }
        return temp;
	}
	@Override
	public String addPriviledgeForDesignerPers(List<TstProjectPrivilegeDefine> allProjectPrivileges, Integer projectRoleId) {
        String temp = "";
//        for (TestProjectPrivilegeDefine projectPrivilege: allProjectPrivileges) {
//            temp += addPriviledgePers(projectPrivilege.getId(), projectRoleId);
//        }
        return temp;
	}
	@Override
	public String addPriviledgeForTesterPers(List<TstProjectPrivilegeDefine> allProjectPrivileges, Integer projectRoleId) {
        String temp = "";
//        for (TestProjectPrivilegeDefine projectPrivilege: allProjectPrivileges) {
//            if (projectPrivilege.getCode().toString().indexOf("result") > -1) {
//                addPriviledgePers(projectPrivilege.getId(), projectRoleId);
//            }
//        }
        return temp;
	}

    @Override
    public String addPriviledgePers(Integer projectPrivilegeId, Integer projectRoleId) {
	    String temp = SQL;
        temp = temp.replace("$project_privilege_define_id$", projectPrivilegeId.toString());
        temp = temp.replace("$project_role_id$", projectRoleId.toString());

        return temp;
    }

}
