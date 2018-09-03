package com.ngtesting.platform.service;

import com.ngtesting.platform.model.TstOrgGroup;
import com.ngtesting.platform.model.TstOrgRole;
import com.ngtesting.platform.model.TstOrgRoleGroupRelation;

import java.util.List;

public interface OrgRoleGroupRelationService extends BaseService {

    List<TstOrgRoleGroupRelation> listRelationsByOrgRole(Integer orgId, Integer orgRoleId);
    List<TstOrgRoleGroupRelation> listRelationsByGroup(Integer orgId, Integer groupId);

    boolean saveRelationsForGroup(Integer orgId, Integer groupId, List<TstOrgRoleGroupRelation> relations);
    boolean saveRelationsForRole(Integer orgId, Integer roleId, List<TstOrgRoleGroupRelation> relations);

    TstOrgRoleGroupRelation genVo(Integer orgId, TstOrgGroup group, Integer orgRoleId);

    TstOrgRoleGroupRelation genVo(Integer orgId, TstOrgRole role, Integer groupId);

}
