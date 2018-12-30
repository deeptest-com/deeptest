package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.dao.CaseInTaskAttachmentDao;
import com.ngtesting.platform.dao.CaseInTaskHistoryDao;
import com.ngtesting.platform.model.*;
import com.ngtesting.platform.service.intf.CaseInTaskAttachmentService;
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
@RequestMapping(Constant.API_PATH_CLIENT + "case_in_task_attachment/")
public class CaseInTaskAttachmentAction extends BaseAction {
	@Autowired
	CaseInTaskAttachmentService caseInTaskAttachmentService;
    @Autowired
	CaseInTaskHistoryDao caseInTaskHistoryDao;
	@Autowired
	CaseInTaskAttachmentDao caseInTaskAttachmentDao;

	@RequestMapping(value = "upload", method = RequestMethod.POST)
	@ResponseBody
	@PrivPrj(perms = {"test_case-maintain"})
	public Map<String, Object> upload(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

        Integer caseInTaskId = json.getInteger("caseInTaskId");
		String path = json.getString("path");
        String name = json.getString("name");

		Boolean result = caseInTaskAttachmentService.save(caseInTaskId, name, path, user);
		if (!result) {
			return authFail();
		}

        List<TstCaseInTaskAttachment> attachments = caseInTaskAttachmentDao.query(caseInTaskId);
        List<TstCaseInTaskHistory> histories = caseInTaskHistoryDao.query(caseInTaskId);

        ret.put("attachments", attachments);
        ret.put("histories", histories);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}


	@RequestMapping(value = "remove", method = RequestMethod.POST)
	@ResponseBody
    @PrivPrj(perms = {"test_case-maintain"})
	public Map<String, Object> remove(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

        Integer caseInTaskId = json.getInteger("caseInTaskId");
		Integer id = json.getInteger("id");

		Boolean result = caseInTaskAttachmentService.delete(id, user);
		if (!result) {
			return authFail();
		}

        List<TstCaseInTaskAttachment> attachments = caseInTaskAttachmentDao.query(caseInTaskId);
        List<TstCaseInTaskHistory> histories = caseInTaskHistoryDao.query(caseInTaskId);

        ret.put("attachments", attachments);
        ret.put("histories", histories);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
