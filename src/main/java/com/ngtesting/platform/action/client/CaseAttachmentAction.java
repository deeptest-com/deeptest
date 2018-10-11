package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.dao.CaseAttachmentDao;
import com.ngtesting.platform.dao.CaseHistoryDao;
import com.ngtesting.platform.model.TstCaseAttachment;
import com.ngtesting.platform.model.TstCaseHistory;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.CaseAttachmentService;
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
@RequestMapping(Constant.API_PATH_CLIENT + "case_attachment/")
public class CaseAttachmentAction extends BaseAction {
	@Autowired
    CaseAttachmentService caseAttachmentService;
    @Autowired
    CaseHistoryDao caseHistoryDao;
	@Autowired
	CaseAttachmentDao caseAttachmentDao;

	@RequestMapping(value = "upload", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> upload(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

        Integer caseId = json.getInteger("caseId");
		String path = json.getString("path");
        String name = json.getString("name");

		Boolean result = caseAttachmentService.save(caseId, name, path, user);
		if (!result) {
			return authFail();
		}

        List<TstCaseAttachment> attachments = caseAttachmentDao.query(caseId);
        List<TstCaseHistory> histories = caseHistoryDao.query(caseId);

        ret.put("attachments", attachments);
        ret.put("histories", histories);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}


	@RequestMapping(value = "remove", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> remove(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

        Integer caseId = json.getInteger("caseId");
		Integer id = json.getInteger("id");

		Boolean result = caseAttachmentService.delete(id, user);
		if (!result) {
			return authFail();
		}

        List<TstCaseAttachment> attachments = caseAttachmentDao.query(caseId);
        List<TstCaseHistory> histories = caseHistoryDao.query(caseId);

        ret.put("attachments", attachments);
        ret.put("histories", histories);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
