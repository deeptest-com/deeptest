package com.ngtesting.platform.service.impl;

import com.github.pagehelper.PageHelper;
import com.github.pagehelper.PageInfo;
import com.ngtesting.platform.dao.OrgDao;
import com.ngtesting.platform.dao.ProjectDao;
import com.ngtesting.platform.dao.UserDao;
import com.ngtesting.platform.model.TstOrg;
import com.ngtesting.platform.model.TstProject;
import com.ngtesting.platform.model.TstProjectAccessHistory;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.ProjectService;
import com.ngtesting.platform.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service(value = "userService")
public class UserServiceImpl implements UserService {

    @Autowired
    private OrgDao orgDao;
    @Autowired
    private ProjectDao projectDao;
    @Autowired
    private UserDao userDao;

    @Autowired
    private ProjectService projectService;

    @Override
    public PageInfo<TstUser> query(int pageNum, int pageSize) {
        PageHelper.startPage(pageNum, pageSize);
        List<TstUser> userDomains = userDao.query();
        PageInfo result = new PageInfo(userDomains);
        return result;
    }

    @Override
    public TstUser get(Integer id) {
        TstUser user = userDao.get(id);
        return user;
    }

    @Override
    public TstUser getByToken(String token) {
        TstUser user = userDao.getByToken(token);
        return user;
    }

    @Override
    public void update(TstUser record) {
        userDao.update(record);
    }

    @Override
    public void setDefaultOrg(TstUser user, Integer orgId) {
        TstOrg org = orgDao.get(orgId);
        userDao.setDefaultOrg(user.getId(), orgId, org.getName());
        user.setDefaultOrgId(orgId);
        user.setDefaultOrgName(org.getName());

        List<TstProjectAccessHistory> recentProjects = projectService.listRecentProject(orgId, user.getId());
        if (recentProjects.size() > 0) {
            TstProjectAccessHistory his = recentProjects.get(0);
            setDefaultPrj(user, his.getPrjId());

        } else {
            setDefaultPrj(user, null);
        }
    }

    @Override
    public void setDefaultPrj(TstUser user, Integer prjId) {
        if (prjId != null) {
            TstProject prj = projectDao.get(prjId);
            userDao.setDefaultPrj(user.getId(), prjId, prj.getName());

            user.setDefaultPrjId(prjId);
            user.setDefaultPrjName(prj.getName());
        } else {
            userDao.setDefaultPrj(user.getId(), null, null);

            user.setDefaultPrjId(null);
            user.setDefaultPrjName(null);
        }
    }

    @Override
    public List<TstUser> search(Integer orgId, String keywords, String exceptIds) {
        PageHelper.startPage(0, 20);
        List<TstUser> users = userDao.search(orgId, keywords, exceptIds);

        return users;
    }

}
