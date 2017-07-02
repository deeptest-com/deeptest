package com.ngtesting.platform.service;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.TestCase;
import com.ngtesting.platform.vo.TestCaseVo;

import java.util.List;

public interface CaseService extends BaseService {

	List<TestCase> query(Long sutieId);
	TestCaseVo getById(Long caseId);

	List<TestCaseVo> genVos(List<TestCase> pos);

	TestCaseVo genVo(TestCase po);
	
	TestCase delete(Long vo, Long userId);
	TestCase move(Long id, Long pid, Long prePid, Long userId);
	TestCase create(Long id, String value, String type, Long pid, Long userId);
	TestCase rename(Long id, String value, Long userId);
	TestCase save(JSONObject json);

}
