package com.ngtesting.platform.service;

import com.alibaba.fastjson.JSONArray;
import com.ngtesting.platform.entity.TestUser;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.RelationProjectRoleEntityVo;
import com.ngtesting.platform.vo.UserVo;

import java.util.List;

public interface UserService extends BaseService {
	Page listByPage(Long orgId, String keywords, String disabled, Integer currentPage, Integer itemsPerPage);
	List<RelationProjectRoleEntityVo> getProjectUsers(Long orgId, Long projectId);
    List<TestUser> search(Long orgId, String keywords, JSONArray exceptIds);
	
	TestUser save(UserVo vo, Long orgId);
	boolean disable(Long userId, Long orgId);
	boolean remove(Long userId, Long orgId);
	boolean setSizePers(Long userId, Integer left, Integer right);

	List<UserVo> genVos(List<TestUser> pos);
	UserVo genVo(TestUser user);

}

