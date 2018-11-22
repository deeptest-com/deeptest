package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.model.TstOrgPrivilegeDefine;
import com.ngtesting.platform.model.TstOrgRole;
import com.ngtesting.platform.model.TstOrgRolePrivilegeRelation;

import java.util.List;

public interface OrgRolePrivilegeRelationService extends BaseService {

	List<TstOrgRolePrivilegeRelation> listRelationsByOrgRole(Integer orgId, Integer orgRoleId);

    List<TstOrgRolePrivilegeRelation> listRelationsByPrivilege(Integer orgId, Integer privilegeId);

    boolean saveRelationsForRole(Integer orgId, Integer roleId, List<TstOrgRolePrivilegeRelation> relations);
    boolean saveRelationsForPrivilege(Integer orgId, Integer privilegeId, List<TstOrgRolePrivilegeRelation> relations);

    TstOrgRolePrivilegeRelation genVo(Integer orgId, TstOrgRole role, Integer privilegeIdId);
    TstOrgRolePrivilegeRelation genVo(Integer orgId, TstOrgPrivilegeDefine priv, Integer orgRoleId);
}
