package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.AuthDao;
import com.ngtesting.platform.model.TstProject;
import com.ngtesting.platform.service.AuthService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class AuthServiceImpl extends BaseServiceImpl implements AuthService {
    @Autowired
    AuthDao authDao;

    @Override
    public Boolean noProjectAndProjectGroupPrivilege(Integer userId, TstProject project) {
        Boolean hasPrjPriv = project.getType().equals(TstProject.ProjectType.project) &&
                !authDao.userNotInProject(userId, project.getId());

        Boolean hasPrjGroupPriv = project.getType().equals(TstProject.ProjectType.group) &&
                !authDao.userNotInOrg(userId, project.getOrgId());

        return !hasPrjPriv && !hasPrjGroupPriv;
    }

}
