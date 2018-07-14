package com.ngtesting.platform.service.impl;

import com.github.pagehelper.PageHelper;
import com.github.pagehelper.PageInfo;
import com.ngtesting.platform.dao.UserDao;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service(value = "userService")
public class UserServiceImpl implements UserService {

    @Autowired
    private UserDao userDao;

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
    public void setDefaultOrg(Integer userId, Integer orgId) {
        userDao.setDefaultOrg(userId, orgId);

//		TestUser user = (TestUser) get(TestUser.class, TstUser.getId());
//
//		user.setDefaultOrgId(orgId);
//
//		List<TstProjectAccessHistory> recentProjects = projectService.listRecentProject(orgId, TstUser.getId());
//        user.setDefaultPrjId(recentProjects.size()>0?recentProjects.get(0).getProjectId(): null);
//        saveOrUpdate(user);
//
//		TstUser.setDefaultOrgId(user.getDefaultOrgId());
//		if (user.getDefaultOrgId()!=null) {
//			TstOrg org = (TstOrg)get(TstOrg.class, user.getDefaultOrgId());
//			TstUser.setDefaultOrgName(org.getName());
//		}
//
//        TstUser.setDefaultPrjId(recentProjects.size()>0?recentProjects.get(0).getProjectId(): null);
//		TstUser.setDefaultPrjName(recentProjects.size()>0?recentProjects.get(0).getProjectName(): "");
    }

    @Override
    public void setDefaultPrj(Integer userId, Integer prjId, String prjName) {
        userDao.setDefaultPrj(userId, prjId, prjName);
    }

}
