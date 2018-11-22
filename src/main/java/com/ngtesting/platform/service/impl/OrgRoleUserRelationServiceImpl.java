package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.ngtesting.platform.dao.OrgRoleUserRelationDao;
import com.ngtesting.platform.model.TstOrgRole;
import com.ngtesting.platform.model.TstOrgRoleUserRelation;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.OrgRoleService;
import com.ngtesting.platform.service.intf.OrgRoleUserRelationService;
import com.ngtesting.platform.service.intf.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;

@Service
public class OrgRoleUserRelationServiceImpl extends BaseServiceImpl implements OrgRoleUserRelationService {
	@Autowired
	UserService userService;
    @Autowired
    OrgRoleService orgRoleService;

	@Autowired
	OrgRoleUserRelationDao orgRoleUserRelationDao;

	@Override
	public List<TstOrgRoleUserRelation> listRelationsByOrgRole(Integer orgId, Integer orgRoleId) {
        List<TstUser> allUsers = userService.listAllOrgUsers(orgId);

        List<TstOrgRoleUserRelation> relations;
        if (orgRoleId == null) {
			relations = new LinkedList<>();
        } else {
			relations = orgRoleUserRelationDao.query(orgId, orgRoleId, null);
        }

		List<TstOrgRoleUserRelation> vos = new LinkedList<>();
		for (TstUser user : allUsers) {
			TstOrgRoleUserRelation vo = genVo(orgId, user, orgRoleId);

			vo.setSelected(false);
			vo.setSelecting(false);
			for (TstOrgRoleUserRelation po : relations) {
				if (po.getUserId().longValue() == user.getId().longValue()
						&& po.getOrgRoleId().longValue() == orgRoleId.longValue()) {
					vo.setSelected(true);
					vo.setSelecting(true);
				}
			}
			vos.add(vo);
		}

		return vos;
	}

    @Override
    public List<TstOrgRoleUserRelation> listRelationsByUser(Integer orgId, Integer userId) {
        List<TstOrgRole> allRoles = orgRoleService.listAllOrgRoles(orgId);

        List<TstOrgRoleUserRelation> relations;
        if (userId == null) {
            relations = new LinkedList<>();
        } else {
            relations = orgRoleUserRelationDao.query(orgId, null, userId);
        }

        List<TstOrgRoleUserRelation> vos = new LinkedList<>();
        for (TstOrgRole role : allRoles) {
            TstOrgRoleUserRelation vo = genVo(orgId, role, userId);

            vo.setSelected(false);
            vo.setSelecting(false);
            for (TstOrgRoleUserRelation po : relations) {
                if (po.getOrgRoleId().longValue() == role.getId().longValue()
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
    public boolean saveRelationsForUser(Integer orgId, Integer userId, List<TstOrgRoleUserRelation> relations) {
        if (relations == null) {
            return false;
        }

        List<TstOrgRoleUserRelation> selectedList = new LinkedList<>();
        for (Object obj: relations) {
            TstOrgRoleUserRelation vo = JSON.parseObject(JSON.toJSONString(obj), TstOrgRoleUserRelation.class);
            if (vo.getSelecting()) {
                vo.setUserId(userId);
                selectedList.add(vo);
            }
        }

        orgRoleUserRelationDao.removeAllRolesForUser(orgId, userId);
        if (selectedList.size() > 0) {
            orgRoleUserRelationDao.saveRelations(selectedList);
        }

        return true;
    }

    @Override
    public boolean saveRelationsForRole(Integer orgId, Integer roleId, List<TstOrgRoleUserRelation> relations) {
        if (relations == null) {
            return false;
        }

        List<TstOrgRoleUserRelation> selectedList = new LinkedList<>();
        for (Object obj: relations) {
            TstOrgRoleUserRelation vo = JSON.parseObject(JSON.toJSONString(obj), TstOrgRoleUserRelation.class);
            if (vo.getSelecting()) {
                vo.setOrgRoleId(roleId);
                selectedList.add(vo);
            }
        }

        orgRoleUserRelationDao.removeAllUsersForRole(orgId, roleId);
        if (selectedList.size() > 0) {
            orgRoleUserRelationDao.saveRelations(selectedList);
        }

        return true;
    }

    @Override
    public TstOrgRoleUserRelation genVo(Integer orgId, TstUser user, Integer orgRoleId) {
        TstOrgRoleUserRelation vo = new TstOrgRoleUserRelation();
        vo.setOrgId(orgId);
        vo.setOrgRoleId(orgRoleId);

        vo.setUserId(user.getId());
        vo.setUserName(user.getNickname());

        return vo;
    }

    @Override
    public TstOrgRoleUserRelation genVo(Integer orgId, TstOrgRole role, Integer userId) {
        TstOrgRoleUserRelation vo = new TstOrgRoleUserRelation();
        vo.setOrgId(orgId);
        vo.setUserId(userId);

        vo.setOrgRoleId(role.getId());
        vo.setOrgRoleName(role.getName());

        return vo;
    }

}
