package com.ngtesting.platform.service;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.model.TstCaseComments;
import com.ngtesting.platform.model.TstUser;

public interface CaseCommentsService extends BaseService {

	TstCaseComments save(JSONObject vo, TstUser userVo);
	Boolean delete(Integer id, TstUser user);

    TstCaseComments genVo(TstCaseComments po);
}
