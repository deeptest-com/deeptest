package com.ngtesting.platform.service;

import java.util.List;

import com.ngtesting.platform.entity.SysOrg;
import com.ngtesting.platform.entity.SysUser;
import com.ngtesting.platform.vo.OrgVo;
import com.ngtesting.platform.vo.TestProjectAccessHistoryVo;
import com.ngtesting.platform.vo.TestProjectVo;
import com.ngtesting.platform.vo.UserVo;



public interface OrgService extends BaseService {

	List<SysOrg> list(String keywords, String disabled, Long userId);
	List<OrgVo> listVo(String keywords, String disabled, Long id);

	SysOrg getDetail(Long id);

	Boolean disable(Long id);
	Boolean delete(Long id);

	List<OrgVo> genVos(List<SysOrg> pos, Long userId);

	OrgVo genVo(SysOrg po);

	SysOrg save(OrgVo vo, Long userId);

	List<TestProjectAccessHistoryVo> setDefaultPers(Long orgId, UserVo user);

	

}
