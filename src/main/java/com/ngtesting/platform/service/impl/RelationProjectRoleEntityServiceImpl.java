package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.model.TstRelationProjectRoleEntity;
import com.ngtesting.platform.service.ProjectPrivilegeService;
import com.ngtesting.platform.service.RelationProjectRoleEntityService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class RelationProjectRoleEntityServiceImpl extends BaseServiceImpl implements RelationProjectRoleEntityService {
    @Autowired
    ProjectPrivilegeService projectPrivilegeService;

    @Override
	public List<TstRelationProjectRoleEntity> listByProject(Integer projectId) {
//		DetachedCriteria dc = DetachedCriteria.forClass(TstRelationProjectRoleEntity.class);
//		dc.add(Restrictions.eq("projectId", projectId));
//
//        dc.addOrder(Order.asc("type"));
//		dc.addOrder(Order.asc("id"));
//		List<TstRelationProjectRoleEntity> ls = findAllByCriteria(dc);
//
//		return ls;

        return null;
	}

	@Override
	public TstRelationProjectRoleEntity getRelationProjectRoleEntity(Integer projectRoleId) {
//		DetachedCriteria dc = DetachedCriteria.forClass(TstRelationProjectRoleEntity.class);
//        dc.add(Restrictions.eq("projectRoleId", projectRoleId));
//
//        dc.addOrder(Order.asc("id"));
//        List<TstRelationProjectRoleEntity> ls = findAllByCriteria(dc);
//
//        if (ls.size() == 0) {
//        	return null;
//        }
//		return ls.get(0);

        return null;
	}

	@Override
	public List<TstRelationProjectRoleEntity> listRelationProjectRoleEntitys(Integer projectRoleId) {
//		DetachedCriteria dc = DetachedCriteria.forClass(TstRelationProjectRoleEntity.class);
//        dc.add(Restrictions.eq("projectRoleId", projectRoleId));
//
//        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//
//        dc.addOrder(Order.asc("id"));
//        List<TstRelationProjectRoleEntity> ls = findAllByCriteria(dc);
//
//		return ls;

        return null;
	}

    @Override
    public TstRelationProjectRoleEntity getByProjectAndEntityId(Integer projectId, Integer entityId) {
//        DetachedCriteria dc = DetachedCriteria.forClass(TstRelationProjectRoleEntity.class);
//        dc.add(Restrictions.eq("projectId", projectId));
//        dc.add(Restrictions.eq("entityId", entityId));
//
//        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//
//        dc.addOrder(Order.asc("id"));
//        List<TstRelationProjectRoleEntity> ls = findAllByCriteria(dc);
//
//        if (ls.size() > 0) {
//            return ls.get(0);
//        }

        return null;
    }

//    @Override
//    public void addUserToProjectAsLeaderPers(Integer userId, Integer defaultRole, Integer projectId) {
//        TstRelationProjectRoleEntity po = new TstRelationProjectRoleEntity(
//                projectId, userId, defaultRole,
//                TstRelationProjectRoleEntity.EntityType.user.toString());
//        saveOrUpdate(po);
//    }

    @Override
	public List<TstRelationProjectRoleEntity> batchSavePers(JSONObject json) {
//        Integer projectId = json.getInteger("projectId");
//		Integer projectRoleId = json.getInteger("roleId");
//        List entityTypeAndIds = json.getJSONArray("entityTypeAndIds");
//
//        List<String> relationEntityAndRoleId = new ArrayList<>();
//        List<String> relationEntityId = new ArrayList<>();
//		List<TstRelationProjectRoleEntity> pos = listByProject(projectId);
//		for (TstRelationProjectRoleEntity po : pos) {
//            relationEntityAndRoleId.add(po.getType() + "-" + po.getEntityId() + "-" + po.getProjectRoleId());
//            relationEntityId.add(po.getType() + "-" + po.getEntityId());
//        }
//
//        for (Object entityTypeAndIdObj : entityTypeAndIds) {
//            String[] arr = entityTypeAndIdObj.toString().split(",");
//            String entityType = arr[0];
//            Integer entityId = Integer.valueOf(arr[1]);
//
//		    String key = entityType + "-" + entityId  + "-" + projectRoleId;
//		    if (relationEntityId.contains(entityType + "-" +entityId) && !relationEntityAndRoleId.contains(key)) { // 目前为其他角色
//                TstRelationProjectRoleEntity po = getByProjectAndEntityId(projectId, entityId);
//                po.setProjectRoleId(projectRoleId);
//
//                TestProjectRoleForOrg projectRole = (TestProjectRoleForOrg)get(TestProjectRoleForOrg.class, projectRoleId);
//                saveOrUpdate(po);
//            } else if (!relationEntityAndRoleId.contains(key)) { // 不存在
//                TestProjectRoleForOrg projectRole = (TestProjectRoleForOrg)get(TestProjectRoleForOrg.class, projectRoleId);
//                String name;
//
//                if(TstRelationProjectRoleEntity.EntityType.user.toString().equals(entityType)) {
//                    TestUser user = (TestUser)get(TestUser.class, entityId);
//                    name = user.getName();
//                } else {
//                    TestOrgGroup group = (TestOrgGroup)get(TestOrgGroup.class, entityId);
//                    name = group.getName();
//                }
//
//                TstRelationProjectRoleEntity po = new TstRelationProjectRoleEntity(
//                        projectRole.getOrgId(), projectId, entityId, projectRoleId, entityType);
//                saveOrUpdate(po);
//            }
//        }
//
//		return listByProject(projectId);

        return null;
	}

    @Override
    public List<TstRelationProjectRoleEntity> changeRolePers(JSONObject json) {
//        Integer projectId = json.getInteger("projectId");
//        Integer projectRoleId = json.getInteger("roleId");
//        Integer entityId = json.getInteger("entityId");
//
//        TestProjectRoleForOrg projectRole = (TestProjectRoleForOrg)get(TestProjectRoleForOrg.class, projectRoleId);
//
//        TstRelationProjectRoleEntity po = (TstRelationProjectRoleEntity)get(TstRelationProjectRoleEntity.class, entityId);
//        po.setProjectRoleId(projectRoleId);
//
//        saveOrUpdate(po);
//
//        return listByProject(projectId);

        return null;
    }

    @Override
    public String getEntityName(TstRelationProjectRoleEntity po) {
        String name = null;
//        if (TstRelationProjectRoleEntity.EntityType.group.equals(po.getType())) {
//            TestOrgGroup group = (TestOrgGroup)get(TestOrgGroup.class, po.getEntityId());
//            name = group.getName();
//        } else {
//            TestUser user = (TestUser)get(TestUser.class, po.getEntityId());
//            name = user.getName();
//        }
        return name;
    }

}
