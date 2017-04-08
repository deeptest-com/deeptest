package com.ngtesting.platform.service;

import java.util.List;

import com.ngtesting.platform.entity.SysCustomField;
import com.ngtesting.platform.entity.SysUser;
import com.ngtesting.platform.vo.CustomFieldVo;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.UserVo;

public interface CustomFieldService extends BaseService {
	List<SysCustomField> list(Long orgId);
	
	SysCustomField save(CustomFieldVo vo, Long orgId);
	boolean delete(Long id);

	List<CustomFieldVo> genVos(List<SysCustomField> pos);
	CustomFieldVo genVo(SysCustomField po);

	List<CustomFieldVo> listVos(Long orgId);

	List<String> listApplyTo();

	List<String> listType();

	List<String> listFormat();

	boolean changeOrderPers(Long id, String act);

}
