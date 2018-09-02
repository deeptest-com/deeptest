package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.ngtesting.platform.dao.OrgRoleGroupRelationDao;
import com.ngtesting.platform.model.TstOrgGroup;
import com.ngtesting.platform.model.TstOrgRole;
import com.ngtesting.platform.model.TstOrgRoleGroupRelation;
import com.ngtesting.platform.service.OrgGroupService;
import com.ngtesting.platform.service.OrgRoleGroupRelationService;
import com.ngtesting.platform.service.OrgRoleService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;

@Service
public class OrgRoleGroupRelationServiceImpl extends BaseServiceImpl implements OrgRoleGroupRelationService {
	@Autowired
	OrgGroupService orgGroupService;
    @Autowired
    OrgRoleService orgRoleService;

	@Autowired
    OrgRoleGroupRelationDao orgRoleGroupRelationDao;

	@Override
	public List<TstOrgRoleGroupRelation> listRelationsByOrgRole(Integer orgId, Integer orgRoleId) {
        List<TstOrgGroup> allGroups = orgGroupService.list(orgId);

        List<TstOrgRoleGroupRelation> relations;
        if (orgRoleId == null) {
			relations = new LinkedList<>();
        } else {
			relations = orgRoleGroupRelationDao.query(orgId, orgRoleId, null);
        }

		List<TstOrgRoleGroupRelation> vos = new LinkedList<>();
		for (TstOrgGroup group : allGroups) {
			TstOrgRoleGroupRelation vo = genVo(orgId, group, orgRoleId);

			vo.setSelected(false);
			vo.setSelecting(false);
			for (TstOrgRoleGroupRelation po : relations) {
				if (po.getGroupId().longValue() == group.getId().longValue()
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
    public List<TstOrgRoleGroupRelation> listRelationsByGroup(Integer orgId, Integer groupId) {
        List<TstOrgRole> allRoles = orgRoleService.listAllOrgRoles(orgId);

        List<TstOrgRoleGroupRelation> relations;
        if (groupId == null) {
            relations = new LinkedList<>();
        } else {
            relations = orgRoleGroupRelationDao.query(orgId, null, groupId);
        }

        List<TstOrgRoleGroupRelation> vos = new LinkedList<>();
        for (TstOrgRole role : allRoles) {
            TstOrgRoleGroupRelation vo = genVo(orgId, role, groupId);

            vo.setSelected(false);
            vo.setSelecting(false);
            for (TstOrgRoleGroupRelation po : relations) {
                if (po.getOrgRoleId().longValue() == role.getId().longValue()
                        && po.getGroupId().longValue() == groupId.longValue()) {
                    vo.setSelected(true);
                    vo.setSelecting(true);
                }
            }
            vos.add(vo);
        }

        return vos;
    }

    @Override
    public boolean saveRelationsForGroup(Integer orgId, Integer groupId, List<TstOrgRoleGroupRelation> relations) {
        if (relations == null) {
            return false;
        }

        List<TstOrgRoleGroupRelation> selectedList = new LinkedList<>();
        for (Object obj: relations) {
            TstOrgRoleGroupRelation vo = JSON.parseObject(JSON.toJSONString(obj), TstOrgRoleGroupRelation.class);
            if (vo.getSelecting()) {
                vo.setGroupId(groupId);
                selectedList.add(vo);
            }
        }

        orgRoleGroupRelationDao.removeAllRolesForGroup(orgId, groupId);
        if (selectedList.size() > 0) {
            orgRoleGroupRelationDao.saveRelations(selectedList);
        }

        return true;
    }

    @Override
    public boolean saveRelationsForRole(Integer orgId, Integer roleId, List<TstOrgRoleGroupRelation> relations) {
        if (relations == null) {
            return false;
        }

        List<TstOrgRoleGroupRelation> selectedList = new LinkedList<>();
        for (Object obj: relations) {
            TstOrgRoleGroupRelation vo = JSON.parseObject(JSON.toJSONString(obj), TstOrgRoleGroupRelation.class);
            if (vo.getSelecting()) {
                vo.setOrgRoleId(roleId);
                selectedList.add(vo);
            }
        }

        orgRoleGroupRelationDao.removeAllGroupsForRole(orgId, roleId);
        if (selectedList.size() > 0) {
            orgRoleGroupRelationDao.saveRelations(selectedList);
        }

        return true;
    }

    @Override
    public TstOrgRoleGroupRelation genVo(Integer orgId, TstOrgGroup group, Integer orgRoleId) {
        TstOrgRoleGroupRelation vo = new TstOrgRoleGroupRelation();
        vo.setOrgId(orgId);
        vo.setOrgRoleId(orgRoleId);

        vo.setGroupId(group.getId());
        vo.setGroupName(group.getName());

        return vo;
    }

    @Override
    public TstOrgRoleGroupRelation genVo(Integer orgId, TstOrgRole role, Integer groupId) {
        TstOrgRoleGroupRelation vo = new TstOrgRoleGroupRelation();
        vo.setOrgId(orgId);
        vo.setGroupId(groupId);

        vo.setOrgRoleId(role.getId());
        vo.setOrgRoleName(role.getName());

        return vo;
    }

}
