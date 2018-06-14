package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.ngtesting.platform.entity.TestProjectPrivilegeDefine;
import com.ngtesting.platform.entity.TestProjectRoleForOrg;
import com.ngtesting.platform.entity.TestProjectRolePriviledgeRelation;
import com.ngtesting.platform.entity.TestRelationProjectRoleEntity;
import com.ngtesting.platform.service.ProjectPrivilegeService;
import com.ngtesting.platform.vo.ProjectPrivilegeDefineVo;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import java.util.*;

@Service
public class ProjectPrivilegeServiceImpl extends BaseServiceImpl implements ProjectPrivilegeService {

	@Override
	public Map<String, Map<String, ProjectPrivilegeDefineVo>> listPrivilegesByOrgAndProjectRole(Long orgId, Long projectRoleId) {

        List<TestProjectPrivilegeDefine> allPrivileges = listAllProjectPrivileges();

        List<TestProjectRolePriviledgeRelation> projectRolePrivileges;
        if (projectRoleId == null) {
        	projectRolePrivileges = new LinkedList();
        } else {
        	projectRolePrivileges = listProjectRolePrivileges(projectRoleId);
        }

        Map<String, Map<String, ProjectPrivilegeDefineVo>> map = new LinkedHashMap<>();
        for (TestProjectPrivilegeDefine po1 : allPrivileges) {
        	String key = po1.getName();

			if (!map.containsKey(key)) {
				map.put(key, new HashMap<String, ProjectPrivilegeDefineVo>());
			}

        	ProjectPrivilegeDefineVo vo = genVo(orgId, po1);

        	vo.setSelected(false);
        	vo.setSelecting(false);
        	for (TestProjectRolePriviledgeRelation po2 : projectRolePrivileges) {
        		if (po1.getId().longValue() == po2.getProjectPrivilegeDefineId().longValue()) {
            		vo.setSelected(true);
            		vo.setSelecting(true);

            		Long relationId = po2.getId();
            		vo.setRelationId(relationId);
            	}
        	}
        	map.get(key).put(vo.getAction(), vo);
        }

		return map;
	}

	@Override
	public List<TestProjectRolePriviledgeRelation> listProjectRolePrivileges(Long projectRoleId) {
		DetachedCriteria dc = DetachedCriteria.forClass(TestProjectRolePriviledgeRelation.class);
        dc.add(Restrictions.eq("projectRoleId", projectRoleId));

        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));

        dc.addOrder(Order.asc("id"));
        List ls = findAllByCriteria(dc);

		return ls;
	}

	@Override
	public List<TestProjectPrivilegeDefine> listAllProjectPrivileges() {
		DetachedCriteria dc = DetachedCriteria.forClass(TestProjectPrivilegeDefine.class);

        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.addOrder(Order.asc("id"));
        List<TestProjectPrivilegeDefine> ls = findAllByCriteria(dc);

		return ls;
	}

	@Override
	public boolean addUserAsProjectTestLeaderPers(Long orgId, Long projectId, String roleCode, Long userId) {
        DetachedCriteria dc = DetachedCriteria.forClass(TestProjectRoleForOrg.class);

        dc.add(Restrictions.eq("orgId", orgId));
        dc.add(Restrictions.eq("code", "test_leader"));
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.addOrder(Order.asc("id"));
        List<TestProjectRoleForOrg> ls = findAllByCriteria(dc);
        if (ls.size() == 0) {
            return false;
        }

        TestProjectRoleForOrg role = ls.get(0);
        TestRelationProjectRoleEntity relation = new TestRelationProjectRoleEntity(
				role.getOrgId(), projectId, userId,  role.getId(), "user");
        saveOrUpdate(relation);
        return true;
	}

	@Override
	public boolean saveProjectPrivileges(Long projectRoleId, Map<String, List<ProjectPrivilegeDefineVo>> map) {
		if (map == null) {
			return false;
		}

        List<TestProjectRolePriviledgeRelation> privilegeSet = listProjectRolePrivileges(projectRoleId);
        List<Long> privilegeDefineIds = new LinkedList<>();
        for (TestProjectRolePriviledgeRelation temp: privilegeSet) {
            privilegeDefineIds.add(temp.getProjectPrivilegeDefineId());
        }

		for (String key: map.keySet()) {
            Map<String, ProjectPrivilegeDefineVo> voMap = JSON.parseObject(JSON.toJSONString(map.get(key)), Map.class);

			for (String key2: voMap.keySet()) {
				ProjectPrivilegeDefineVo vo = JSON.parseObject(JSON.toJSONString(voMap.get(key2)),
						ProjectPrivilegeDefineVo.class);

                if (vo.getSelecting() != vo.getSelected()) { // 变化了
	    			if (vo.getSelecting() && !privilegeDefineIds.contains(vo.getId())) { // 勾选
                        TestProjectRolePriviledgeRelation temp = new TestProjectRolePriviledgeRelation(vo.getId(), projectRoleId);
                        saveOrUpdate(temp);
	    			} else { // 取消
                        Long id = vo.getRelationId();
                        TestProjectRolePriviledgeRelation temp = (TestProjectRolePriviledgeRelation)get(
                                TestProjectRolePriviledgeRelation.class, id);

                        getDao().delete(temp);
	    			}

				}
			}
		}

		return true;
	}

	@Override
	public Map<String, Boolean> listByUserPers(Long userId, Long prjId, Long orgId) {
        Map<String, Boolean> map = new HashMap();
	    if (prjId == null) {
            return map;
        }
		List<Object[]> ls = getDao().getListBySQL("{call get_project_privilege_by_project_for_user(?,?,?)}", userId, prjId, orgId);
		for (Object[] arr : ls) {
			map.put(arr[0].toString() + "-" + arr[1].toString(), true);
		}
		return map;
	}

	private ProjectPrivilegeDefineVo genVo(Long orgId, TestProjectPrivilegeDefine po1) {
		ProjectPrivilegeDefineVo vo = new ProjectPrivilegeDefineVo(po1.getId(), po1.getCode().toString(), po1.getAction().toString(),
				po1.getName(), po1.getDescr(), orgId);

		return vo;
	}

}
