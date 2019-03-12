package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstCaseInTask;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.*;
import com.ngtesting.platform.servlet.PrivPrj;
import org.apache.shiro.SecurityUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

import javax.servlet.http.HttpServletRequest;
import java.util.HashMap;
import java.util.List;
import java.util.Map;


@RestController
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
    CustomFieldService customFieldService;

    @RequestMapping(value = "query", method = RequestMethod.POST)
    @PrivPrj(perms = {"test_case-view", "test_plan-view", "test_task-view"})
    public Map<String, Object> query(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();

        Integer orgId = user.getDefaultOrgId();
        Integer prjId = user.getDefaultPrjId();

        Integer taskId = json.getInteger("taskId");

        List<TstCaseInTask> vos = caseInTaskService.query(taskId, prjId);

//        Map<String, Object> map = customFieldService.fetchProjectFieldForCase(orgId, prjId);

        ret.put("data", vos);
//        ret.put("customFields", map.get("fields"));
//        ret.put("casePropMap", map.get("props"));
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @RequestMapping(value = "get", method = RequestMethod.POST)
    @PrivPrj(perms = {"test_case-view", "test_plan-view", "test_task-view"})
    public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();
        Integer prjId = user.getDefaultPrjId();

        Integer id = json.getInteger("id");

        TstCaseInTask vo = caseInTaskService.getDetail(id, prjId);

        ret.put("data", vo);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @RequestMapping(value = "rename", method = RequestMethod.POST)
    @PrivPrj(perms = {"test_case-maintain"})
    public Map<String, Object> rename(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();

        TstCaseInTask testCase = caseInTaskService.rename(json, user);
        if (testCase == null) {
            return authorFail();
        }

        ret.put("data", testCase);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @RequestMapping(value = "setResult", method = RequestMethod.POST)
    @PrivPrj(perms = {"test_task-exe"})
    public Map<String, Object> setResult(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();

        Integer caseInTaskId = json.getInteger("id");
        Integer caseId = json.getInteger("caseId");
        String result = json.getString("result");
        String status = json.getString("status");
        Integer nextId = json.getInteger("nextId");

        TstCaseInTask testCase = caseInTaskService.setResult(caseInTaskId, caseId, result, status, nextId, user);
        if (testCase == null) {
            return authorFail();
        }

        ret.put("data", testCase);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

}
