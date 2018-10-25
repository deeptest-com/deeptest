package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSONObject;
import com.github.pagehelper.PageHelper;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.dao.*;
import com.ngtesting.platform.model.*;
import com.ngtesting.platform.service.*;
import com.ngtesting.platform.utils.PasswordEncoder;
import com.ngtesting.platform.utils.StringUtil;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

@Service(value = "userService")
public class UserServiceImpl implements UserService {
    @Autowired
    private PropService propService;
    @Autowired
    private MailService mailService;

    @Autowired
    private OrgDao orgDao;
    @Autowired
    private OrgUserRelationDao orgUserRelationDao;
    @Autowired
    private ProjectRoleEntityRelationDao projectRoleEntityRelationDao;

    @Autowired
    private ProjectDao projectDao;
    @Autowired
    private ProjectRoleDao projectRoleDao;
    @Autowired
    private UserDao userDao;

    @Autowired
    private AccountService accountService;
    @Autowired
    private AccountVerifyCodeService accountVerifyCodeService;
    @Autowired
    private ProjectService projectService;
    @Autowired
    OrgGroupUserRelationService orgGroupUserRelationService;
    @Autowired
    OrgPrivilegeService orgPrivilegeService;

    @Override
    public List<TstUser> list(Integer orgId, String keywords, Boolean disabled, int pageNum, int pageSize) {
        List<TstUser> users = userDao.query(orgId, keywords, disabled);

        return users;
    }

    @Override
    public List<TstUser> listAllOrgUsers(Integer orgId) {
        List<TstUser> ls = userDao.query(orgId, null, false);

        return ls;
    }

    @Override
    public List<TstUser> getProjectUsers(Integer orgId, Integer projectId) {
        List<TstUser> ls = userDao.getProjectUsers(projectId, null);

        return ls;
    }

    @Override
    public Boolean removeFromOrg(Integer userId, Integer orgId) {
        userDao.removeFromOrg(userId, orgId);
        return true;
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
    public TstUser getByPhone(String phone) {
        TstUser user = userDao.getByPhone(phone);
        return user;
    }

    @Override
    public TstUser getByEmail(String email) {
        TstUser user = userDao.getByEmail(email);
        return user;
    }

    @Override
    @Transactional
    public TstUser invite(TstUser user, TstUser vo, List<TstOrgGroupUserRelation> relations) {
        Integer orgId = user.getDefaultOrgId();
        Integer prjId = user.getDefaultPrjId();
        String orgName = user.getDefaultOrgName();
        String prjName = user.getDefaultPrjName();

        TstUser existUser  = getByEmail(vo.getEmail());
        boolean isNew;
        if (existUser != null) {
            isNew = false;
            vo = existUser;
        } else {
            isNew = true;
            vo.setDefaultOrgId(orgId);
            vo.setAvatar("upload/sample/user/avatar.png");

            vo.setDefaultOrgId(orgId);
            vo.setDefaultPrjId(prjId);
            vo.setDefaultOrgName(orgName);
            vo.setDefaultPrjName(prjName);

            String salt = PasswordEncoder.genSalt();
            PasswordEncoder passwordEncoder = new  PasswordEncoder(salt);

            user.setTemp(salt);
            user.setPassword(passwordEncoder.encodePassword(StringUtil.RandomString(6)));

            userDao.save(vo);

            userDao.saveSettings(vo);
        }

        if (isNew || orgUserRelationDao.userInOrg(vo.getId(), orgId) == 0) { // 不在组织里
            orgUserRelationDao.addUserToOrg(vo.getId(), orgId);

            Integer projectRoleId = projectRoleDao.getRoleByCode(orgId, "test_designer").getId();
            projectRoleEntityRelationDao.addRole(orgId, prjId, projectRoleId, vo.getId(), "user");

            projectService.changeDefaultPrj(vo, prjId);

            orgGroupUserRelationService.saveRelationsForUser(orgId, vo.getId(), relations);

            // 发送邮件
            String sys = propService.getSysName();
            Map<String, String> map = new HashMap<String, String>();
            map.put("user", user.getNickname() + "(" + user.getEmail() + ")");
            map.put("name", vo.getNickname());
            map.put("sys", sys);

            String url;
            if (isNew) {
                String verifyCode = accountVerifyCodeService.genVerifyCode(vo.getId());

                url = propService.getUrlResetPassword();
                if (!url.startsWith("http")) {
                    url = Constant.WEB_ROOT + url;
                }
                url += "/" + verifyCode;
            } else {
                url = propService.getUrlLogin();
                if (!url.startsWith("http")) {
                    url = Constant.WEB_ROOT + url;
                }
            }
            map.put("url", url);
            mailService.sendTemplateMail("来自[" + sys + "]的邀请", "invite-user.ftl",
                    vo.getEmail(), map);
            return vo;
        } else {
            return null;
        }
    }

    @Override
    @Transactional
    public TstUser update(TstUser vo) {
        userDao.update(vo);

        TstUser user = userDao.get(vo.getId());
        return user;
    }

    @Override
    @Transactional
    public TstUser modifyProp(JSONObject json) {
        Integer id = json.getInteger("id");
        String prop = json.getString("prop");
        String value = json.getString("value");

        userDao.modifyProp(id, prop, value);

        TstUser user = userDao.get(id);
        return user;
    }

    @Override
    public List<TstUser> search(Integer orgId, String keywords, String exceptIds) {
        PageHelper.startPage(0, 20);
        List<TstUser> users = userDao.search(orgId, keywords, exceptIds);

        return users;
    }

    @Override
    public TstUser setIssueView(TstUser user, String issueView) {
        user.setIssueView(issueView);
        userDao.setIssueView(user);
        return user;
    }

    @Override
    public TstUser setLeftSizePers(TstUser user, Integer left, String prop) {
        if ("design".equals(prop)) {
            user.setLeftSizeDesign(left);
        } else if ("exe".equals(prop)) {
            user.setLeftSizeExe(left);
        } else if ("issue".equals(prop)) {
            user.setLeftSizeIssue(left);
        }

        userDao.setLeftSize(user);
        return user;
    }

    @Override
    public void saveIssueColumns(String columnsStr, TstUser user) {
        user.setIssueColumns(columnsStr);
        userDao.saveIssueColumns(columnsStr, user.getId());
    }

    @Override
    public void saveIssueFields(String fieldStr, TstUser user) {
        user.setIssueFileds(fieldStr);
        userDao.saveIssueFields(fieldStr, user.getId());
    }

}
