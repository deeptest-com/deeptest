package com.ngtesting.platform.service.inf;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.model.TstRelationProjectRoleEntity;

import java.util.List;

public interface RelationProjectRoleEntityService extends BaseService {

	List<TstRelationProjectRoleEntity> listRelationProjectRoleEntitys(Integer projectRoleId);
    List<TstRelationProjectRoleEntity> listByProject(Integer projectId);

	TstRelationProjectRoleEntity getRelationProjectRoleEntity(Integer projectRoleId);
    TstRelationProjectRoleEntity getByProjectAndEntityId(Integer projectId, Integer usereId);
//    void addUserToProjectAsLeaderPers(Integer userId, Integer roleId, Integer projectId);

    List<TstRelationProjectRoleEntity> batchSavePers(JSONObject json);
    List<TstRelationProjectRoleEntity> changeRolePers(JSONObject json);

    String getEntityName(TstRelationProjectRoleEntity po);
}
