package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.ngtesting.platform.dao.ProjectPrivilegeDao;
import com.ngtesting.platform.dao.ProjectRoleDao;
import com.ngtesting.platform.dao.ProjectRoleEntityRelationDao;
import com.ngtesting.platform.dao.ProjectRolePrivilegeRelationDao;
import com.ngtesting.platform.model.TstProjectPrivilegeDefine;
import com.ngtesting.platform.model.TstProjectRole;
import com.ngtesting.platform.model.TstProjectRolePriviledgeRelation;
import com.ngtesting.platform.service.OrgPrivilegeService;
import com.ngtesting.platform.service.ProjectPrivilegeService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.*;

@Service
public class ProjectPrivilegeServiceImpl extends BaseServiceImpl implements ProjectPrivilegeService {
	@Autowired
	OrgPrivilegeService orgPrivilegeService;

    @Autowired
    ProjectRoleDao projectRoleDao;
	@Autowired
	private ProjectPrivilegeDao projectPrivilegeDao;

    @Autowired
    ProjectRolePrivilegeRelationDao projectRolePrivilegeRelationDao;

    @Autowired
    ProjectRoleEntityRelationDao projectRoleEntityRelationDao;

	@Override
	public Map<String, Map<String, TstProjectPrivilegeDefine>> listPrivilegesByOrgAndProjectRole(
			Integer orgId, Integer projectRoleId) {

        List<TstProjectPrivilegeDefine> allPrivileges = projectPrivilegeDao.listAllProjectPrivileges();

        List<TstProjectRolePriviledgeRelation> projectRolePrivileges;
        if (projectRoleId == null) {
        	projectRolePrivileges = new LinkedList();
        } else {
        	projectRolePrivileges = projectRolePrivilegeRelationDao.listProjectRolePrivileges(projectRoleId);
        }

        Map<String, Map<String, TstProjectPrivilegeDefine>> map = new LinkedHashMap<>();
        for (TstProjectPrivilegeDefine po1 : allPrivileges) {
        	String key = po1.getName();

			if (!map.containsKey(key)) {
				map.put(key, new HashMap());
			}

            po1.setSelected(false);
            po1.setSelecting(false);
        	for (TstProjectRolePriviledgeRelation po2 : projectRolePrivileges) {
        		if (po1.getId().longValue() == po2.getProjectPrivilegeDefineId().longValue()) {
                    po1.setSelected(true);
                    po1.setSelecting(true);

//            		Integer relationId = po2.getCode();
//            		vo.setRelationId(relationId);
            	}
        	}
        	map.get(key).put(po1.getAction(), po1);
        }

		return map;
	}

	@Override
	public boolean addUserAsProjectTestLeaderPers(Integer orgId, Integer projectId, String roleCode, Integer userId) {
        TstProjectRole projectRole = projectRoleDao.getRoleByCode(orgId, roleCode);

	    projectRoleEntityRelationDao.addRole(orgId, projectId, projectRole.getId(), userId, "user");
        return true;
	}

	@Override
	public boolean saveProjectPrivileges(Integer orgId, Integer projectRoleId,
                                         Map<String, List<TstProjectPrivilegeDefine>> map) {
		if (map == null) {
			return false;
		}

        List<TstProjectPrivilegeDefine> selectedList = new LinkedList<>();
		for (String key: map.keySet()) {
            Map<String, TstProjectPrivilegeDefine> voMap = JSON.parseObject(JSON.toJSONString(map.get(key)), Map.class);

			for (String key2: voMap.keySet()) {
                TstProjectPrivilegeDefine vo = JSON.parseObject(JSON.toJSONString(voMap.get(key2)),
                        TstProjectPrivilegeDefine.class);

                if (vo.getSelecting()) {
                    selectedList.add(vo);
                }
            }
		}

        projectRolePrivilegeRelationDao.removeAllPrivilegsForRole(projectRoleId);
        if (selectedList.size() > 0) {
            projectRolePrivilegeRelationDao.saveRelations(orgId, projectRoleId, selectedList);
        }

		return true;
	}

	@Override
	public Map<String, Boolean> listByUser(Integer userId, Integer prjId, Integer orgId) {
        Map<String, Boolean> map = new HashMap();
	    if (prjId == null) {
            return map;
        }

		List<Map<String, String>> ls = projectPrivilegeDao.listByProjectForUser(userId, prjId, orgId);
		for (Map<String, String> item : ls) {
			map.put(item.get("code") + "-" + item.get("action"), true);
		}

        Map<String, Boolean> orgPrivileges = orgPrivilegeService.listByUser(userId, orgId);
        if (orgPrivileges.containsKey("project-admin") && orgPrivileges.get("project-admin")) {
            map.put("project-view", true);
            map.put("project-maintain", true);
            map.put("project-delete", true);
        }

		return map;
	}

}
