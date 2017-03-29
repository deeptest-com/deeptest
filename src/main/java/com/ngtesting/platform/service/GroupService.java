package com.ngtesting.platform.service;

import java.util.List;
import java.util.Map;

import com.ngtesting.platform.entity.SysGroup;
import com.ngtesting.platform.entity.SysUser;
import com.ngtesting.platform.entity.SysVerifyCode;
import com.ngtesting.platform.vo.GroupVo;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.UserVo;

public interface GroupService extends BaseService {

	Page listByPage(Long companyId, String keywords, String disabled, Integer currentPage, Integer itemsPerPage);
	
	SysGroup save(GroupVo vo, Long companyId);
	boolean delete(Long id);
	boolean disable(Long id);

	List<GroupVo> genVos(List<SysGroup> pos);
	GroupVo genVo(SysGroup user);

}
