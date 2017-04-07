package com.ngtesting.platform.service;

import java.util.List;

import com.ngtesting.platform.entity.SysCaseType;
import com.ngtesting.platform.vo.CaseTypeVo;

public interface CaseTypeService extends BaseService {
	List<SysCaseType> list(Long orgId);
	
	SysCaseType save(CaseTypeVo vo, Long orgId);
	boolean delete(Long id);

	List<CaseTypeVo> genVos(List<SysCaseType> pos);
	CaseTypeVo genVo(SysCaseType user);

	boolean setDefaultPers(Long orgId, Long orgId2);

	List<CaseTypeVo> listVos(Long orgId);

}
