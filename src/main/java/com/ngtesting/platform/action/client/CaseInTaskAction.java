package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.*;
import com.ngtesting.platform.service.*;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.ResponseBody;

import javax.servlet.http.HttpServletRequest;
import java.util.HashMap;
import java.util.List;
import java.util.Map;


@Controller
@RequestMapping(Constant.API_PATH_CLIENT + "caseInTask/")
public class CaseInTaskAction extends BaseAction {
    @Autowired
    CaseService caseService;
	@Autowired
    CaseInTaskService caseInTaskService;

    @Autowired
    CaseTypeService caseTypeService;
    @Autowired
    CasePriorityService casePriorityService;
    @Autowired
    TestCustomFieldService customFieldService;

    @RequestMapping(value = "query", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> query(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

        Integer orgId = user.getDefaultOrgId();
        Integer projectId = user.getDefaultPrjId();

        Integer taskId = json.getInteger("taskId");

        List<TstCaseInTask> vos = caseInTaskService.query(taskId, projectId);

        List<TstCaseType> caseTypePos = caseTypeService.list(orgId);
        List<TstCasePriority> casePriorityPos = casePriorityService.list(orgId);
        List<TstCustomField> customFieldList = customFieldService.listForCaseByProject(orgId, projectId);

        ret.put("data", vos);
        ret.put("caseTypeList", caseTypePos);
        ret.put("casePriorityList", casePriorityPos);
        ret.put("customFields", customFieldList);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @RequestMapping(value = "get", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer prjId = user.getDefaultPrjId();

        Integer id = json.getInteger("id");

        TstCaseInTask vo = caseInTaskService.getDetail(id, prjId);

        ret.put("data", vo);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @RequestMapping(value = "rename", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> rename(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

        TstCaseInTask testCase = caseInTaskService.rename(json, user);
        if (testCase == null) {
            return authFail();
        }

        ret.put("data", testCase);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @RequestMapping(value = "setResult", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> setResult(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

        Integer caseInTaskId = json.getInteger("id");
        Integer caseId = json.getInteger("caseId");
        String result = json.getString("result");
        String status = json.getString("status");
        Integer nextId = json.getInteger("nextId");

        TstCaseInTask testCase = caseInTaskService.setResult(caseInTaskId, caseId, result, status, nextId, user);
        if (testCase == null) {
            return authFail();
        }

        ret.put("data", testCase);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

}
