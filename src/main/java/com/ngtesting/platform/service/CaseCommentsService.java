package com.ngtesting.platform.service;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.TestCaseComments;
import com.ngtesting.platform.vo.TestCaseCommentsVo;
import com.ngtesting.platform.vo.UserVo;

public interface CaseCommentsService extends BaseService {

	TestCaseCommentsVo save(JSONObject vo, UserVo userVo);
	boolean delete(Long d, Long userId);

    TestCaseCommentsVo genVo(TestCaseComments po);
}
