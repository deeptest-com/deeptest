package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.ngtesting.platform.config.WsConstant;
import com.ngtesting.platform.model.IsuQuery;
import com.ngtesting.platform.model.TstOrg;
import com.ngtesting.platform.model.TstProjectAccessHistory;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.*;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.messaging.simp.SimpMessagingTemplate;
import org.springframework.stereotype.Service;
import org.springframework.web.socket.TextMessage;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

@Service
public class PushSettingsServiceImpl extends BaseServiceImpl implements PushSettingsService {
    @Autowired
    private SimpMessagingTemplate simpMessagingTemplate;

    @Autowired
    SysPrivilegeService sysPrivilegeService;

    @Autowired
    OrgPrivilegeService orgRolePrivilegeService;
    @Autowired
    ProjectPrivilegeService projectPrivilegeService;
    @Autowired
    CasePropertyService casePropertyService;

    @Autowired
    private IssueQueryService issueQueryService;

    @Autowired
    OrgService orgService;
    @Autowired
    ProjectService projectService;

    @Autowired
    CustomFieldService customFieldService;
    @Autowired
    IssueDynamicFormService dynamicFormService;
    @Autowired
    IssueWorkflowTransitionService issueWorkflowTransitionService;

    @Override
    public void pushUserSettings(TstUser user) {
        Map<String, Object> ret = new HashMap<>();
        ret.put("code", 1);
        ret.put("type", WsConstant.WS_USER_SETTINGS);

        Map<String, Boolean> sysPrivileges = sysPrivilegeService.listByUser(user.getId());
        ret.put("sysPrivileges", sysPrivileges);

        ret.put("profile", user);
        sendMsg(user, ret);
    }

    @Override
    public void pushOrgSettings(TstUser user) {
        Map<String, Object> ret = new HashMap<>();
        ret.put("code", 1);
        ret.put("type", WsConstant.WS_ORG_SETTINGS);

        Integer userId = user.getId();
        Integer orgId = user.getDefaultOrgId();

        Map<String, Boolean> orgPrivileges = orgRolePrivilegeService.listByUser(userId, orgId);

        ret.put("orgPrivileges", orgPrivileges);

        ret.put("defaultOrgId", user.getDefaultOrgId());
        ret.put("defaultOrgName", user.getDefaultOrgName());

        ret.put("defaultPrjId", user.getDefaultPrjId());
        ret.put("defaultPrjName", user.getDefaultPrjName());

        sendMsg(user, ret);
    }

    @Override
    public void pushPrjSettings(TstUser user) {
        Map<String, Object> ret = new HashMap<>();
        ret.put("code", 1);
        ret.put("type", WsConstant.WS_PRJ_SETTINGS);

        Integer userId = user.getId();
        Integer orgId = user.getDefaultOrgId();
        Integer prjId = user.getDefaultPrjId();

        ret.put("prjId", user.getDefaultPrjId());
        ret.put("prjName", user.getDefaultPrjName());

        // 权限
        Map<String, Boolean> prjPrivileges = projectPrivilegeService.listByUser(userId, prjId, orgId);
        ret.put("prjPrivileges", prjPrivileges);

        // 用例
        Map<String, Object> map = customFieldService.fetchProjectFieldForCase(orgId, prjId);
        ret.put("caseCustomFields", map.get("fields"));
        ret.put("casePropMap", map.get("props"));
        Map<String,Map<String,String>> casePropValMap = casePropertyService.getMap(orgId);
        ret.put("casePropValMap", casePropValMap);

        // 缺陷
        Map issuePropMap = dynamicFormService.genIssuePropMap(orgId, prjId);
        ret.put("issuePropMap", issuePropMap);
        Map<String, Object> issuePropValMap = dynamicFormService.genIssueBuldInPropValMap(orgId, prjId);
        ret.put("issuePropValMap", issuePropValMap);

        Map issueTransMap = issueWorkflowTransitionService.getStatusTrainsMap(prjId, userId);
        ret.put("issueTransMap", issueTransMap);

        sendMsg(user, ret);
    }

    @Override
    public void pushRecentQueries(TstUser user) {
        Map<String, Object> ret = new HashMap<>();
        ret.put("code", 1);
        ret.put("type", WsConstant.WS_RECENT_QUERIES);

        List<IsuQuery> pos = issueQueryService.listRecentQuery(user.getDefaultOrgId(), user.getId());
        ret.put("recentQueries", pos);

        sendMsg(user, ret);
    }

    @Override
    public void pushMyOrgs(TstUser user) {
        Map<String, Object> ret = new HashMap<>();
        ret.put("code", 1);
        ret.put("type", WsConstant.WS_MY_ORGS);

        Integer userId = user.getId();

        List<TstOrg> orgs = orgService.listByUser(userId);
        ret.put("myOrgs", orgs);

        ret.put("defaultOrgId", user.getDefaultOrgId());

        sendMsg(user, ret);
    }

    @Override
    public void pushRecentProjects(TstUser user) {
        Map<String, Object> ret = new HashMap<>();
        ret.put("code", 1);
        ret.put("type", WsConstant.WS_RECENT_PROJECTS);

        Integer userId = user.getId();
        Integer orgId = user.getDefaultOrgId();

        List<TstProjectAccessHistory> recentProjects = projectService.listRecentProject(orgId, userId);
        ret.put("recentProjects", recentProjects);

        ret.put("defaultOrgId", orgId);
        ret.put("defaultPrjId", user.getDefaultPrjId());

        sendMsg(user, ret);
    }

    @Override
    public void sendMsg(TstUser user, Map ret) {
        simpMessagingTemplate.convertAndSendToUser(user.getToken(), "/notification",
                new TextMessage(JSON.toJSONString(ret)));
    }

}

