package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.TestProjectRole;
import com.ngtesting.platform.entity.TestRelationProjectRoleUser;
import com.ngtesting.platform.entity.TestUser;
import com.ngtesting.platform.service.RelationProjectRoleUserService;
import com.ngtesting.platform.vo.RelationProjectRoleUserVo;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;

@Service
public class RelationProjectRoleUserServiceImpl extends BaseServiceImpl implements RelationProjectRoleUserService {
	@Override
	public List<TestRelationProjectRoleUser> listByProject(Long projectId) {
		DetachedCriteria dc = DetachedCriteria.forClass(TestRelationProjectRoleUser.class);
		dc.add(Restrictions.eq("projectId", projectId));

		dc.addOrder(Order.asc("id"));
		List<TestRelationProjectRoleUser> ls = findAllByCriteria(dc);

		return ls;
	}

	@Override
	public TestRelationProjectRoleUser getRelationProjectRoleUser(Long projectRoleId) {
		DetachedCriteria dc = DetachedCriteria.forClass(TestRelationProjectRoleUser.class);
        dc.add(Restrictions.eq("projectRoleId", projectRoleId));
        
        dc.addOrder(Order.asc("id"));
        List<TestRelationProjectRoleUser> ls = findAllByCriteria(dc);
        
        if (ls.size() == 0) {
        	return null;
        }
		return ls.get(0);
	}

	@Override
	public List<TestRelationProjectRoleUser> listRelationProjectRoleUsers(Long projectRoleId) {
		DetachedCriteria dc = DetachedCriteria.forClass(TestRelationProjectRoleUser.class);
        dc.add(Restrictions.eq("projectRoleId", projectRoleId));
        
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        
        dc.addOrder(Order.asc("id"));
        List<TestRelationProjectRoleUser> ls = findAllByCriteria(dc);
        
		return ls;
	}

    @Override
    public TestRelationProjectRoleUser getByUserId(Long userId) {
        DetachedCriteria dc = DetachedCriteria.forClass(TestRelationProjectRoleUser.class);
        dc.add(Restrictions.eq("userId", userId));

        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));

        dc.addOrder(Order.asc("id"));
        List<TestRelationProjectRoleUser> ls = findAllByCriteria(dc);

        if (ls.size() > 0) {
            return ls.get(0);
        }

        return null;
    }

	@Override
	public List<TestRelationProjectRoleUser> batchSavePers(JSONObject json) {
        Long projectId = json.getLong("projectId");
		Long projectRoleId = json.getLong("roleId");
        List userIds = json.getJSONArray("userIds");

        List<String> relationUserAndRoleId = new ArrayList<>();
        List<Long> relationUserId = new ArrayList<>();
		List<TestRelationProjectRoleUser> pos = listByProject(projectId);
		for (TestRelationProjectRoleUser po : pos) {
            relationUserAndRoleId.add(po.getUserId() + "-" + po.getProjectRoleId());
            relationUserId.add(po.getUserId());
        }

        for (Object userIdObj : userIds) {
            Long userId = Long.valueOf(userIdObj.toString());

		    String key = userId  + "-" + projectRoleId;
		    if (relationUserId.contains(userId) && !relationUserAndRoleId.contains(key)) { // 目前为其他角色
                TestRelationProjectRoleUser po = getByUserId(userId);
                po.setProjectRoleId(projectRoleId);

                TestProjectRole projectRole = (TestProjectRole)get(TestProjectRole.class, projectRoleId);
                po.setProjectRoleName(projectRole.getName());
                saveOrUpdate(po);
            } else if (!relationUserAndRoleId.contains(key)) { // 用户不存在
                TestProjectRole projectRole = (TestProjectRole)get(TestProjectRole.class, projectRoleId);
                TestUser user = (TestUser)get(TestUser.class, userId);

                TestRelationProjectRoleUser po = new TestRelationProjectRoleUser(
                        projectId, userId, projectRoleId, projectRole.getName(), user.getName());
                saveOrUpdate(po);
            }
        }

		return listByProject(projectId);
	}

    @Override
    public RelationProjectRoleUserVo genVo(TestRelationProjectRoleUser po) {
        return new RelationProjectRoleUserVo(po.getProjectId(), po.getUserId(), po.getProjectRoleId(), po.getProjectRoleName(), po.getUserName());
    }

    @Override
    public List<RelationProjectRoleUserVo> genVos(List<TestRelationProjectRoleUser> pos) {
        List<RelationProjectRoleUserVo> vos = new LinkedList<>();
        for (TestRelationProjectRoleUser po: pos) {
            RelationProjectRoleUserVo vo = genVo(po);
            vos.add(vo);
        }
        return vos;
    }

}
