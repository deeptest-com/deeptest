package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.ProjectRoleDao;
import com.ngtesting.platform.model.TstProjectRole;
import com.ngtesting.platform.service.ProjectPrivilegeService;
import com.ngtesting.platform.service.ProjectRoleService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.UUID;

@Service
public class ProjectRoleServiceImpl extends BaseServiceImpl implements ProjectRoleService {
	@Autowired
	private ProjectRoleDao projectRoleDao;

    @Autowired
    private ProjectPrivilegeService projectPrivilegeService;

	@Override
	public List list(Integer orgId, String keywords, Boolean disabled) {
		List<TstProjectRole> ls = projectRoleDao.query(orgId, keywords, disabled);

		return ls;
	}

	@Override
	public TstProjectRole get(Integer roleId, Integer orgId) {
		return projectRoleDao.get(roleId, orgId);
	}

	@Override
	public TstProjectRole save(TstProjectRole vo, Integer orgId) {
        vo.setOrgId(orgId);

		if (vo.getId() == null) {
			vo.setCode(UUID.randomUUID().toString());
			projectRoleDao.save(vo);
		} else {
			Integer count = projectRoleDao.update(vo);
			if (count == 0) {
                return null;
            }
		}

		return vo;
	}

	@Override
	public boolean delete(Integer id, Integer orgId) {
        Integer count = projectRoleDao.delete(id, orgId);

        return count > 0;
	}

}
