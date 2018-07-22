package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.ProjectPrivilegeDao;
import com.ngtesting.platform.dao.ProjectRoleDao;
import com.ngtesting.platform.dao.ProjectRoleEntityRelationDao;
import com.ngtesting.platform.model.TstProjectPrivilegeDefine;
import com.ngtesting.platform.model.TstProjectRole;
import com.ngtesting.platform.model.TstProjectRolePriviledgeRelation;
import com.ngtesting.platform.service.ProjectPrivilegeService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

@Service
public class ProjectPrivilegeServiceImpl extends BaseServiceImpl implements ProjectPrivilegeService {
    @Autowired
    ProjectRoleDao projectRoleDao;
	@Autowired
	private ProjectPrivilegeDao projectPrivilegeDao;
    @Autowired
    ProjectRoleEntityRelationDao projectRoleEntityRelationDao;

	@Override
	public Map<String, Map<String, TstProjectPrivilegeDefine>> listPrivilegesByOrgAndProjectRole(Integer orgId, Integer projectRoleId) {

//        List<TstProjectPrivilegeDefine> allPrivileges = listAllProjectPrivileges();
//
//        List<TstProjectRolePriviledgeRelation> projectRolePrivileges;
//        if (projectRoleId == null) {
//        	projectRolePrivileges = new LinkedList();
//        } else {
//        	projectRolePrivileges = listProjectRolePrivileges(projectRoleId);
//        }
//
//        Map<String, Map<String, TstProjectPrivilegeDefine>> map = new LinkedHashMap<>();
//        for (TstProjectPrivilegeDefine po1 : allPrivileges) {
//        	String key = po1.getName();
//
//			if (!map.containsKey(key)) {
//				map.put(key, new HashMap<String, TstProjectPrivilegeDefine>());
//			}
//
//        	TstProjectPrivilegeDefine vo = genVo(orgId, po1);
//
//        	vo.setSelected(false);
//        	vo.setSelecting(false);
//        	for (TstProjectRolePriviledgeRelation po2 : projectRolePrivileges) {
//        		if (po1.getId().longValue() == po2.getProjectPrivilegeDefineId().longValue()) {
//            		vo.setSelected(true);
//            		vo.setSelecting(true);
//
//            		Integer relationId = po2.getId();
//            		vo.setRelationId(relationId);
//            	}
//        	}
//        	map.get(key).put(vo.getAction(), vo);
//        }
//
//		return map;

		return null;
	}

	@Override
	public List<TstProjectRolePriviledgeRelation> listProjectRolePrivileges(Integer projectRoleId) {
//		DetachedCriteria dc = DetachedCriteria.forClass(TstProjectRolePriviledgeRelation.class);
//        dc.add(Restrictions.eq("projectRoleId", projectRoleId));
//
//        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//
//        dc.addOrder(Order.asc("id"));
//        List ls = findAllByCriteria(dc);
//
//		return ls;

		return null;
	}

	@Override
	public List<TstProjectPrivilegeDefine> listAllProjectPrivileges() {
//		DetachedCriteria dc = DetachedCriteria.forClass(TstProjectPrivilegeDefine.class);
//
//        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//        dc.addOrder(Order.asc("id"));
//        List<TstProjectPrivilegeDefine> ls = findAllByCriteria(dc);
//
//		return ls;

		return null;
	}

	@Override
	public boolean addUserAsProjectTestLeaderPers(Integer orgId, Integer projectId, String roleCode, Integer userId) {
        TstProjectRole projectRole = projectRoleDao.getRoleByCode(orgId, roleCode);

	    projectRoleEntityRelationDao.addRole(orgId, projectId, projectRole.getId(), userId, "user");
        return true;
	}

	@Override
	public boolean saveProjectPrivileges(Integer projectRoleId, Map<String, List<TstProjectPrivilegeDefine>> map) {
//		if (map == null) {
//			return false;
//		}
//
//        List<TstProjectRolePriviledgeRelation> privilegeSet = listProjectRolePrivileges(projectRoleId);
//        List<Integer> privilegeDefineIds = new LinkedList<>();
//        for (TstProjectRolePriviledgeRelation temp: privilegeSet) {
//            privilegeDefineIds.add(temp.getProjectPrivilegeDefineId());
//        }
//
//		for (String key: map.keySet()) {
//            Map<String, TstProjectPrivilegeDefine> voMap = JSON.parseObject(JSON.toJSONString(map.get(key)), Map.class);
//
//			for (String key2: voMap.keySet()) {
//				TstProjectPrivilegeDefine vo = JSON.parseObject(JSON.toJSONString(voMap.get(key2)),
//						TstProjectPrivilegeDefine.class);
//
//                if (vo.getSelecting() != vo.getSelected()) { // 变化了
//	    			if (vo.getSelecting() && !privilegeDefineIds.contains(vo.getId())) { // 勾选
//                        TstProjectRolePriviledgeRelation temp = new TstProjectRolePriviledgeRelation(vo.getId(), projectRoleId);
//                        saveOrUpdate(temp);
//	    			} else { // 取消
//                        Integer id = vo.getRelationId();
//                        TstProjectRolePriviledgeRelation temp = (TstProjectRolePriviledgeRelation)get(
//                                TstProjectRolePriviledgeRelation.class, id);
//
//                        getDao().delete(temp);
//	    			}
//
//				}
//			}
//		}

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
		return map;
	}

	private TstProjectPrivilegeDefine genVo(Integer orgId, TstProjectPrivilegeDefine po1) {
//		TstProjectPrivilegeDefine vo = new TstProjectPrivilegeDefine(po1.getId(), po1.getCode().toString(), po1.getAction().toString(),
//				po1.getName(), po1.getDescr(), orgId);
//
//		return vo;

		return null;
	}

}
