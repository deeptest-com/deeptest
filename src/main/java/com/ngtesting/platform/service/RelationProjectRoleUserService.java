package com.ngtesting.platform.service;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.TestRelationProjectRoleUser;
import com.ngtesting.platform.vo.RelationProjectRoleUserVo;

import java.util.List;

public interface RelationProjectRoleUserService extends BaseService {

	List<TestRelationProjectRoleUser> listRelationProjectRoleUsers(Long projectRoleId);
    List<TestRelationProjectRoleUser> listByProject(Long projectId);

	TestRelationProjectRoleUser getRelationProjectRoleUser(Long projectRoleId);
    TestRelationProjectRoleUser getByUserId(Long usereId);

    List<TestRelationProjectRoleUser> batchSavePers(JSONObject json);

    RelationProjectRoleUserVo genVo(TestRelationProjectRoleUser po);
    List<RelationProjectRoleUserVo> genVos(List<TestRelationProjectRoleUser> po);
}
