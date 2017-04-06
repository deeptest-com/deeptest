package com.ngtesting.platform.service;

import java.util.List;

import com.ngtesting.platform.entity.SysCustomField;
import com.ngtesting.platform.entity.SysUser;
import com.ngtesting.platform.vo.CustomFieldVo;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.UserVo;

public interface CustomFieldService extends BaseService {
	Page listByPage(Long orgId, String keywords, String disabled, Integer currentPage, Integer itemsPerPage);
	
	SysCustomField save(CustomFieldVo vo, Long orgId);
	boolean delete(Long id);

	List<CustomFieldVo> genVos(List<SysCustomField> pos);
	CustomFieldVo genVo(SysCustomField user);

}
