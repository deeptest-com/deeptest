package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.ngtesting.platform.bean.ApplicationScopeBean;
import com.ngtesting.platform.config.WsConstant;
import com.ngtesting.platform.service.*;
import com.ngtesting.platform.vo.OrgVo;
import com.ngtesting.platform.vo.TestProjectAccessHistoryVo;
import com.ngtesting.platform.vo.UserVo;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.web.socket.TextMessage;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

@Service
public class PushSettingsServiceImpl extends BaseServiceImpl implements PushSettingsService {
    @Autowired
    private ApplicationScopeBean scopeBean;

    @Autowired
    SysPrivilegeService sysPrivilegeService;
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
    public void pushUserSettings(UserVo userVo) {
        Map<String, Object> ret = new HashMap<>();
        ret.put("code", 1);
        ret.put("type", WsConstant.WS_USER_SETTINGS);

        Long userId = userVo.getId();

        Map<String, Boolean> sysPrivileges = sysPrivilegeService.listByUser(userId);
        ret.put("sysPrivileges", sysPrivileges);

        ret.put("profile", userVo);
        sendMsg(userVo.getId(), ret);
    }

    @Override
    public void pushMyOrgs(UserVo userVo) {
        Map<String, Object> ret = new HashMap<>();
        ret.put("code", 1);
        ret.put("type", WsConstant.WS_MY_ORGS);

        Long userId = userVo.getId();

        List<OrgVo> orgs = orgService.listVo(null, "false", userId);
        ret.put("myOrgs", orgs);

        ret.put("defaultOrgId", userVo.getDefaultOrgId());

        sendMsg(userVo.getId(), ret);
    }

    @Override
    public void pushOrgSettings(UserVo userVo) {
        Map<String, Object> ret = new HashMap<>();
        ret.put("code", 1);
        ret.put("type", WsConstant.WS_ORG_SETTINGS);

        Long userId = userVo.getId();
        Long orgId = userVo.getDefaultOrgId();

        Map<String, Boolean> orgPrivileges = orgRolePrivilegeService.listByUser(userVo.getId(), orgId);
        Map<String,Map<String,String>> casePropertyMap = casePropertyService.getMap(orgId);

        ret.put("defaultPrjId", userVo.getDefaultPrjId());
        ret.put("orgPrivileges", orgPrivileges);
        ret.put("casePropertyMap", casePropertyMap);

        sendMsg(userId, ret);
    }

    @Override
    public void pushRecentProjects(UserVo userVo) {
        Map<String, Object> ret = new HashMap<>();
        ret.put("code", 1);
        ret.put("type", WsConstant.WS_RECENT_PROJECTS);

        Long userId = userVo.getId();
        Long orgId = userVo.getDefaultOrgId();

        List<TestProjectAccessHistoryVo> recentProjects = projectService.listRecentProjectVo(orgId, userId);
        ret.put("recentProjects", recentProjects);

        ret.put("defaultOrgId", orgId);
        ret.put("defaultPrjId", userVo.getDefaultPrjId());

        sendMsg(userVo.getId(), ret);
    }

    @Override
    public void pushPrjSettings(UserVo userVo) {
        Map<String, Object> ret = new HashMap<>();
        ret.put("code", 1);
        ret.put("type", WsConstant.WS_PRJ_SETTINGS);

        Long userId = userVo.getId();
        Long orgId = userVo.getDefaultOrgId();
        Long prjId = userVo.getDefaultPrjId();

        Map<String, Boolean> prjPrivileges = projectPrivilegeService.listByUserPers(userId, prjId, orgId);
        ret.put("prjPrivileges", prjPrivileges);
        ret.put("prjName", userVo.getDefaultPrjName());

        sendMsg(userId, ret);
    }

    @Override
    public void sendMsg(Long userId, Map ret) {
        scopeBean.sendMessageToClient(userId.toString(), new TextMessage(JSON.toJSONString(ret)));
    }

}

