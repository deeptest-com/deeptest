package com.ngtesting.platform.service.intf;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.model.TstCaseInTaskComments;
import com.ngtesting.platform.model.TstUser;

public interface CaseInTaskCommentsService extends BaseService {

    TstCaseInTaskComments save(JSONObject vo, TstUser userVo);
    TstCaseInTaskComments save(TstCaseInTaskComments vo, TstUser user);

    Boolean delete(Integer id, TstUser user);
}
