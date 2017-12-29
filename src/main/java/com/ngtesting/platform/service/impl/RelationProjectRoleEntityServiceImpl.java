package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.*;
import com.ngtesting.platform.service.ProjectPrivilegeService;
import com.ngtesting.platform.service.RelationProjectRoleEntityService;
import com.ngtesting.platform.vo.RelationProjectRoleEntityVo;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;

@Service
public class RelationProjectRoleEntityServiceImpl extends BaseServiceImpl implements RelationProjectRoleEntityService {
    @Autowired
    ProjectPrivilegeService projectPrivilegeService;

    @Override
	public List<TestRelationProjectRoleEntity> listByProject(Long projectId) {
		DetachedCriteria dc = DetachedCriteria.forClass(TestRelationProjectRoleEntity.class);
		dc.add(Restrictions.eq("projectId", projectId));

        dc.addOrder(Order.asc("type"));
		dc.addOrder(Order.asc("id"));
		List<TestRelationProjectRoleEntity> ls = findAllByCriteria(dc);

		return ls;
	}

	@Override
	public TestRelationProjectRoleEntity getRelationProjectRoleEntity(Long projectRoleId) {
		DetachedCriteria dc = DetachedCriteria.forClass(TestRelationProjectRoleEntity.class);
        dc.add(Restrictions.eq("projectRoleId", projectRoleId));
        
        dc.addOrder(Order.asc("id"));
        List<TestRelationProjectRoleEntity> ls = findAllByCriteria(dc);
        
        if (ls.size() == 0) {
        	return null;
        }
		return ls.get(0);
	}

	@Override
	public List<TestRelationProjectRoleEntity> listRelationProjectRoleEntitys(Long projectRoleId) {
		DetachedCriteria dc = DetachedCriteria.forClass(TestRelationProjectRoleEntity.class);
        dc.add(Restrictions.eq("projectRoleId", projectRoleId));
        
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        
        dc.addOrder(Order.asc("id"));
        List<TestRelationProjectRoleEntity> ls = findAllByCriteria(dc);
        
		return ls;
	}

    @Override
    public TestRelationProjectRoleEntity getByProjectAndEntityId(Long projectId, Long entityId) {
        DetachedCriteria dc = DetachedCriteria.forClass(TestRelationProjectRoleEntity.class);
        dc.add(Restrictions.eq("projectId", projectId));
        dc.add(Restrictions.eq("entityId", entityId));

        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));

        dc.addOrder(Order.asc("id"));
        List<TestRelationProjectRoleEntity> ls = findAllByCriteria(dc);

        if (ls.size() > 0) {
            return ls.get(0);
        }

        return null;
    }

    @Override
    public void addUserToProjectAsLeaderPers(Long userId, Long defaultRole, Long projectId) {
        TestRelationProjectRoleEntity po = new TestRelationProjectRoleEntity(
                projectId, userId, defaultRole,
                TestRelationProjectRoleEntity.EntityType.user.toString());
        saveOrUpdate(po);
    }

    @Override
	public List<TestRelationProjectRoleEntity> batchSavePers(JSONObject json) {
        Long projectId = json.getLong("projectId");
		Long projectRoleId = json.getLong("roleId");
        List entityTypeAndIds = json.getJSONArray("entityTypeAndIds");

        List<String> relationEntityAndRoleId = new ArrayList<>();
        List<String> relationEntityId = new ArrayList<>();
		List<TestRelationProjectRoleEntity> pos = listByProject(projectId);
		for (TestRelationProjectRoleEntity po : pos) {
            relationEntityAndRoleId.add(po.getType() + "-" + po.getEntityId() + "-" + po.getProjectRoleId());
            relationEntityId.add(po.getType() + "-" + po.getEntityId());
        }

        for (Object entityTypeAndIdObj : entityTypeAndIds) {
            String[] arr = entityTypeAndIdObj.toString().split(",");
            String entityType = arr[0];
            Long entityId = Long.valueOf(arr[1]);

		    String key = entityType + "-" + entityId  + "-" + projectRoleId;
		    if (relationEntityId.contains(entityType + "-" +entityId) && !relationEntityAndRoleId.contains(key)) { // 目前为其他角色
                TestRelationProjectRoleEntity po = getByProjectAndEntityId(projectId, entityId);
                po.setProjectRoleId(projectRoleId);

                TestProjectRoleForOrg projectRole = (TestProjectRoleForOrg)get(TestProjectRoleForOrg.class, projectRoleId);
                saveOrUpdate(po);
            } else if (!relationEntityAndRoleId.contains(key)) { // 不存在
                TestProjectRoleForOrg projectRole = (TestProjectRoleForOrg)get(TestProjectRoleForOrg.class, projectRoleId);
                String name;

                if(TestRelationProjectRoleEntity.EntityType.user.toString().equals(entityType)) {
                    TestUser user = (TestUser)get(TestUser.class, entityId);
                    name = user.getName();
                } else {
                    TestOrgGroup group = (TestOrgGroup)get(TestOrgGroup.class, entityId);
                    name = group.getName();
                }

                TestRelationProjectRoleEntity po = new TestRelationProjectRoleEntity(
                        projectId, entityId, projectRoleId, entityType);
                saveOrUpdate(po);
            }
        }

		return listByProject(projectId);
	}

    @Override
    public List<TestRelationProjectRoleEntity> changeRolePers(JSONObject json) {
        Long projectId = json.getLong("projectId");
        Long projectRoleId = json.getLong("roleId");
        Long entityId = json.getLong("entityId");

        TestProjectRoleForOrg projectRole = (TestProjectRoleForOrg)get(TestProjectRoleForOrg.class, projectRoleId);

        TestRelationProjectRoleEntity po = (TestRelationProjectRoleEntity)get(TestRelationProjectRoleEntity.class, entityId);
        po.setProjectRoleId(projectRoleId);

        saveOrUpdate(po);

        return listByProject(projectId);
    }

    @Override
    public RelationProjectRoleEntityVo genVo(TestRelationProjectRoleEntity po) {
        return new RelationProjectRoleEntityVo(po.getId(), po.getProjectId(), po.getEntityId(), po.getProjectRoleId(),
                po.getProjectRole().getName(), getEntityName(po), po.getType().toString());
    }

    @Override
    public List<RelationProjectRoleEntityVo> genVos(List<TestRelationProjectRoleEntity> pos) {
        List<RelationProjectRoleEntityVo> vos = new LinkedList<>();
        for (TestRelationProjectRoleEntity po: pos) {
            RelationProjectRoleEntityVo vo = genVo(po);
            vos.add(vo);
        }
        return vos;
    }

    @Override
    public String getEntityName(TestRelationProjectRoleEntity po) {
        String name;
        if (TestRelationProjectRoleEntity.EntityType.group.equals(po.getType())) {
            TestOrgGroup group = (TestOrgGroup)get(TestOrgGroup.class, po.getEntityId());
            name = group.getName();
        } else {
            TestUser user = (TestUser)get(TestUser.class, po.getEntityId());
            name = user.getName();
        }
        return name;
    }

}
