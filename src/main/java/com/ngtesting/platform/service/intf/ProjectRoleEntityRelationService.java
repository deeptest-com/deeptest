package com.ngtesting.platform.service.intf;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.model.TstProjectRoleEntityRelation;
import com.ngtesting.platform.model.TstUser;

import java.util.List;

public interface ProjectRoleEntityRelationService extends BaseService {

    List<TstProjectRoleEntityRelation> listByProject(Integer projectId);

    List<TstProjectRoleEntityRelation> batchSavePers(JSONObject json, TstUser user);

    List<TstProjectRoleEntityRelation> changeRolePers(JSONObject json, TstUser user);
    List<TstProjectRoleEntityRelation> remove(String type, Integer entityId, TstUser user);
}
