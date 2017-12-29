package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.ngtesting.platform.entity.TestProjectPrivilegeDefine;
import com.ngtesting.platform.entity.TestProjectRolePriviledgeRelation;
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
	public Map<String, List<ProjectPrivilegeDefineVo>> listPrivilegesByOrgAndProjectRole(Long orgId, Long projectRoleId) {

        List<TestProjectPrivilegeDefine> allPrivileges = listAllProjectPrivileges();

        List<TestProjectRolePriviledgeRelation> projectRolePrivileges;
        if (projectRoleId == null) {
        	projectRolePrivileges = new LinkedList();
        } else {
        	projectRolePrivileges = listProjectRolePrivileges(projectRoleId);
        }

        Map<String, List<ProjectPrivilegeDefineVo>> map = new LinkedHashMap<String, List<ProjectPrivilegeDefineVo>>();
        for (TestProjectPrivilegeDefine po1 : allPrivileges) {
        	String key = po1.getName();
        	if (!map.containsKey(key)) {
        		List<ProjectPrivilegeDefineVo> vos = new LinkedList();
        		map.put(key, vos);
        	}

        	ProjectPrivilegeDefineVo vo = genVo(orgId, po1);

        	vo.setSelected(false);
        	vo.setSelecting(false);
        	for (TestProjectRolePriviledgeRelation po2 : projectRolePrivileges) {
        		if (po1.getId().longValue() == po2.getProjectPrivilegeDefineId().longValue()) {
            		vo.setSelected(true);
            		vo.setSelecting(true);
            		vo.setRelationId(po2.getId());
            	}
        	}
        	map.get(key).add(vo);
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
			List<ProjectPrivilegeDefineVo> ls = JSON.parseObject(JSON.toJSONString(map.get(key)), List.class);

			for (ProjectPrivilegeDefineVo vo: ls) {
                if (vo.getSelecting() != vo.getSelected()) { // 变化了

	    			if (vo.getSelecting() && !privilegeDefineIds.contains(vo.getId())) { // 勾选
                        TestProjectRolePriviledgeRelation temp = new TestProjectRolePriviledgeRelation(vo.getId(), projectRoleId);
                        saveOrUpdate(temp);
	    			} else { // 取消
                        TestProjectRolePriviledgeRelation temp = (TestProjectRolePriviledgeRelation)get(TestProjectRolePriviledgeRelation.class,
                                vo.getRelationId());
                        temp.setDeleted(true);
                        saveOrUpdate(temp);
	    			}

				}
			}
		}

		return true;
	}

	@Override
	public Map<String, Boolean> listByUser(Long userId, Long prjId) {
        Map<String, Boolean> map = new HashMap();
	    if (prjId == null) {
            return map;
        }
//		String hql = "select entiy.projectId, priv.code, priv.action from TestProjectPrivilegeDefine priv" +
//				" join priv.projectRoleSet roles, " +
//				"  TestRelationProjectRoleEntity entiy " +
//
//				" where entiy.entityId = ?" +
//				" and entiy.projectRoleId = roles.id" +
//				" and entiy.type = 'user'" +
//
//				" and priv.deleted != true and priv.disabled!= true " +
//				" order by priv.id asc";
//
//		List<Object[]> ls = getDao().getListByHQL(hql, userId);
//
//		for (Object[] raw: ls) {
//		    System.out.println(raw.getClass());
//		    Long projectId = Long.valueOf(raw[0].toString());
//			if (!map.containsKey(projectId)) {
//                map.put("prj" + projectId, new HashMap());
//			}
//
//			map.get("prj" + projectId).put(raw[1].toString() + "-" + raw[2].toString(), true);
//		}

		return null;
	}

	private ProjectPrivilegeDefineVo genVo(Long orgId, TestProjectPrivilegeDefine po1) {
		ProjectPrivilegeDefineVo vo = new ProjectPrivilegeDefineVo(po1.getId(), po1.getCode().toString(), po1.getAction().toString(),
				po1.getName(), po1.getDescr(), orgId);

		return vo;
	}

}
