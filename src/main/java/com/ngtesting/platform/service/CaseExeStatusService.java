package com.ngtesting.platform.service;

import java.util.List;

import com.ngtesting.platform.entity.SysCaseExeStatus;
import com.ngtesting.platform.entity.SysCustomField;
import com.ngtesting.platform.entity.SysUser;
import com.ngtesting.platform.vo.CaseExeStatusVo;
import com.ngtesting.platform.vo.CustomFieldVo;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.UserVo;

public interface CaseExeStatusService extends BaseService {
	List<SysCaseExeStatus> list(Long orgId);
	
	SysCaseExeStatus save(CaseExeStatusVo vo, Long orgId);
	boolean delete(Long id);

	List<CaseExeStatusVo> genVos(List<SysCaseExeStatus> pos);
	CaseExeStatusVo genVo(SysCaseExeStatus user);

	List<CaseExeStatusVo> listVos(Long orgId);

}
