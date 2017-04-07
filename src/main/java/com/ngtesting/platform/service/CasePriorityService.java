package com.ngtesting.platform.service;

import java.util.List;

import com.ngtesting.platform.entity.SysCasePriority;
import com.ngtesting.platform.entity.SysCustomField;
import com.ngtesting.platform.entity.SysUser;
import com.ngtesting.platform.vo.CasePriorityVo;
import com.ngtesting.platform.vo.CustomFieldVo;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.UserVo;

public interface CasePriorityService extends BaseService {
	List<SysCasePriority> list(Long orgId);
	
	SysCasePriority save(CasePriorityVo vo, Long orgId);
	boolean delete(Long id);

	List<CasePriorityVo> genVos(List<SysCasePriority> pos);
	CasePriorityVo genVo(SysCasePriority user);

	boolean setDefaultPers(Long id, Long orgId);

	List<CasePriorityVo> listVos(Long orgId);

}
