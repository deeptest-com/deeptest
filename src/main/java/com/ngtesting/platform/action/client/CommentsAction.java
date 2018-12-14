package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.CaseCommentsService;
import com.ngtesting.platform.service.intf.IssueCommentsService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.ResponseBody;

import javax.servlet.http.HttpServletRequest;
import java.util.HashMap;
import java.util.Map;

@Controller
@RequestMapping(Constant.API_PATH_CLIENT + "comments/")
public class CommentsAction extends BaseAction {
    @Autowired
    CaseCommentsService caseCommentsService;

    @Autowired
    IssueCommentsService issueCommentsService;

    @RequestMapping(value = "save", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

        Object vo = null;
        if ("case".equals(json.getString("modelType"))) {
            vo = caseCommentsService.save(json, user);
        } else if ("issue".equals(json.getString("modelType"))) {
            vo = issueCommentsService.save(json, user);
        }

        if (vo == null) {
            return authFail();
        }

        ret.put("data", vo);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @RequestMapping(value = "delete", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Boolean result = null;

        if ("case".equals(json.getString("modelType"))) {
            result = caseCommentsService.delete(json.getInteger("id"), user);
        } else if ("issue".equals(json.getString("modelType"))) {
            result = issueCommentsService.delete(json.getInteger("id"), user);
        }

        if (!result) {
            return authFail();
        }

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }
}
