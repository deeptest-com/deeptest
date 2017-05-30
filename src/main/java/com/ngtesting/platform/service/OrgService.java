package com.ngtesting.platform.service;

import java.util.List;

import com.ngtesting.platform.entity.TestOrg;
import com.ngtesting.platform.entity.TestUser;
import com.ngtesting.platform.vo.OrgVo;
import com.ngtesting.platform.vo.TestProjectAccessHistoryVo;
import com.ngtesting.platform.vo.TestProjectVo;
import com.ngtesting.platform.vo.UserVo;



public interface OrgService extends BaseService {

	List<TestOrg> list(String keywords, String disabled, Long userId);
	List<OrgVo> listVo(String keywords, String disabled, Long id);

	TestOrg getDetail(Long id);

	Boolean disable(Long id);
	Boolean delete(Long id);

	List<OrgVo> genVos(List<TestOrg> pos, Long userId);

	OrgVo genVo(TestOrg po);

	TestOrg save(OrgVo vo, Long userId);

	List<TestProjectAccessHistoryVo> setDefaultPers(Long orgId, UserVo user);

	

}
