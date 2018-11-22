package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.ngtesting.platform.dao.OrgRolePrivilegeRelationDao;
import com.ngtesting.platform.model.TstOrgPrivilegeDefine;
import com.ngtesting.platform.model.TstOrgRole;
import com.ngtesting.platform.model.TstOrgRolePrivilegeRelation;
import com.ngtesting.platform.service.intf.OrgPrivilegeService;
import com.ngtesting.platform.service.intf.OrgRolePrivilegeRelationService;
import com.ngtesting.platform.service.intf.OrgRoleService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;

@Service
public class OrgRolePrivilegeRelationServiceImpl extends BaseServiceImpl implements OrgRolePrivilegeRelationService {
	@Autowired
	private OrgRolePrivilegeRelationDao orgRolePrivilegeRelationDao;

    @Autowired
    OrgPrivilegeService orgPrivilegeService;
    @Autowired
    OrgRoleService orgRoleService;

	@Override
	public List<TstOrgRolePrivilegeRelation> listRelationsByOrgRole(Integer orgId, Integer orgRoleId) {
        List<TstOrgPrivilegeDefine> allPrivileges = orgPrivilegeService.listAllOrgPrivileges();

        List<TstOrgRolePrivilegeRelation> relations;
        if (orgRoleId == null) {
			relations = new LinkedList<>();
        } else {
			relations = orgRolePrivilegeRelationDao.query(orgId, orgRoleId, null);
        }

		List<TstOrgRolePrivilegeRelation> vos = new LinkedList<>();
		for (TstOrgPrivilegeDefine priv : allPrivileges) {
			TstOrgRolePrivilegeRelation vo = genVo(orgId, priv, orgRoleId);

			vo.setSelected(false);
			vo.setSelecting(false);
			for (TstOrgRolePrivilegeRelation po : relations) {
				if (po.getOrgPrivilegeId().longValue() == priv.getId().longValue()
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
    public List<TstOrgRolePrivilegeRelation> listRelationsByPrivilege(Integer orgId, Integer privilegeId) {
        List<TstOrgRole> allRoles = orgRoleService.listAllOrgRoles(orgId);

        List<TstOrgRolePrivilegeRelation> relations;
        if (privilegeId == null) {
            relations = new LinkedList<>();
        } else {
            relations = orgRolePrivilegeRelationDao.query(orgId, null, privilegeId);
        }

        List<TstOrgRolePrivilegeRelation> vos = new LinkedList<>();
        for (TstOrgRole role : allRoles) {
            TstOrgRolePrivilegeRelation vo = genVo(orgId, role, privilegeId);

            vo.setSelected(false);
            vo.setSelecting(false);
            for (TstOrgRolePrivilegeRelation po : relations) {
                if (po.getOrgRoleId().longValue() == role.getId().longValue()
                        && po.getOrgPrivilegeId().longValue() == privilegeId.longValue()) {
                    vo.setSelected(true);
                    vo.setSelecting(true);
                }
            }
            vos.add(vo);
        }

        return vos;
    }

    @Override
    public boolean saveRelationsForRole(Integer orgId, Integer roleId, List<TstOrgRolePrivilegeRelation> relations) {
        if (relations == null) {
            return false;
        }

        List<TstOrgRolePrivilegeRelation> selectedList = new LinkedList<>();
        for (Object obj: relations) {
            TstOrgRolePrivilegeRelation vo = JSON.parseObject(JSON.toJSONString(obj), TstOrgRolePrivilegeRelation.class);
            if (vo.getSelecting()) {
                vo.setOrgRoleId(roleId);
                selectedList.add(vo);
            }
        }

        orgRolePrivilegeRelationDao.removeAllPrivilegesForRole(orgId, roleId);
        if (selectedList.size() > 0) {
            orgRolePrivilegeRelationDao.saveRelations(selectedList);
        }

        return true;
    }

    @Override
    public boolean saveRelationsForPrivilege(Integer orgId, Integer privilegeId, List<TstOrgRolePrivilegeRelation> relations) {
        if (relations == null) {
            return false;
        }

        List<TstOrgRolePrivilegeRelation> selectedList = new LinkedList<>();
        for (Object obj: relations) {
            TstOrgRolePrivilegeRelation vo = JSON.parseObject(JSON.toJSONString(obj), TstOrgRolePrivilegeRelation.class);
            if (vo.getSelecting()) {
                vo.setOrgPrivilegeId(privilegeId);
                selectedList.add(vo);
            }
        }

        orgRolePrivilegeRelationDao.removeAllRolesForPrivilege(orgId, privilegeId);
        if (selectedList.size() > 0) {
            orgRolePrivilegeRelationDao.saveRelations(selectedList);
        }

        return true;
    }

    @Override
    public TstOrgRolePrivilegeRelation genVo(Integer orgId, TstOrgPrivilegeDefine priv, Integer orgRoleId) {
        TstOrgRolePrivilegeRelation vo = new TstOrgRolePrivilegeRelation();
        vo.setOrgId(orgId);
        vo.setOrgRoleId(orgRoleId);

        vo.setOrgPrivilegeId(priv.getId());
        vo.setOrgPrivilegeName(priv.getName());

        return vo;
    }

    @Override
    public TstOrgRolePrivilegeRelation genVo(Integer orgId, TstOrgRole role, Integer privilegeId) {
        TstOrgRolePrivilegeRelation vo = new TstOrgRolePrivilegeRelation();
        vo.setOrgId(orgId);
        vo.setOrgPrivilegeId(privilegeId);

        vo.setOrgRoleId(role.getId());
        vo.setOrgRoleName(role.getName());

        return vo;
    }

}
