package com.ngtesting.platform.service;

import java.util.List;

import com.ngtesting.platform.entity.SysOrgGroup;
import com.ngtesting.platform.vo.OrgGroupVo;
import com.ngtesting.platform.vo.OrgPrivilegeVo;
import com.ngtesting.platform.vo.Page;

public interface OrgPrivilegeService extends BaseService {

	List<OrgPrivilegeVo> listPrivilegesByOrg(Long orgId, Long orgRoleId);

	boolean saveOrgPrivileges(Long roleId, List<OrgPrivilegeVo> orgPrivileges);

}
