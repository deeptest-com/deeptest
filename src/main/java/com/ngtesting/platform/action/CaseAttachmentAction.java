package com.ngtesting.platform.action;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.dao.CaseAttachmentDao;
import com.ngtesting.platform.model.TstCaseAttachment;
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
	CaseAttachmentDao caseAttachmentDao;

	@RequestMapping(value = "uploadAttachment", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> uploadAttachment(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

        TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

        Integer caseId = json.getInteger("caseId");
		String path = json.getString("path");
        String name = json.getString("name");

		caseAttachmentService.uploadAttachmentPers(caseId, name, path, userVo);
        List<TstCaseAttachment> vos = caseAttachmentDao.query(caseId);

        ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}


	@RequestMapping(value = "removeAttachment", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> removeAttachment(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

        TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

        Integer caseId = json.getInteger("caseId");
		Integer id = json.getInteger("id");

        caseAttachmentService.removeAttachmentPers(id, userVo);
        List<TstCaseAttachment> vos = caseAttachmentDao.query(caseId);

		ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
