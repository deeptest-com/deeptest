package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.config.WsConstant;
import com.ngtesting.platform.model.TstOrg;
import com.ngtesting.platform.model.TstProjectAccessHistory;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.inf.*;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

@Service
public class PushSettingsServiceImpl extends BaseServiceImpl implements PushSettingsService {
//    @Autowired
//    private ApplicationScopeBean scopeBean;
//
//    @Autowired
//    SysPrivilegeService sysPrivilegeService;
    @Autowired
    OrgRolePrivilegeService orgRolePrivilegeService;
    @Autowired
    ProjectPrivilegeService projectPrivilegeService;
    @Autowired
    CasePropertyService casePropertyService;

    @Autowired
    OrgService orgService;
    @Autowired
    ProjectService projectService;

    @Override
    public void pushUserSettings(TstUser TstUser) {
        Map<String, Object> ret = new HashMap<>();
        ret.put("code", 1);
        ret.put("type", WsConstant.WS_USER_SETTINGS);

        Integer userId = TstUser.getId();

//        Map<String, Boolean> sysPrivileges = sysPrivilegeService.listByUser(userId);
//        ret.put("sysPrivileges", sysPrivileges);

        ret.put("profile", TstUser);
        sendMsg(TstUser.getId(), ret);
    }

    @Override
    public void pushMyOrgs(TstUser TstUser) {
        Map<String, Object> ret = new HashMap<>();
        ret.put("code", 1);
        ret.put("type", WsConstant.WS_MY_ORGS);

        Integer userId = TstUser.getId();

        List<TstOrg> orgs = orgService.listVo(null, "false", userId);
        ret.put("myOrgs", orgs);

        ret.put("defaultOrgId", TstUser.getDefaultOrgId());

        sendMsg(TstUser.getId(), ret);
    }

    @Override
    public void pushOrgSettings(TstUser TstUser) {
        Map<String, Object> ret = new HashMap<>();
        ret.put("code", 1);
        ret.put("type", WsConstant.WS_ORG_SETTINGS);

        Integer userId = TstUser.getId();
        Integer orgId = TstUser.getDefaultOrgId();

//        TstOrg org = (TstOrg)get(TstOrg.class, orgId);
        Map<String, Boolean> orgPrivileges = orgRolePrivilegeService.listByUser(TstUser.getId(), orgId);
        Map<String,Map<String,String>> casePropertyMap = casePropertyService.getMap(orgId);

//        ret.put("org", org);

        ret.put("defaultOrgId", TstUser.getDefaultOrgId());
        ret.put("defaultOrgName", TstUser.getDefaultOrgName());
        ret.put("defaultPrjId", TstUser.getDefaultPrjId());
        ret.put("defaultPrjName", TstUser.getDefaultPrjName());
        ret.put("orgPrivileges", orgPrivileges);
        ret.put("casePropertyMap", casePropertyMap);

        sendMsg(userId, ret);
    }

    @Override
    public void pushRecentProjects(TstUser TstUser) {
        Map<String, Object> ret = new HashMap<>();
        ret.put("code", 1);
        ret.put("type", WsConstant.WS_RECENT_PROJECTS);

        Integer userId = TstUser.getId();
        Integer orgId = TstUser.getDefaultOrgId();

        List<TstProjectAccessHistory> recentProjects = projectService.listRecentProjectVo(orgId, userId);
        ret.put("recentProjects", recentProjects);

        ret.put("defaultOrgId", orgId);
        ret.put("defaultPrjId", TstUser.getDefaultPrjId());

        sendMsg(TstUser.getId(), ret);
    }

    @Override
    public void pushPrjSettings(TstUser TstUser) {
        Map<String, Object> ret = new HashMap<>();
        ret.put("code", 1);
        ret.put("type", WsConstant.WS_PRJ_SETTINGS);

        Integer userId = TstUser.getId();
        Integer orgId = TstUser.getDefaultOrgId();
        Integer prjId = TstUser.getDefaultPrjId();

        Map<String, Boolean> prjPrivileges = projectPrivilegeService.listByUserPers(userId, prjId, orgId);
        ret.put("prjPrivileges", prjPrivileges);
        ret.put("prjName", TstUser.getDefaultPrjName());

        sendMsg(userId, ret);
    }

    @Override
    public void sendMsg(Integer userId, Map ret) {
//        scopeBean.sendMessageToClient(userId.toString(), new TextMessage(JSON.toJSONString(ret)));
    }

}

