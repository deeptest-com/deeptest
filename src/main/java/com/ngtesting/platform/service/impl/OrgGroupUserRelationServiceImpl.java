package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.OrgGroupUserRelationDao;
import com.ngtesting.platform.model.TstOrgGroup;
import com.ngtesting.platform.model.TstOrgGroupUserRelation;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.OrgGroupService;
import com.ngtesting.platform.service.OrgGroupUserRelationService;
import com.ngtesting.platform.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;

@Service
public class OrgGroupUserRelationServiceImpl extends BaseServiceImpl implements OrgGroupUserRelationService {

	@Autowired
    OrgGroupService orgGroupService;
    @Autowired
    UserService userService;
    @Autowired
    OrgGroupUserRelationDao orgGroupUserRelationDao;

	@Override
	public List<TstOrgGroupUserRelation> listRelationsByUser(Integer orgId, Integer userId) {

        List<TstOrgGroup> allOrgGroups = listAllOrgGroups(orgId);

        List<TstOrgGroupUserRelation> relations;
        if (userId == null) {
        	relations = new LinkedList<>();
        } else {
        	relations = listRelations(orgId, null, userId);
        }

        List<TstOrgGroupUserRelation> vos = new LinkedList<>();
        for (TstOrgGroup orgGroup : allOrgGroups) {
        	TstOrgGroupUserRelation vo = genVo(orgId, orgGroup, userId);

        	vo.setSelected(false);
        	vo.setSelecting(false);
        	for (TstOrgGroupUserRelation po : relations) {
        		if (po.getOrgGroupId().longValue() == orgGroup.getId().longValue()
						&& po.getUserId().longValue() == userId.longValue()) {
            		vo.setSelected(true);
            		vo.setSelecting(true);
            	}
        	}
        	vos.add(vo);
        }

		return vos;
	}

	@Override
	public List<TstOrgGroupUserRelation> listRelationsByGroup(Integer orgId, Integer groupId) {

        List<TstUser> allOrgUsers = listAllOrgUsers(orgId);

        List<TstOrgGroupUserRelation> relations;
        if (groupId == null) {
            relations = new LinkedList<>();
        } else {
            relations = listRelations(orgId, groupId, null);
        }

        List<TstOrgGroupUserRelation> vos = new LinkedList<>();
        for (TstUser user : allOrgUsers) {
            TstOrgGroupUserRelation vo = genVo(orgId, user, groupId);

            vo.setSelected(false);
            vo.setSelecting(false);
            for (TstOrgGroupUserRelation po : relations) {
                if (po.getUserId().longValue() == user.getId().longValue()
                        && po.getOrgGroupId().longValue() == groupId.longValue()) {
                    vo.setSelected(true);
                    vo.setSelecting(true);
                }
            }
            vos.add(vo);
        }

        return vos;
	}

    @Override
    public List<TstOrgGroupUserRelation> listRelations(Integer orgId, Integer groupId, Integer userId) {
        List<TstOrgGroupUserRelation> ls = orgGroupUserRelationDao.query(orgId, groupId, userId);

        return ls;
	}

    @Override
    public List<TstOrgGroup> listAllOrgGroups(Integer orgId) {
        List<TstOrgGroup> ls = orgGroupService.search(orgId, null, null);

		return ls;
	}
    @Override
    public List<TstUser> listAllOrgUsers(Integer orgId) {
        List<TstUser> ls = userService.search(orgId, null, null);

        return ls;
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

    @Override
    public TstOrgGroupUserRelation getRelationOrgGroupUser(Integer orgGroupId, Integer userId) {
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

    @Override
    public TstOrgGroupUserRelation genVo(Integer orgId, TstOrgGroup group, Integer userId) {

		TstOrgGroupUserRelation vo = new TstOrgGroupUserRelation();
		vo.setOrgId(orgId);

        vo.setOrgGroupId(group.getId());
        vo.setOrgGroupName(group.getName());

        vo.setUserId(userId);

		return vo;
	}

    @Override
    public TstOrgGroupUserRelation genVo(Integer orgId, TstUser user, Integer groupId) {
        TstOrgGroupUserRelation vo = new TstOrgGroupUserRelation();
        vo.setOrgId(orgId);

        vo.setOrgGroupId(groupId);

        vo.setUserId(user.getId());
        vo.setUserName(user.getNickname());

        return vo;
    }
}
