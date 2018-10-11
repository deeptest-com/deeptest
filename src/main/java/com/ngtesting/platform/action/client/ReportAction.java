package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstProject;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.ProjectService;
import com.ngtesting.platform.service.TestReportService;
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
@RequestMapping(Constant.API_PATH_CLIENT + "report/")
public class ReportAction extends BaseAction {

	@Autowired
    TestReportService reportService;
    @Autowired
    ProjectService projectService;

    @PostMapping(value = "org")
    @ResponseBody
    public Map<String, Object> org(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
        Map<String, Object> data = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

        Integer orgId = json.getInteger("orgId");
        if (userNotInOrg(user.getId(), orgId)) {
            return authFail();
        }

        Map<String, List<Object>> designReport =
                reportService.chartDesignProgressByProject(orgId, TstProject.ProjectType.org, 14);
        Map<String, List<Object>> exeReport =
                reportService.chartExcutionProcessByProject(orgId, TstProject.ProjectType.org, 14);

        data.put("design", designReport);
        data.put("exe", exeReport);

        ret.put("data", data);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @RequestMapping(value = "project")
    @ResponseBody
    public Map<String, Object> project(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
        Map<String, Object> data = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

        Integer projectId = json.getInteger("projectId");
        TstProject prj = projectService.get(projectId);
        if (prj.getType().equals(TstProject.ProjectType.project) &&
                userNotInProject(user.getId(), projectId)) {
            return authFail();
        }

        Map<String, List<Object>> designReport =
                reportService.chartDesignProgressByProject(projectId, prj.getType(), 14);
        Map<String, List<Object>> exeReport =
                reportService.chartExcutionProcessByProject(projectId, prj.getType(), 14);

        data.put("design", designReport);
        data.put("exe", exeReport);

        ret.put("data", data);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @PostMapping(value = "plan")
    @ResponseBody
    public Map<String, Object> plan(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
        Map<String, Object> data = new HashMap<String, Object>();

        Integer planId = json.getInteger("planId");

        List<Map<Object, Object>> resultReport =
                reportService.chartExecutionResultByPlan(planId);
        Map<String, List<Object>> processReport =
                reportService.chartExecutionProcessByPlan(planId, 14);
        Map<String, Object> processByUserReport =
                reportService.chartExecutionProcessByPlanUser(planId, 14);
        Map<String, Object> progressReport =
                reportService.chartExecutionProgressByPlan(planId, 14);

        data.put("result", resultReport);
        data.put("process", processReport);
        data.put("processByUser", processByUserReport);
        data.put("progress", progressReport);

        ret.put("data", data);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());

        return ret;
    }

}
