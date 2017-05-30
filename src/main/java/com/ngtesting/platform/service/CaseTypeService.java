package com.ngtesting.platform.service;

import java.util.List;

import com.ngtesting.platform.entity.TestCaseType;
import com.ngtesting.platform.vo.CaseTypeVo;

public interface CaseTypeService extends BaseService {
	List<TestCaseType> list(Long orgId);
	
	TestCaseType save(CaseTypeVo vo, Long orgId);
	boolean delete(Long id);

	List<CaseTypeVo> genVos(List<TestCaseType> pos);
	CaseTypeVo genVo(TestCaseType user);

	boolean setDefaultPers(Long orgId, Long orgId2);

	List<CaseTypeVo> listVos(Long orgId);

	boolean changeOrderPers(Long id, String act);

}
