package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstProjectRole;
import com.ngtesting.platform.model.TstProjectRoleEntityRelation;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.AuthService;
import com.ngtesting.platform.service.intf.ProjectRoleEntityRelationService;
import com.ngtesting.platform.service.intf.ProjectRoleService;
import com.ngtesting.platform.service.intf.PushSettingsService;
import com.ngtesting.platform.servlet.PrivPrj;
import org.apache.shiro.SecurityUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import javax.servlet.http.HttpServletRequest;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

@RestController
@RequestMapping(value = Constant.API_PATH_CLIENT + "/project_member")
public class ProjectMemberAction extends BaseAction {

    @Autowired
    private ProjectRoleService projectRoleService;
    @Autowired
    private ProjectRoleEntityRelationService projectRoleEntityRelationService;

    @Autowired
    private PushSettingsService pushSettingsService;

    @Autowired
    AuthService authService;

    @PostMapping("/getUsers")
    @PrivPrj(perms = {"project:*"})
    public Map<String, Object> getUsers(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();
        Integer orgId = user.getDefaultOrgId();
        Integer prjId = user.getDefaultPrjId();

        List<TstProjectRole> projectRoles = projectRoleService.list(orgId, null, null);

        List<TstProjectRoleEntityRelation> entityInRoles = projectRoleEntityRelationService.listByProject(prjId);

        ret.put("projectRoles", projectRoles);
        ret.put("entityInRoles", entityInRoles);

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @PostMapping(value = "saveMembers")
    @PrivPrj(perms = {"project:*"})
    public Map<String, Object> saveMembers(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();

        List<TstProjectRoleEntityRelation> entityInRoles = projectRoleEntityRelationService.batchSavePers(json, user);

        ret.put("entityInRoles", entityInRoles);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @PostMapping(value = "changeRole")
    @PrivPrj(perms = {"project:*"})
    public Map<String, Object> changeRole(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();

        List<TstProjectRoleEntityRelation> entityInRoles = projectRoleEntityRelationService.changeRolePers(json, user);

        pushSettingsService.pushPrjSettings(user);

        ret.put("entityInRoles", entityInRoles);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @PostMapping(value = "remove")
    @PrivPrj(perms = {"project:*"})
    public Map<String, Object> remove(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();

        String type = json.getString("type");
        Integer entityId = json.getInteger("entityId");

        List<TstProjectRoleEntityRelation> entityInRoles =
                projectRoleEntityRelationService.remove(type, entityId, user);

        pushSettingsService.pushPrjSettings(user);

        ret.put("entityInRoles", entityInRoles);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

}
