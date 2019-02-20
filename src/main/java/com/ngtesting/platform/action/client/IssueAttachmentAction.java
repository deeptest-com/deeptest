package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.IsuAttachment;
import com.ngtesting.platform.model.IsuHistory;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.IssueAttachmentService;
import com.ngtesting.platform.service.intf.IssueHistoryService;
import com.ngtesting.platform.service.intf.IssueService;
import com.ngtesting.platform.service.intf.MsgService;
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
@RequestMapping(Constant.API_PATH_CLIENT + "issue_attachment/")
public class IssueAttachmentAction extends BaseAction {

    @Autowired
    IssueAttachmentService issueAttachmentService;
    @Autowired
    IssueHistoryService issueHistoryService;

    @Autowired
    IssueService issueService;
    @Autowired
    MsgService msgService;

    @RequestMapping(value = "save", method = RequestMethod.POST)
    @ResponseBody
    @PrivPrj(perms = {"issue-maintain"})
    public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer prjId = user.getDefaultPrjId();

        Integer issueId = json.getInteger("issueId");
        String path = json.getString("path");
        String name = json.getString("name");

        Boolean result = issueAttachmentService.save(issueId, name, path, user);
        if (!result) {
            return authFail();
        }

        List<IsuAttachment> attachments = issueAttachmentService.query(issueId);
        List<IsuHistory> histories = issueHistoryService.query(issueId);

        ret.put("attachments", attachments);
        ret.put("histories", histories);

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @RequestMapping(value = "remove", method = RequestMethod.POST)
    @ResponseBody
    @PrivPrj(perms = {"issue-maintain"})
    public Map<String, Object> remove(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer prjId = user.getDefaultPrjId();

        Integer issueId = json.getInteger("issueId");
        Integer id = json.getInteger("id");

        Boolean result = issueAttachmentService.delete(id, user);

        List<IsuAttachment> attachments = issueAttachmentService.query(issueId);
        List<IsuHistory> histories = issueHistoryService.query(issueId);

        ret.put("attachments", attachments);
        ret.put("histories", histories);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

}
