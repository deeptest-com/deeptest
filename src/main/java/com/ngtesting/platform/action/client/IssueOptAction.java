package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.IssueDynamicFormService;
import com.ngtesting.platform.service.intf.IssueFieldService;
import com.ngtesting.platform.service.intf.IssueOptService;
import com.ngtesting.platform.servlet.PrivPrj;
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
@RequestMapping(Constant.API_PATH_CLIENT + "issue_opt/")
public class IssueOptAction extends BaseAction {
    @Autowired
    IssueOptService issueOptService;
    @Autowired
    IssueFieldService fieldService;
	@Autowired
	IssueDynamicFormService dynamicFormService;

    @RequestMapping(value = "statusTran", method = RequestMethod.POST)
    @ResponseBody
    @PrivPrj(perms = {"issue-maintain"})
    public Map<String, Object> statusTran(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<>();
        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer prjId = user.getDefaultPrjId();

        Integer id = json.getInteger("id");
        Integer dictStatusId = json.getInteger("dictStatusId");
        String dictStatusName = json.getString("dictStatusName");

        issueOptService.statusTran(id, dictStatusId, dictStatusName, user);

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @PrivPrj(perms = {"issue-maintain"})
    @RequestMapping(value = "assign", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> assign(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<>();
        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

        Integer id = json.getInteger("id");
        Integer userId = json.getInteger("userId");
        String comments = json.getString("comments");

        issueOptService.assign(id, userId, comments, user);

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

}
