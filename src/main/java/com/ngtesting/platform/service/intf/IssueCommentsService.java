package com.ngtesting.platform.service.intf;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.model.TstCaseComments;
import com.ngtesting.platform.model.TstUser;

public interface IssueCommentsService extends BaseService {

	TstCaseComments save(JSONObject vo, TstUser TstUser);
	boolean delete(Long d, Long userId);

    TstCaseComments genVo(TstCaseComments po);
}
