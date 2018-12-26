package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.IsuTag;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.IssueTagService;
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
@RequestMapping(Constant.API_PATH_CLIENT + "issue_tag/")
public class IssueTagAction extends BaseAction {
    @Autowired
    IssueTagService issueTagService;

    @RequestMapping(value = "search", method = RequestMethod.POST)
    @ResponseBody
    @PrivPrj
    public Map<String, Object> search(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<>();
        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer orgId = user.getDefaultOrgId();

        Integer issueId = json.getInteger("issueId");
        String keywords = json.getString("keywords");
        List<Integer> exceptIds = json.getObject("exceptIds", List.class);

        List<IsuTag> tags = issueTagService.search(issueId, orgId, keywords, exceptIds);

        ret.put("data", tags);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @RequestMapping(value = "save", method = RequestMethod.POST)
    @ResponseBody
    @PrivPrj(perms = {"issue-maintain"})
    public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<>();
        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

        Integer issueId = json.getInteger("issueId");
        List<Map> tags = json.getObject("tags", List.class);

        issueTagService.save(issueId, tags, user);

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

}
