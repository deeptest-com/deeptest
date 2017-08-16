package com.ngtesting.platform.service;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.TestRelationProjectRoleEntity;
import com.ngtesting.platform.vo.RelationProjectRoleEntityVo;

import java.util.List;

public interface RelationProjectRoleEntityService extends BaseService {

	List<TestRelationProjectRoleEntity> listRelationProjectRoleEntitys(Long projectRoleId);
    List<TestRelationProjectRoleEntity> listByProject(Long projectId);

	TestRelationProjectRoleEntity getRelationProjectRoleEntity(Long projectRoleId);
    TestRelationProjectRoleEntity getByProjectAndEntityId(Long projectId, Long usereId);

    List<TestRelationProjectRoleEntity> batchSavePers(JSONObject json);
    List<TestRelationProjectRoleEntity> changeRolePers(JSONObject json);

    RelationProjectRoleEntityVo genVo(TestRelationProjectRoleEntity po);
    List<RelationProjectRoleEntityVo> genVos(List<TestRelationProjectRoleEntity> po);

}
