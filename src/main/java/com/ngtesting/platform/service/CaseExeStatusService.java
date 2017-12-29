package com.ngtesting.platform.service;

import com.ngtesting.platform.entity.TestCaseExeStatus;
import com.ngtesting.platform.vo.CaseExeStatusVo;

import java.util.List;

public interface CaseExeStatusService extends BaseService {
	List<TestCaseExeStatus> list(Long orgId);
	
	TestCaseExeStatus save(CaseExeStatusVo vo, Long orgId);
	boolean delete(Long id);

	List<CaseExeStatusVo> genVos(List<TestCaseExeStatus> pos);
	CaseExeStatusVo genVo(TestCaseExeStatus user);

	List<CaseExeStatusVo> listVos(Long orgId);

	boolean changeOrderPers(Long id, String act);

    void createDefaultBasicDataPers(Long id);
}
