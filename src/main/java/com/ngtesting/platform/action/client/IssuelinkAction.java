package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.IsuLinkReason;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.IssueLinkService;
import com.ngtesting.platform.service.intf.MsgService;
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
@RequestMapping(Constant.API_PATH_CLIENT + "issue_link/")
public class IssuelinkAction extends BaseAction {
    @Autowired
    IssueLinkService issueLinkService;

    @Autowired
    MsgService msgService;

    @RequestMapping(value = "link", method = RequestMethod.POST)
    @PrivPrj(perms = {"issue:view"})
    public Map<String, Object> link(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();

        Integer srcIssueId = json.getInteger("srcIssueId");
        Integer dictIssueId = json.getInteger("dictIssueId");
        Integer reasonId = json.getInteger("reasonId");
        String reasonName = json.getString("reasonName");

        issueLinkService.link(srcIssueId, dictIssueId, reasonId, reasonName, user);

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @RequestMapping(value = "listIssueLinkReasons", method = RequestMethod.POST)
    @PrivPrj(perms = {"issue:view"})
    public Map<String, Object> listIssueLinkReasons(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        List<IsuLinkReason> issueLinkReasons = issueLinkService.listLinkReason();

        ret.put("data", issueLinkReasons);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

}
