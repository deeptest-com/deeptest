package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.*;
import com.ngtesting.platform.service.intf.*;
import com.ngtesting.platform.servlet.PrivPrj;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;

import javax.servlet.http.HttpServletRequest;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

@Controller
@RequestMapping(value = Constant.API_PATH_CLIENT + "/project_member")
public class ProjectMemberAction extends BaseAction {
    @Autowired
    private ProjectService projectService;
    @Autowired
    private ProjectConfigService projectConfigService;

    @Autowired
    private ProjectRoleService projectRoleService;
    @Autowired
    private ProjectRoleEntityRelationService projectRoleEntityRelationService;
    @Autowired
    private HistoryService historyService;

    @Autowired
    private PushSettingsService pushSettingsService;

    @Autowired
    AuthService authService;

    @ResponseBody
    @PostMapping("/getUsers")
    @PrivPrj(perms = {"project-admin"})
    public Map<String, Object> getUsers(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer orgId = user.getDefaultOrgId();
        Integer prjId = user.getDefaultPrjId();

        if (userNotInProject(user.getId(), prjId)) {
            return authFail();
        }

        List<TstProjectRole> projectRoles = projectRoleService.list(orgId, null, null);

        List<TstProjectRoleEntityRelation> entityInRoles = projectRoleEntityRelationService.listByProject(prjId);

        ret.put("projectRoles", projectRoles);
        ret.put("entityInRoles", entityInRoles);

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @PostMapping(value = "saveMembers")
    @ResponseBody
    @PrivPrj(perms = {"project-admin"})
    public Map<String, Object> saveMembers(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer orgId = user.getDefaultOrgId();
        Integer prjId = user.getDefaultPrjId();

        if (userNotInProject(user.getId(), prjId)) {
            return authFail();
        }

        List<TstProjectRoleEntityRelation> entityInRoles = projectRoleEntityRelationService.batchSavePers(json, orgId);

        TstProject project = projectService.get(prjId);
        historyService.create(prjId, user, Constant.MsgType.update.msg,
                TstHistory.TargetType.project_member, prjId, project.getName());

        ret.put("entityInRoles", entityInRoles);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @PostMapping(value = "changeRole")
    @ResponseBody
    @PrivPrj(perms = {"project-admin"})
    public Map<String, Object> changeRole(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer prjId = user.getDefaultPrjId();

        if (userNotInProject(user.getId(), prjId)) {
            return authFail();
        }

        List<TstProjectRoleEntityRelation> entityInRoles = projectRoleEntityRelationService.changeRolePers(json, prjId);

        pushSettingsService.pushPrjSettings(user);

        ret.put("entityInRoles", entityInRoles);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @PostMapping(value = "remove")
    @ResponseBody
    @PrivPrj(perms = {"project-admin"})
    public Map<String, Object> remove(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer prjId = user.getDefaultPrjId();

        String type = json.getString("type");
        Integer entityId = json.getInteger("entityId");

        if (userNotInProject(user.getId(), prjId)) {
            return authFail();
        }

        List<TstProjectRoleEntityRelation> entityInRoles =
                projectRoleEntityRelationService.remove(prjId, type, entityId);

        pushSettingsService.pushPrjSettings(user);

        ret.put("entityInRoles", entityInRoles);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

}
