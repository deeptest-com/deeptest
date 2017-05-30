package com.ngtesting.platform.service;

import java.util.List;

import com.ngtesting.platform.entity.TestCasePriority;
import com.ngtesting.platform.entity.TestCustomField;
import com.ngtesting.platform.entity.TestUser;
import com.ngtesting.platform.vo.CasePriorityVo;
import com.ngtesting.platform.vo.CustomFieldVo;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.UserVo;

public interface CasePriorityService extends BaseService {
	List<TestCasePriority> list(Long orgId);
	
	TestCasePriority save(CasePriorityVo vo, Long orgId);
	boolean delete(Long id);

	List<CasePriorityVo> genVos(List<TestCasePriority> pos);
	CasePriorityVo genVo(TestCasePriority user);

	boolean setDefaultPers(Long id, Long orgId);

	List<CasePriorityVo> listVos(Long orgId);

	boolean changeOrderPers(Long id, String act);

}
