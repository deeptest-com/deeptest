package com.ngtesting.platform.service;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.model.TstProjectRoleEntityRelation;

import java.util.List;

public interface ProjectRoleEntityRelationService extends BaseService {

    List<TstProjectRoleEntityRelation> listByProject(Integer projectId);

    List<TstProjectRoleEntityRelation> batchSavePers(JSONObject json, Integer orgId);
    List<TstProjectRoleEntityRelation> changeRolePers(JSONObject json);
}
