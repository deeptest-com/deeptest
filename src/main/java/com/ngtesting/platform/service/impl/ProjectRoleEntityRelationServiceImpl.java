package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.dao.ProjectRoleEntityRelationDao;
import com.ngtesting.platform.model.TstProjectRoleEntityRelation;
import com.ngtesting.platform.service.ProjectPrivilegeService;
import com.ngtesting.platform.service.ProjectRoleEntityRelationService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.List;

@Service
public class ProjectRoleEntityRelationServiceImpl extends BaseServiceImpl implements ProjectRoleEntityRelationService {
    @Autowired
    ProjectRoleEntityRelationDao projectRoleEntityRelationDao;

    @Autowired
    ProjectPrivilegeService projectPrivilegeService;

    @Override
	public List<TstProjectRoleEntityRelation> listByProject(Integer projectId) {
        List<TstProjectRoleEntityRelation> ls = projectRoleEntityRelationDao.listByProject(projectId);

//		DetachedCriteria dc = DetachedCriteria.forClass(TstProjectRoleEntityRelation.class);
//		dc.add(Restrictions.eq("projectId", projectId));
//
//        dc.addOrder(Order.asc("type"));
//		dc.addOrder(Order.asc("id"));
//		List<TstProjectRoleEntityRelation> ls = findAllByCriteria(dc);
//
//		return ls;

        return ls;
	}

	@Override
	public TstProjectRoleEntityRelation getRelationProjectRoleEntity(Integer projectRoleId) {
//		DetachedCriteria dc = DetachedCriteria.forClass(TstProjectRoleEntityRelation.class);
//        dc.add(Restrictions.eq("projectRoleId", projectRoleId));
//
//        dc.addOrder(Order.asc("id"));
//        List<TstProjectRoleEntityRelation> ls = findAllByCriteria(dc);
//
//        if (ls.size() == 0) {
//        	return null;
//        }
//		return ls.getDetail(0);

        return null;
	}

	@Override
	public List<TstProjectRoleEntityRelation> listRelationProjectRoleEntitys(Integer projectRoleId) {
//		DetachedCriteria dc = DetachedCriteria.forClass(TstProjectRoleEntityRelation.class);
//        dc.add(Restrictions.eq("projectRoleId", projectRoleId));
//
//        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//
//        dc.addOrder(Order.asc("id"));
//        List<TstProjectRoleEntityRelation> ls = findAllByCriteria(dc);
//
//		return ls;

        return null;
	}

    @Override
    public TstProjectRoleEntityRelation getByProjectAndEntityId(Integer projectId, Integer entityId) {
//        DetachedCriteria dc = DetachedCriteria.forClass(TstProjectRoleEntityRelation.class);
//        dc.add(Restrictions.eq("projectId", projectId));
//        dc.add(Restrictions.eq("entityId", entityId));
//
//        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//
//        dc.addOrder(Order.asc("id"));
//        List<TstProjectRoleEntityRelation> ls = findAllByCriteria(dc);
//
//        if (ls.size() > 0) {
//            return ls.getDetail(0);
//        }

        return null;
    }

//    @Override
//    public void addUserToProjectAsLeaderPers(Integer userId, Integer defaultRole, Integer projectId) {
//        TstProjectRoleEntityRelation po = new TstProjectRoleEntityRelation(
//                projectId, userId, defaultRole,
//                TstProjectRoleEntityRelation.EntityType.user.toString());
//        saveOrUpdate(po);
//    }

    @Override
	public List<TstProjectRoleEntityRelation> batchSavePers(JSONObject json, Integer orgId) {
        Integer projectId = json.getInteger("projectId");
		Integer projectRoleId = json.getInteger("roleId");
        List entityTypeAndIds = json.getJSONArray("entityTypeAndIds");

        List<String> relationEntityAndRoleId = new ArrayList<>();
        List<String> relationEntityId = new ArrayList<>();
		List<TstProjectRoleEntityRelation> pos = listByProject(projectId);
		for (TstProjectRoleEntityRelation po : pos) {
            relationEntityAndRoleId.add(po.getType() + "-" + po.getEntityId() + "-" + po.getProjectRoleId());
            relationEntityId.add(po.getType() + "-" + po.getEntityId());
        }

        for (Object entityTypeAndIdObj : entityTypeAndIds) {
            String[] arr = entityTypeAndIdObj.toString().split(",");
            String entityType = arr[0];
            Integer entityId = Integer.valueOf(arr[1]);

		    String key1 = entityType + "-" + entityId  + "-" + projectRoleId;
            String key2 = entityType + "-" +entityId;
		    if (relationEntityId.contains(key2) && !relationEntityAndRoleId.contains(key1)) { // 目前为其他角色
                projectRoleEntityRelationDao.changeRole(projectId, projectRoleId, entityId);
            } else if (!relationEntityAndRoleId.contains(key1)) { // 不存在
                projectRoleEntityRelationDao.addRole(orgId, projectId, projectRoleId, entityId, entityType);
            }
        }

		return listByProject(projectId);
	}

    @Override
    public List<TstProjectRoleEntityRelation> changeRolePers(JSONObject json) {

        Integer projectId = json.getInteger("projectId");
        Integer projectRoleId = json.getInteger("roleId");
        Integer entityId = json.getInteger("entityId");

        projectRoleEntityRelationDao.changeRole(projectId, projectRoleId, entityId);

        return listByProject(projectId);
    }

    @Override
    public String getEntityName(TstProjectRoleEntityRelation po) {
        String name = null;
//        if (TstProjectRoleEntityRelation.EntityType.group.equals(po.getType())) {
//            TestOrgGroup group = (TestOrgGroup)getDetail(TestOrgGroup.class, po.getEntityId());
//            name = group.getName();
//        } else {
//            TestUser user = (TestUser)getDetail(TestUser.class, po.getEntityId());
//            name = user.getName();
//        }
        return name;
    }

}
