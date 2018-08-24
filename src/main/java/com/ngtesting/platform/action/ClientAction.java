package com.ngtesting.platform.action;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstOrg;
import com.ngtesting.platform.model.TstProjectAccessHistory;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.*;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

import javax.servlet.http.HttpServletRequest;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

@Controller
@RequestMapping(value = Constant.API_PATH_CLIENT + "/client")
public class ClientAction extends BaseAction {
    @Autowired
    private UserService userService;

    @Autowired
    private OrgService orgService;

    @Autowired
    OrgGroupUserRelationService orgGroupUserRelationService;

    @Autowired
    private ProjectService projectService;

    @Autowired
    SysPrivilegeService sysPrivilegeService;
    @Autowired
    OrgPrivilegeService orgRolePrivilegeService;
    @Autowired
    CasePropertyService casePropertyService;
    @Autowired
    ProjectPrivilegeService projectPrivilegeService;
    @Autowired
    PushSettingsService pushSettingsService;

    @PostMapping(value = "getProfile")
    @ResponseBody
    public Map<String, Object> getProfile(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer orgId = user.getDefaultOrgId();
        Integer prjId = user.getDefaultPrjId();
        Integer userId = user.getId();

        Integer orgIdNew = json.getInteger("orgId");
        Integer prjIdNew = json.getInteger("prjId");

        if (orgIdNew != null && orgIdNew.longValue() != orgId.longValue()) { // org不能为空
            userService.setDefaultOrg(user, orgId);
        }
        if (prjIdNew != null && (prjId == null || prjIdNew.longValue() != prjId.longValue())) { // prj可能为空
            projectService.view(prjIdNew, user);
        }

        Map<String, Boolean> sysPrivileges = sysPrivilegeService.listByUser(userId);
        ret.put("sysPrivileges", sysPrivileges);
        Map<String, Boolean> orgPrivileges = orgRolePrivilegeService.listByUser(user.getId(), orgId);
        ret.put("orgPrivileges", orgPrivileges);

        List<TstOrg> orgs = orgService.listByUser(userId);
        ret.put("myOrgs", orgs);

        Map<String,Map<String,String>> casePropertyMap = casePropertyService.getMap(orgId);
        ret.put("casePropertyMap", casePropertyMap);

        List<TstProjectAccessHistory> recentProjects = projectService.listRecentProject(orgId, userId);
        ret.put("recentProjects", recentProjects);

        Map<String, Boolean> prjPrivileges = projectPrivilegeService.listByUser(userId, prjId, orgId);
        ret.put("prjPrivileges", prjPrivileges);

        ret.put("profile", user);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());

        return ret;
    }

    // 用户修改某个字段
    @RequestMapping(value = "saveInfo", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> saveInfo(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        json.put("id", user.getId()); // 强制设置当前用户的属性

        TstUser po = userService.modifyProp(json);
        request.getSession().setAttribute(Constant.HTTP_SESSION_USER_PROFILE, po);

        pushSettingsService.pushUserSettings(po);

        ret.put("data", po);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @PostMapping(value = "setLeftSize")
    @ResponseBody
    public Map<String, Object> setLeftSize(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

        Integer left = json.getInteger("left");
        String prop = json.getString("prop");

        user = userService.setLeftSizePers(user, left, prop);

        ret.put("data", user);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

}
