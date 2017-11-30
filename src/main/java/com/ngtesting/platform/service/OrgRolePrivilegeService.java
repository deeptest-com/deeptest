package com.ngtesting.platform.service;

import com.ngtesting.platform.vo.OrgPrivilegeVo;

import java.util.List;
import java.util.Map;

public interface OrgRolePrivilegeService extends BaseService {

	List<OrgPrivilegeVo> listPrivilegesByOrgRole(Long orgId, Long orgRoleId);

	boolean saveOrgRolePrivileges(Long roleId, List<OrgPrivilegeVo> orgPrivileges);

	Map<String, Boolean> listByUser(Long userId, Long orgId);

}
