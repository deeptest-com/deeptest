package com.ngtesting.platform.service;

import java.util.List;

import com.ngtesting.platform.entity.SysOrgGroup;
import com.ngtesting.platform.vo.OrgGroupVo;
import com.ngtesting.platform.vo.OrgPriviledgeVo;
import com.ngtesting.platform.vo.Page;

public interface OrgPriviledgeService extends BaseService {

	List<OrgPriviledgeVo> listPriviledgesByOrg(Long orgId, Long orgRoleId);

	boolean saveOrgPriviledges(Long roleId, List<OrgPriviledgeVo> orgPriviledges);

}
