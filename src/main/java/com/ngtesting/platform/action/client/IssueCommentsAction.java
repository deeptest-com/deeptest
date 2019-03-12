package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.IsuComments;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.IssueCommentsService;
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
@RequestMapping(Constant.API_PATH_CLIENT + "issue_comments/")
public class IssueCommentsAction extends BaseAction {
    @Autowired
    IssueCommentsService issueCommentsService;

    @RequestMapping(value = "save", method = RequestMethod.POST)
    @PrivPrj(perms = {"issue-maintain"})
    public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();

        IsuComments vo = issueCommentsService.save(json, user);

        if (vo == null) {
            return authorFail();
        }

        ret.put("data", vo);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @RequestMapping(value = "delete", method = RequestMethod.POST)
    @PrivPrj(perms = {"issue-maintain"})
    public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();

        Integer id = json.getInteger("id");

        Boolean result = issueCommentsService.delete(id, user);
        if (!result) {
            return authorFail();
        }

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }
}
