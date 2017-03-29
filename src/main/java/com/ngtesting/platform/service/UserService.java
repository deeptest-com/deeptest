package com.ngtesting.platform.service;

import java.util.List;
import java.util.Map;

import com.ngtesting.platform.entity.SysUser;
import com.ngtesting.platform.entity.SysVerifyCode;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.UserVo;

public interface UserService extends BaseService {

	Page listByPage(Long companyId, String keywords, String disabled, Integer currentPage, Integer itemsPerPage);
	
	SysUser save(UserVo vo, Long companyId);
	boolean remove(Long id);
	boolean disablePers(Long id);

	List<UserVo> genVos(List<SysUser> pos);
	UserVo genVo(SysUser user);

}
