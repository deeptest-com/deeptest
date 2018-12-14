package com.ngtesting.platform.service.intf;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.model.TstCaseComments;
import com.ngtesting.platform.model.TstUser;

public interface CaseCommentsService extends BaseService {

	TstCaseComments save(JSONObject vo, TstUser userVo);
    TstCaseComments save(TstCaseComments vo, TstUser user);

    Boolean delete(Integer id, TstUser user);
}
