package com.ngtesting.platform.service;

import java.util.List;

import com.ngtesting.platform.entity.SysOrgGroup;
import com.ngtesting.platform.vo.OrgGroupVo;
import com.ngtesting.platform.vo.OrgPriviledgeVo;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.ProjectPriviledgeVo;

public interface ProjectPriviledgeService extends BaseService {

	List<ProjectPriviledgeVo> listPriviledgesByOrg(Long orgId, Long projectRoleId);

	boolean saveProjectPriviledges(Long roleId, List<ProjectPriviledgeVo> projectPriviledges);

}
