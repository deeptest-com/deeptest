package com.ngtesting.platform.action;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.service.CaseAttachmentService;
import com.ngtesting.platform.util.AuthPassport;
import com.ngtesting.platform.vo.TestCaseAttachmentVo;
import com.ngtesting.platform.vo.UserVo;
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

	@AuthPassport(validate = true)
	@RequestMapping(value = "uploadAttachment", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> uploadAttachment(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

        UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

        Long caseId = json.getLong("caseId");
		String path = json.getString("path");
        String name = json.getString("name");

		caseAttachmentService.uploadAttachmentPers(caseId, name, path, userVo);
        List<TestCaseAttachmentVo> vos = caseAttachmentService.listByCase(caseId);

        ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@AuthPassport(validate = true)
	@RequestMapping(value = "removeAttachment", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> removeAttachment(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

        UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

        Long caseId = json.getLong("caseId");
		Long id = json.getLong("id");

        caseAttachmentService.removeAttachmentPers(id, userVo);
        List<TestCaseAttachmentVo> vos = caseAttachmentService.listByCase(caseId);

		ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
