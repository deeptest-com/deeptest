package com.ngtesting.platform.service;

import java.util.List;

import com.ngtesting.platform.entity.SysGroupUser;
import com.ngtesting.platform.entity.SysUser;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.UserVo;

public interface UserGroupService extends BaseService {

	SysGroupUser getGroupUser(Long companyId, Long userId, Long groupId);

	List<SysGroupUser> listUserGroups(Long companyId, Long userId, Long groupId);


}
