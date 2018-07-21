package com.ngtesting.platform.service;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.model.TstProjectRoleEntityRelation;

import java.util.List;

public interface ProjectRoleEntityRelationService extends BaseService {

	List<TstProjectRoleEntityRelation> listRelationProjectRoleEntitys(Integer projectRoleId);
    List<TstProjectRoleEntityRelation> listByProject(Integer projectId);

	TstProjectRoleEntityRelation getRelationProjectRoleEntity(Integer projectRoleId);
    TstProjectRoleEntityRelation getByProjectAndEntityId(Integer projectId, Integer usereId);
//    void addUserToProjectAsLeaderPers(Integer userId, Integer roleId, Integer projectId);

    List<TstProjectRoleEntityRelation> batchSavePers(JSONObject json, Integer orgId);
    List<TstProjectRoleEntityRelation> changeRolePers(JSONObject json);

    String getEntityName(TstProjectRoleEntityRelation po);
}
