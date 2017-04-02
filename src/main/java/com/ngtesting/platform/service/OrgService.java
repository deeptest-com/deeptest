package com.ngtesting.platform.service;

import java.util.List;

import com.ngtesting.platform.entity.SysOrg;
import com.ngtesting.platform.vo.OrgVo;



public interface OrgService extends BaseService {

	List<SysOrg> list(String keywords, Boolean disabled, Long userId);

	SysOrg getDetail(Long id);

	Boolean disable(Long id);
	Boolean delete(Long id);

	List<OrgVo> genVos(List<SysOrg> pos);

	OrgVo genVo(SysOrg po);

	SysOrg save(OrgVo vo, Long userId);

}
