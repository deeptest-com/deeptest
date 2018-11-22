package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.ngtesting.platform.dao.OrgGroupUserRelationDao;
import com.ngtesting.platform.model.TstOrgGroup;
import com.ngtesting.platform.model.TstOrgGroupUserRelation;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.OrgGroupService;
import com.ngtesting.platform.service.intf.OrgGroupUserRelationService;
import com.ngtesting.platform.service.intf.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

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
    public List<TstOrgGroupUserRelation> listRelationsByGroup(Integer orgId, Integer groupId) {
        List<TstUser> allOrgUsers = listAllOrgUsers(orgId);

        List<TstOrgGroupUserRelation> relations;
        if (groupId == null) {
            relations = new LinkedList<>();
        } else {
            relations = orgGroupUserRelationDao.query(orgId, groupId, null);
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
	public List<TstOrgGroupUserRelation> listRelationsByUser(Integer orgId, Integer userId) {

        List<TstOrgGroup> allOrgGroups = listAllOrgGroups(orgId);

        List<TstOrgGroupUserRelation> relations;
        if (userId == null) {
        	relations = new LinkedList<>();
        } else {
        	relations = orgGroupUserRelationDao.query(orgId, null, userId);
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
    public List<TstOrgGroup> listAllOrgGroups(Integer orgId) {
        List<TstOrgGroup> ls = orgGroupService.list(orgId);

		return ls;
	}
    @Override
    public List<TstUser> listAllOrgUsers(Integer orgId) {
        List<TstUser> ls = userService.search(orgId, null, null);

        return ls;
    }

	@Override
    @Transactional
	public boolean saveRelationsForUser(Integer orgId, Integer userId, List<TstOrgGroupUserRelation> orgGroupUserRelation) {
		if (orgGroupUserRelation == null) {
            return false;
        }

        List<TstOrgGroupUserRelation> selectedList = new LinkedList<>();
		for (Object obj: orgGroupUserRelation) {
			TstOrgGroupUserRelation vo = JSON.parseObject(JSON.toJSONString(obj), TstOrgGroupUserRelation.class);
			if (vo.getSelecting()) {
			    vo.setUserId(userId);
                selectedList.add(vo);
            }
		}

        orgGroupUserRelationDao.removeAllGroupsForUser(orgId, userId);
		if (selectedList.size() > 0) {
            orgGroupUserRelationDao.saveRelations(selectedList);
        }

		return true;
	}

    @Override
    public boolean saveRelationsForGroup(Integer orgId, Integer groupId, List<TstOrgGroupUserRelation> orgGroupUserRelation) {
        if (orgGroupUserRelation == null) {
            return false;
        }

        List<TstOrgGroupUserRelation> selectedList = new LinkedList<>();
        for (Object obj: orgGroupUserRelation) {
            TstOrgGroupUserRelation vo = JSON.parseObject(JSON.toJSONString(obj), TstOrgGroupUserRelation.class);
            if (vo.getSelecting()) {
                vo.setOrgGroupId(groupId);
                selectedList.add(vo);
            }
        }

        orgGroupUserRelationDao.removeAllUsersForGroup(orgId, groupId);
        if (selectedList.size() > 0) {
            orgGroupUserRelationDao.saveRelations(selectedList);
        }

        return true;
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
