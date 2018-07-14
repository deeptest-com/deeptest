package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.model.TstOrgGroup;
import com.ngtesting.platform.model.TstOrgGroupUserRelation;
import com.ngtesting.platform.service.OrgGroupUserRelationService;
import com.ngtesting.platform.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class OrgGroupUserRelationServiceImpl extends BaseServiceImpl implements OrgGroupUserRelationService {

	@Autowired
	UserService userService;

	@Override
	public List<TstOrgGroupUserRelation> listRelationsByUser(Integer orgId, Integer userId) {

//        List<TestOrgGroup> allOrgGroups = listAllOrgGroups(orgId);
//
//        List<TestRelationOrgGroupUser> relations;
//        if (userId == null) {
//        	relations = new LinkedList<>();
//        } else {
//        	relations = listRelations(orgId, null, userId);
//        }
//
//        List<TstOrgGroupUserRelation> vos = new LinkedList<>();
//        for (TestOrgGroup orgGroup : allOrgGroups) {
//        	TstOrgGroupUserRelation vo = genVo(orgId, orgGroup.getId(), userId);
//
//        	vo.setSelected(false);
//        	vo.setSelecting(false);
//        	for (TestRelationOrgGroupUser po : relations) {
//        		if (po.getOrgGroupId() == orgGroup.getId() && po.getUserId() == userId) {
//            		vo.setSelected(true);
//            		vo.setSelecting(true);
//            	}
//        	}
//        	vos.add(vo);
//        }
//
//		return vos;

		return null;
	}

	@Override
	public List<TstOrgGroupUserRelation> listRelationsByGroup(Integer orgId, Integer orgGroupId) {

//        List<TestUser> allUsers = userService.listAllOrgUsers(orgId);
//
//        List<TestRelationOrgGroupUser> relations;
//        if (orgGroupId == null) {
//        	relations = new LinkedList<>();
//        } else {
//        	relations = listRelations(orgId, orgGroupId, null);
//        }
//
//        List<TstOrgGroupUserRelation> vos = new LinkedList<>();
//        for (TestUser user : allUsers) {
//        	TstOrgGroupUserRelation vo = genVo(orgId, orgGroupId, user.getId());
//
//        	vo.setSelected(false);
//        	vo.setSelecting(false);
//        	for (TestRelationOrgGroupUser po : relations) {
//        		if (po.getUserId().longValue() == user.getId().longValue()
//						&& po.getOrgGroupId().longValue() == orgGroupId.longValue()) {
//            		vo.setSelected(true);
//            		vo.setSelecting(true);
//            	}
//        	}
//        	vos.add(vo);
//        }
//
//		return vos;

		return null;
	}

	private List<TstOrgGroupUserRelation> listRelations(Integer orgId, Integer orgGroupId, Integer userId) {
//		DetachedCriteria dc2 = DetachedCriteria.forClass(TestRelationOrgGroupUser.class);
//		if (orgId != null) {
//        	dc2.add(Restrictions.eq("orgId", orgId));
//        }
//		// 以下2个条件只会有一个
//        if (orgGroupId != null) {
//        	dc2.add(Restrictions.eq("orgGroupId", orgGroupId));
//        }
//        if (userId != null) {
//        	dc2.add(Restrictions.eq("userId", userId));
//        }
//
//        dc2.add(Restrictions.eq("deleted", Boolean.FALSE));
//        dc2.add(Restrictions.eq("disabled", Boolean.FALSE));
//        dc2.addOrder(Order.asc("id"));
//        List<TestRelationOrgGroupUser> relations = findAllByCriteria(dc2);
//
//		return relations;

		return null;
	}

	private List<TstOrgGroup> listAllOrgGroups(Integer orgId) {
//		DetachedCriteria dc = DetachedCriteria.forClass(TestOrgGroup.class);
//        dc.add(Restrictions.eq("orgId", orgId));
//
//        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//        dc.addOrder(Order.asc("id"));
//        List<TestOrgGroup> ls = findAllByCriteria(dc);
//
//		return ls;

		return null;
	}

	@Override
	public boolean saveRelations(List<TstOrgGroupUserRelation> orgGroupTstUsers) {
//		return saveRelations(null, orgGroupTstUsers);

		return true;
	}
	@Override
	public boolean saveRelations(Integer userId, List<TstOrgGroupUserRelation> orgGroupTstUsers) {
//		if (orgGroupTstUsers == null) {
//			return false;
//		}
//		for (Object obj: orgGroupTstUsers) {
//			TstOrgGroupUserRelation vo = JSON.parseObject(JSON.toJSONString(obj), TstOrgGroupUserRelation.class);
//			if (vo.getSelecting() != vo.getSelected()) { // 变化了
//				TestRelationOrgGroupUser relationOrgGroupUser = this.getRelationOrgGroupUser(vo.getOrgGroupId(), vo.getUserId());
//
//				if (vo.getSelecting() && relationOrgGroupUser == null) { // 勾选
//					relationOrgGroupUser = new TestRelationOrgGroupUser(vo.getOrgId(), vo.getOrgGroupId(), vo.getUserId());
//					if (relationOrgGroupUser.getUserId() == null) {
//						relationOrgGroupUser.setUserId(userId);
//					}
//					saveOrUpdate(relationOrgGroupUser);
//				} else if (relationOrgGroupUser != null) { // 取消
//					getDao().delete(relationOrgGroupUser);
//				}
//			}
//		}

		return true;
	}

	private TstOrgGroupUserRelation getRelationOrgGroupUser(Integer orgGroupId, Integer userId) {
//		DetachedCriteria dc = DetachedCriteria.forClass(TestRelationOrgGroupUser.class);
//        dc.add(Restrictions.eq("orgGroupId", orgGroupId));
//        dc.add(Restrictions.eq("userId", userId));
//
//        dc.addOrder(Order.asc("id"));
//        List<TestRelationOrgGroupUser> ls = findAllByCriteria(dc);
//
//        if (ls.size() == 0) {
//        	return null;
//        }
//		return ls.get(0);

		return null;
	}

	private TstOrgGroupUserRelation genVo(Integer orgId, Integer orgGroupId, Integer userId) {

//		TstOrgGroupUserRelation vo = new TstOrgGroupUserRelation();
//		vo.setOrgId(orgId);
//
//		if (orgGroupId != null) {
//			TestOrgGroup orgGroup = (TestOrgGroup) get(TestOrgGroup.class, orgGroupId);
//			vo.setOrgGroupId(orgGroupId);
//			vo.setOrgGroupName(orgGroup.getName());
//		}
//
//		if (userId != null) {
//			TestUser user = (TestUser) get(TestUser.class, userId);
//			vo.setUserId(user.getId());
//			vo.setUserName(user.getName());
//		}
//
//		return vo;

		return null;
	}
}
