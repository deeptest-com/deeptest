package com.ngtesting.platform.action;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.*;
import com.ngtesting.platform.service.*;
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
@RequestMapping(value = Constant.API_PATH_CLIENT + "/project")
public class ProjectAction {
    @Autowired
    private ProjectService projectService;

    @Autowired
    private TestPlanService planService;
    @Autowired
    private HistoryService historyService;

    @Autowired
    private ProjectRoleService projectRoleService;
    @Autowired
    private ProjectRoleEntityRelationService projectRoleEntityRelationService;

    @Autowired
    private PushSettingsService pushSettingsService;

    @ResponseBody
    @PostMapping("/list")
    public Object list(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
        Integer orgId = user.getDefaultOrgId();

        String keywords = json.getString("keywords");
        Boolean disabled = json.getBoolean("disabled");

        List<TstProject> vos = projectService.list(orgId, user.getId(), keywords, disabled);

        ret.put("data", vos);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());

        return ret;
    }

    @ResponseBody
    @PostMapping("/getInfo")
    public Map<String, Object> getInfo(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
        Integer orgId = user.getDefaultOrgId();

        Integer projectId = json.getInteger("id");

        if (projectId != null) {
            TstProject project = projectService.get(projectId);
            TstProject vo = projectService.genVo(project, null);

            if (TstProject.ProjectType.group.equals(project.getType())) {
                vo.setLastestProjectGroup(projectService.isLastestProjectGroup(orgId, projectId));
            }

            ret.put("data", vo);
        }
        List<TstProject> groups = projectService.listProjectGroups(orgId);
        ret.put("groups", groups);

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @ResponseBody
    @PostMapping("/view")
    public Map<String, Object> view(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
        Integer id = json.getInteger("id");

        TstProject po = projectService.get(id);
        TstProject vo = projectService.genVo(po, null);

        List<TstPlan> planPos = planService.listByProject(id, vo.getType());
        planService.genVos(planPos);

        List<TstHistory> historyPos = historyService.listByProject(id, vo.getType());
        Map<String, List<TstHistory>> historyVos = historyService.genVosByDate(historyPos);

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        ret.put("project", vo);
        ret.put("plans", planPos);
        ret.put("histories", historyVos);

        return ret;
    }

    @ResponseBody
    @PostMapping("/change")
    public Map<String, Object> change(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
        Integer id = json.getInteger("id");

        TstProject vo = projectService.viewPers(id, user);

        if (vo.getType().equals(TstProject.ProjectType.project)) {
            pushSettingsService.pushRecentProjects(user);
            pushSettingsService.pushPrjSettings(user);
        }

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        ret.put("data", vo);

        return ret;
    }

    @ResponseBody
    @PostMapping("/save")
    public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
        Integer userId = user.getId();
        Integer orgId = user.getDefaultOrgId();

        TstProject vo = json.getObject("model", TstProject.class);

        TstProject po = projectService.save(vo, orgId, user);
        if (TstProject.ProjectType.project.equals(po.getType())) {
            projectService.updateNameInHisotyPers(po.getId(), userId);
        }

        pushSettingsService.pushRecentProjects(user);
        pushSettingsService.pushPrjSettings(user);

        ret.put("data", vo);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @PostMapping(value = "delete")
    @ResponseBody
    public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
        Integer id = json.getInteger("id");

        projectService.delete(id, user.getId());

        pushSettingsService.pushRecentProjects(user);
        pushSettingsService.pushPrjSettings(user);

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @ResponseBody
    @PostMapping("/getUsers")
    public Map<String, Object> getUsers(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
        Integer orgId = user.getDefaultOrgId();

        Integer projectId = json.getInteger("id");

        List<TstProjectRole> projectRoles = projectRoleService.list(orgId, null, null);

        List<TstProjectRoleEntityRelation> entityInRoles = projectRoleEntityRelationService.listByProject(projectId);

        ret.put("projectRoles", projectRoles);
        ret.put("entityInRoles", entityInRoles);

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @PostMapping(value = "saveMembers")
    @ResponseBody
    public Map<String, Object> saveMembers(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
        Integer orgId = user.getDefaultOrgId();
        Integer projectId = json.getInteger("projectId");

        List<TstProjectRoleEntityRelation> entityInRoles = projectRoleEntityRelationService.batchSavePers(json, orgId);

        TstProject project = projectService.get(projectId);
        historyService.create(projectId, user, Constant.MsgType.update.msg,
                TstHistory.TargetType.project_member, projectId, project.getName());

        ret.put("entityInRoles", entityInRoles);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @PostMapping(value = "changeRole")
    @ResponseBody
    public Map<String, Object> changeRole(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

        List<TstProjectRoleEntityRelation> entityInRoles = projectRoleEntityRelationService.changeRolePers(json);

        pushSettingsService.pushPrjSettings(user);

        ret.put("entityInRoles", entityInRoles);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

}
