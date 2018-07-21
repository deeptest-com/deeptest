package com.ngtesting.platform.action;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstCaseInRun;
import com.ngtesting.platform.model.TstCustomField;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.CaseInRunService;
import com.ngtesting.platform.service.CaseService;
import com.ngtesting.platform.service.CustomFieldService;
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
@RequestMapping(Constant.API_PATH_CLIENT + "caseInRun/")
public class CaseInRunAction extends BaseAction {
    @Autowired
    CaseService caseService;
	@Autowired
    CaseInRunService caseInRunService;

    @Autowired
    CustomFieldService customFieldService;

    @RequestMapping(value = "query", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> query(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        Integer orgId = json.getInteger("orgId");
        Integer projectId = json.getInteger("projectId");
        Integer runId = json.getInteger("runId");

        List<TstCaseInRun> vos = caseInRunService.query(runId);
        List<TstCustomField> customFieldList = customFieldService.listForCaseByProject(orgId, projectId);

        ret.put("data", vos);
        ret.put("customFields", customFieldList);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @RequestMapping(value = "get", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
        Integer orgId = userVo.getDefaultOrgId();
        Integer caseId = json.getInteger("id");

        TstCaseInRun vo = caseInRunService.getById(caseId);

        ret.put("data", vo);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @RequestMapping(value = "setResult", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> setResult(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
        Integer orgId = userVo.getDefaultOrgId();

        Integer caseInRunId = json.getInteger("id");
        String result = json.getString("result");
        String status = json.getString("status");
        Integer nextId = json.getInteger("nextId");

        TstCaseInRun vo = caseInRunService.setResultPers(caseInRunId, result, status, nextId, userVo);

        ret.put("data", vo);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @RequestMapping(value = "rename", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> rename(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

        TstCaseInRun vo = caseInRunService.renamePers(json, userVo);

        ret.put("data", vo);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @RequestMapping(value = "move", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> move(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
        TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

        TstCaseInRun vo = caseInRunService.movePers(json, userVo);

        ret.put("data", vo);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

//    @RequestMapping(value = "delete", method = RequestMethod.POST)
//    @ResponseBody
//    public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject json) {
//        Map<String, Object> ret = new HashMap<String, Object>();
//
//        TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
//
//        Integer entityId = json.getInteger("entityId");
//        TstCaseInRun caseInRun = caseInRunService.removeCaseFromRunPers(entityId, userVo);
//
//        ret.put("code", Constant.RespCode.SUCCESS.getCode());
//        return ret;
//    }

}
