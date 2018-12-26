package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstProject;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.ProjectService;
import com.ngtesting.platform.service.intf.ReportIssueService;
import com.ngtesting.platform.servlet.PrivOrg;
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
@RequestMapping(Constant.API_PATH_CLIENT + "report_issue/")
public class ReportIssueAction extends BaseAction {

    @Autowired
    ReportIssueService reportIssueService;

    @Autowired
    ProjectService projectService;

    @PostMapping(value = "orgIssue")
    @ResponseBody
    @PrivOrg
    public Map<String, Object> orgIssue(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
        Map<String, Object> data = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer orgId = user.getDefaultOrgId();

        Map trendData =
                reportIssueService.chartIssueTrend(orgId, TstProject.ProjectType.org, 14);
        Map ageData =
                reportIssueService.chartIssueAgeByOrgOrGroup(orgId, TstProject.ProjectType.org, 7);
        List distribDataByPriority =
                reportIssueService.chartIssueDistribByPriority(orgId, TstProject.ProjectType.org);
        List distribDataByStatus =
                reportIssueService.chartIssueDistribByStatus(orgId, TstProject.ProjectType.org);

        data.put("trend", trendData);
        data.put("age", ageData);
        data.put("distribByPriority", distribDataByPriority);
        data.put("distribByStatus", distribDataByStatus);

        ret.put("data", data);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @RequestMapping(value = "projectIssue")
    @ResponseBody
    @PrivPrj
    public Map<String, Object> projectIssue(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
        Map<String, Object> data = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer orgId = user.getDefaultOrgId();
        Integer prjId = json.getInteger("projectId");

        TstProject prj = projectService.get(prjId);

        Map trendData =
                reportIssueService.chartIssueTrend(prjId, prj.getType(), 14);
        List distribDataByPriority =
                reportIssueService.chartIssueDistribByPriority(prjId, prj.getType());
        List distribDataByStatus =
                reportIssueService.chartIssueDistribByStatus(prjId, prj.getType());

        Map ageData;
        if (prj.getType().equals(TstProject.ProjectType.project)) {
            ageData = reportIssueService.chartIssueAgeByProject(prjId, 7, orgId);
        } else {
            ageData = reportIssueService.chartIssueAgeByOrgOrGroup(prjId, prj.getType(), 7);
        }

        data.put("trend", trendData);
        data.put("age", ageData);
        data.put("distribByPriority", distribDataByPriority);
        data.put("distribByStatus", distribDataByStatus);

        ret.put("data", data);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

}
