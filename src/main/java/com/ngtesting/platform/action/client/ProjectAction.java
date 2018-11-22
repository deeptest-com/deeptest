package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.*;
import com.ngtesting.platform.service.intf.AuthService;
import com.ngtesting.platform.service.intf.HistoryService;
import com.ngtesting.platform.service.intf.ProjectService;
import com.ngtesting.platform.service.intf.TestPlanService;
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
public class ProjectAction extends BaseAction {
    @Autowired
    private ProjectService projectService;

    @Autowired
    private TestPlanService planService;
    @Autowired
    private HistoryService historyService;
    @Autowired
    AuthService authService;

    @ResponseBody
    @PostMapping("/list")
    public Object list(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer orgId = user.getDefaultOrgId();

        String keywords = json.getString("keywords");
        Boolean disabled = json.getBoolean("disabled");

        List<TstProject> vos = projectService.list(orgId, user.getId(), keywords, disabled);

        ret.put("data", vos);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());

        return ret;
    }

    @ResponseBody
    @PostMapping("/get")
    public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

        Integer projectId = json.getInteger("id");

        if (projectId != null) {
            TstProject project = projectService.get(projectId);
            if (authService.noProjectAndProjectGroupPrivilege(user.getId(), project)) {
                return authFail();
            }

            ret.put("data", project);
        }

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @ResponseBody
    @PostMapping("/getInfo")
    public Map<String, Object> getInfo(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer orgId = user.getDefaultOrgId();

        Integer projectId = json.getInteger("id");

        if (projectId != null) {
            TstProject project = projectService.get(projectId);
            if (authService.noProjectAndProjectGroupPrivilege(user.getId(), project)) {
                return authFail();
            }

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
        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer orgId = user.getDefaultOrgId();
        Integer projectId = json.getInteger("id");

        TstProject po = projectService.getWithPrivs(projectId, user.getId());
        if (authService.noProjectAndProjectGroupPrivilege(user.getId(), po)) {
            return authFail();
        }

        List<TstPlan> planPos = planService.listByProject(projectId, po.getType());
        planService.genVos(planPos);

        List<TstHistory> historyPos = historyService.listByProject(projectId, po.getType());
        Map<String, List<TstHistory>> historyVos = historyService.genVosByDate(historyPos);

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        ret.put("project", po);
        ret.put("plans", planPos);
        ret.put("histories", historyVos);

        return ret;
    }

    @ResponseBody
    @PostMapping("/save")
    public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer orgId = user.getDefaultOrgId();
        Integer userId = user.getId();

        TstProject vo = json.getObject("model", TstProject.class);

        TstProject po = projectService.save(vo, orgId, user);
        if (po == null) {
            return authFail();
        }

        ret.put("data", vo);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @PostMapping(value = "delete")
    @ResponseBody
    public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

        Integer projectId = json.getInteger("id");
        TstProject project = projectService.get(projectId);
        if (authService.noProjectAndProjectGroupPrivilege(user.getId(), project)) {
            return authFail();
        }

        projectService.delete(projectId, user);

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    // 来源于前端上下文的变化
    @ResponseBody
    @PostMapping("/change")
    public Map<String, Object> change(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer projectId = json.getInteger("id");

        TstProject vo = projectService.changeDefaultPrj(user, projectId);
        if (vo == null) {
            return authFail();
        }

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        ret.put("data", vo);

        return ret;
    }

}
