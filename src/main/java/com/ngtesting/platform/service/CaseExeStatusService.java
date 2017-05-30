package com.ngtesting.platform.service;

import java.util.List;

import com.ngtesting.platform.entity.TestCaseExeStatus;
import com.ngtesting.platform.entity.TestCustomField;
import com.ngtesting.platform.entity.TestUser;
import com.ngtesting.platform.vo.CaseExeStatusVo;
import com.ngtesting.platform.vo.CustomFieldVo;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.UserVo;

public interface CaseExeStatusService extends BaseService {
	List<TestCaseExeStatus> list(Long orgId);
	
	TestCaseExeStatus save(CaseExeStatusVo vo, Long orgId);
	boolean delete(Long id);

	List<CaseExeStatusVo> genVos(List<TestCaseExeStatus> pos);
	CaseExeStatusVo genVo(TestCaseExeStatus user);

	List<CaseExeStatusVo> listVos(Long orgId);

	boolean changeOrderPers(Long id, String act);

}
