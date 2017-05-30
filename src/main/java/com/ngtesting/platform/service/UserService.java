package com.ngtesting.platform.service;

import java.util.List;

import com.ngtesting.platform.entity.TestUser;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.UserVo;

public interface UserService extends BaseService {
	Page listByPage(Long orgId, String keywords, String disabled, Integer currentPage, Integer itemsPerPage);
	
	TestUser save(UserVo vo, Long orgId);
	boolean disable(Long userId, Long orgId);
	boolean remove(Long userId, Long orgId);

	List<UserVo> genVos(List<TestUser> pos);
	UserVo genVo(TestUser user);

}
