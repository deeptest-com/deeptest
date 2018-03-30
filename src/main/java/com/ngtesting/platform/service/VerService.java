package com.ngtesting.platform.service;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.TestVer;
import com.ngtesting.platform.vo.TestVerVo;
import com.ngtesting.platform.vo.UserVo;

import java.util.List;

public interface VerService extends BaseService {
	List<TestVer> list(Long projectId);
	TestVerVo getById(Long caseId);
	TestVer save(JSONObject json, UserVo optUser);
	TestVer delete(Long vo, Long userId);

	List<TestVerVo> genVos(List<TestVer> pos);
	TestVerVo genVo(TestVer po);
}
