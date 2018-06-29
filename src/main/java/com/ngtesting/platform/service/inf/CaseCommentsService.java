package com.ngtesting.platform.service.inf;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.model.TstCaseComments;
import com.ngtesting.platform.model.TstUser;

public interface CaseCommentsService extends BaseService {

	TstCaseComments save(JSONObject vo, TstUser userVo);
	boolean delete(Long d, Long userId);

    TstCaseComments genVo(TstCaseComments po);
}
