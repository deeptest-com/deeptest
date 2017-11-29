package com.ngtesting.platform.service;

import com.ngtesting.platform.entity.SysPrivilege;
import com.ngtesting.platform.vo.SysPrivilegeVo;

import java.util.List;
import java.util.Map;

public interface SysPrivilegeService extends BaseService {

	Map<String, Boolean> listByUser(Long userId);

	List<SysPrivilegeVo> genVos(List<SysPrivilege> pos);

	SysPrivilegeVo genVo(SysPrivilege po);
}
