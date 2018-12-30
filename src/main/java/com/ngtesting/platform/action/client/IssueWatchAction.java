package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.IssueWatchService;
import com.ngtesting.platform.servlet.PrivPrj;
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
@RequestMapping(Constant.API_PATH_CLIENT + "issue_watch/")
public class IssueWatchAction extends BaseAction {
    @Autowired
    IssueWatchService issueWatchService;

    @RequestMapping(value = "list", method = RequestMethod.POST)
    @ResponseBody
    @PrivPrj
    public Map<String, Object> list(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<>();
        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer prjId = user.getDefaultPrjId();

        Integer issueId = json.getInteger("issueId");

        List<Map> users = issueWatchService.list(issueId);

        ret.put("data", users);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @RequestMapping(value = "search", method = RequestMethod.POST)
    @ResponseBody
    @PrivPrj
    public Map<String, Object> search(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<>();
        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer prjId = user.getDefaultPrjId();

        Integer issueId = json.getInteger("issueId");
        String keywords = json.getString("keywords");
        List<Integer> exceptIds = json.getObject("exceptIds", List.class);

        List<TstUser> users = issueWatchService.search(issueId, prjId, keywords, exceptIds, user);

        ret.put("data", users);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @RequestMapping(value = "batchSave", method = RequestMethod.POST)
    @ResponseBody
    @PrivPrj(perms = {"issue-maintain"})
    public Map<String, Object> batchSave(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<>();
        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

        Integer issueId = json.getInteger("issueId");
        List<Integer> userIds = json.getObject("userIds", List.class);

        issueWatchService.batchSave(issueId, userIds, user);

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @RequestMapping(value = "remove", method = RequestMethod.POST)
    @ResponseBody
    @PrivPrj(perms = {"issue-maintain"})
    public Map<String, Object> remove(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<>();
        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

        Integer id = json.getInteger("id");
        Integer issueId = json.getInteger("issueId");

        issueWatchService.remove(id, issueId, user);

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @RequestMapping(value = "watch", method = RequestMethod.POST)
    @ResponseBody
    @PrivPrj
    public Map<String, Object> watch(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<>();
        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

        Integer id = json.getInteger("id");
        Boolean status = json.getBoolean("status");

        issueWatchService.watch(id, user, status);

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

}
