package com.ngtesting.platform.service.intf;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.model.IsuComments;
import com.ngtesting.platform.model.TstUser;

public interface IssueCommentsService extends BaseService {
    IsuComments get(Integer id);

    IsuComments save(JSONObject vo, TstUser userVo);
	IsuComments save(IsuComments vo, TstUser user);

	Boolean delete(Integer id, TstUser user);
}
