package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstCaseStep;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.CaseStepService;
import com.ngtesting.platform.servlet.PrivPrj;
import org.apache.shiro.SecurityUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

import javax.servlet.http.HttpServletRequest;
import java.util.HashMap;
import java.util.Map;

@RestController
@RequestMapping(Constant.API_PATH_CLIENT + "case_step/")
public class CaseStepAction extends BaseAction {
    @Autowired
    CaseStepService caseStepService;

    @RequestMapping(value = "save", method = RequestMethod.POST)
    @PrivPrj(perms = {"test_case-maintain"})
    public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();

        TstCaseStep po = caseStepService.save(json, user);
        if (po == null) {
            return authorFail();
        }

        ret.put("data", po);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @RequestMapping(value = "delete", method = RequestMethod.POST)
    @PrivPrj(perms = {"test_case-maintain"})
    public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();

        Boolean result = caseStepService.delete(json.getInteger("id"), user);
        if (!result) {
            return authorFail();
        }

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @RequestMapping(value = "up", method = RequestMethod.POST)
    @PrivPrj(perms = {"test_case-maintain"})
    public Map<String, Object> up(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();

        Boolean result = caseStepService.changeOrder(json, "up", user);
        if (!result) {
            return authorFail();
        }

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @RequestMapping(value = "down", method = RequestMethod.POST)
    @PrivPrj(perms = {"test_case-maintain"})
    public Map<String, Object> down(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();

        Boolean result = caseStepService.changeOrder(json, "down", user);
        if (!result) {
            return authorFail();
        }

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }
}
