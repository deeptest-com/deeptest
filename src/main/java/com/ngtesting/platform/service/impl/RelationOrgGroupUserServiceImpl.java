package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.ngtesting.platform.entity.TestOrgGroup;
import com.ngtesting.platform.entity.TestRelationOrgGroupUser;
import com.ngtesting.platform.entity.TestUser;
import com.ngtesting.platform.service.RelationOrgGroupUserService;
import com.ngtesting.platform.vo.RelationOrgGroupUserVo;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;

@Service
public class RelationOrgGroupUserServiceImpl extends BaseServiceImpl implements RelationOrgGroupUserService {

	@Override
	public List<RelationOrgGroupUserVo> listRelationsByUser(Long orgId, Long userId) {

        List<TestOrgGroup> allOrgGroups = listAllOrgGroups(orgId);
        
        List<TestRelationOrgGroupUser> relations;
        if (userId == null) {
        	relations = new LinkedList<>();
        } else {
        	relations = listRelations(orgId, null, userId);
        }
        
        List<RelationOrgGroupUserVo> vos = new LinkedList<>();
        for (TestOrgGroup orgGroup : allOrgGroups) {
        	RelationOrgGroupUserVo vo = genVo(orgId, orgGroup.getId(), userId);
        	
        	vo.setSelected(false);
        	vo.setSelecting(false);
        	for (TestRelationOrgGroupUser po : relations) {
        		if (po.getOrgGroupId() == orgGroup.getId() && po.getUserId() == userId) {
            		vo.setSelected(true);
            		vo.setSelecting(true);
            	}
        	}
        	vos.add(vo);
        }
        
		return vos;
	}
	
	@Override
	public List<RelationOrgGroupUserVo> listRelationsByGroup(Long orgId, Long orgGroupId) {

        List<TestUser> allUsers = listAllOrgUsers(orgId);
        
        List<TestRelationOrgGroupUser> relations;
        if (orgGroupId == null) {
        	relations = new LinkedList<>();
        } else {
        	relations = listRelations(orgId, orgGroupId, null);
        }
        
        List<RelationOrgGroupUserVo> vos = new LinkedList<>();
        for (TestUser user : allUsers) {
        	RelationOrgGroupUserVo vo = genVo(orgId, orgGroupId, user.getId());
        	
        	vo.setSelected(false);
        	vo.setSelecting(false);
        	for (TestRelationOrgGroupUser po : relations) {
        		if (po.getUserId().longValue() == user.getId().longValue()
						&& po.getOrgGroupId().longValue() == orgGroupId.longValue()) {
            		vo.setSelected(true);
            		vo.setSelecting(true);
            	}
        	}
        	vos.add(vo);
        }
        
		return vos;
	}

	private List<TestRelationOrgGroupUser> listRelations(Long orgId, Long orgGroupId, Long userId) {
		DetachedCriteria dc2 = DetachedCriteria.forClass(TestRelationOrgGroupUser.class);
		if (orgId != null) {
        	dc2.add(Restrictions.eq("orgId", orgId));
        }
		// 以下2个条件只会有一个
        if (orgGroupId != null) {
        	dc2.add(Restrictions.eq("orgGroupId", orgGroupId));
        }
        if (userId != null) {
        	dc2.add(Restrictions.eq("userId", userId));
        }
        
        dc2.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc2.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc2.addOrder(Order.asc("id"));
        List<TestRelationOrgGroupUser> relations = findAllByCriteria(dc2);
        
		return relations;
	}

	private List<TestOrgGroup> listAllOrgGroups(Long orgId) {
		DetachedCriteria dc = DetachedCriteria.forClass(TestOrgGroup.class);
        dc.add(Restrictions.eq("orgId", orgId));

        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.addOrder(Order.asc("id"));
        List<TestOrgGroup> ls = findAllByCriteria(dc);
        
		return ls;
	}
	
	private List<TestUser> listAllOrgUsers(Long orgId) {
		DetachedCriteria dc = DetachedCriteria.forClass(TestUser.class);
        
        dc.createAlias("orgSet", "orgs");
        dc.add(Restrictions.eq("orgs.id", orgId));
        
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        
        dc.addOrder(Order.asc("id"));
        List<TestUser> ls = findAllByCriteria(dc);
		
		return ls;
	}

	@Override
	public boolean saveRelations(List<RelationOrgGroupUserVo> orgGroupUserVos) {
		if (orgGroupUserVos == null) {
			return false;
		}
		for (Object obj: orgGroupUserVos) {
			RelationOrgGroupUserVo vo = JSON.parseObject(JSON.toJSONString(obj), RelationOrgGroupUserVo.class);
			if (vo.getSelecting() != vo.getSelected()) { // 变化了
				TestRelationOrgGroupUser relationOrgGroupUser = this.getRelationOrgGroupUser(vo.getOrgGroupId(), vo.getUserId());
				
    			if (vo.getSelecting() && relationOrgGroupUser == null) { // 勾选
    				relationOrgGroupUser = new TestRelationOrgGroupUser(vo.getOrgId(), vo.getOrgGroupId(), vo.getUserId());
    				saveOrUpdate(relationOrgGroupUser);
    			} else if (relationOrgGroupUser != null) { // 取消
    				getDao().delete(relationOrgGroupUser);
    			}
			}
		}
		
		return true;
	}

	private TestRelationOrgGroupUser getRelationOrgGroupUser(Long orgGroupId, Long userId) {
		DetachedCriteria dc = DetachedCriteria.forClass(TestRelationOrgGroupUser.class);
        dc.add(Restrictions.eq("orgGroupId", orgGroupId));
        dc.add(Restrictions.eq("userId", userId));
        
        dc.addOrder(Order.asc("id"));
        List<TestRelationOrgGroupUser> ls = findAllByCriteria(dc);
        
        if (ls.size() == 0) {
        	return null;
        }
		return ls.get(0);
	}
	
	private RelationOrgGroupUserVo genVo(Long orgId, Long orgGroupId, Long userId) {
		
		RelationOrgGroupUserVo vo = new RelationOrgGroupUserVo();
		vo.setOrgId(orgId);
		
		if (orgGroupId != null) {
			TestOrgGroup orgGroup = (TestOrgGroup) get(TestOrgGroup.class, orgGroupId);
			vo.setOrgGroupId(orgGroupId);
			vo.setOrgGroupName(orgGroup.getName());
		}
		
		if (userId != null) {
			TestUser user = (TestUser) get(TestUser.class, userId);
			vo.setUserId(user.getId());
			vo.setUserName(user.getName());
		}
		
		return vo;
	}
}
