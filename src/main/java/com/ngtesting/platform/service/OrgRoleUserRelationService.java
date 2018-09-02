package com.ngtesting.platform.service;

import com.ngtesting.platform.model.TstOrgRole;
import com.ngtesting.platform.model.TstOrgRoleUserRelation;
import com.ngtesting.platform.model.TstUser;

import java.util.List;

public interface OrgRoleUserRelationService extends BaseService {

    List<TstOrgRoleUserRelation> listRelationsByOrgRole(Integer orgId, Integer orgRoleId);
    List<TstOrgRoleUserRelation> listRelationsByUser(Integer orgId, Integer userId);

    boolean saveRelationsForUser(Integer orgId, Integer userId, List<TstOrgRoleUserRelation> relations);
    boolean saveRelationsForRole(Integer orgId, Integer roleId, List<TstOrgRoleUserRelation> relations);

	TstOrgRoleUserRelation genVo(Integer orgId, TstUser user, Integer orgRoleId);

    TstOrgRoleUserRelation genVo(Integer orgId, TstOrgRole role, Integer userId);
}
