package com.ngtesting.platform.action;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstHistory;
import com.ngtesting.platform.model.TstPlan;
import com.ngtesting.platform.model.TstProject;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.HistoryService;
import com.ngtesting.platform.service.ProjectService;
import com.ngtesting.platform.service.PushSettingsService;
import com.ngtesting.platform.service.TestPlanService;
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
    @PostMapping("/view")
    public Map<String, Object> view(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
        Integer id = json.getInteger("id");

        TstProject po = projectService.get(id);
        TstProject vo = projectService.genVo(po, null);

        List<TstPlan> planPos = planService.listByProject(id, vo.getType());
        List<TstPlan> planVos = planService.genVos(planPos);

        List<TstHistory> historyPos = historyService.listByProject(id, vo.getType());
        Map<String, List<TstHistory>> historyVos = historyService.genVosByDate(historyPos);

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        ret.put("project", vo);
        ret.put("plans", planVos);
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

}
