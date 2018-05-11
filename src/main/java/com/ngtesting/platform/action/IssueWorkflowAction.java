package com.ngtesting.platform.action;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.entity.TestCaseType;
import com.ngtesting.platform.service.CasePropertyService;
import com.ngtesting.platform.service.IssueWorkflowService;
import com.ngtesting.platform.util.AuthPassport;
import com.ngtesting.platform.vo.CaseTypeVo;
import com.ngtesting.platform.vo.UserVo;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
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
@RequestMapping(Constant.API_PATH_CLIENT + "issue_workflow/")
public class IssueWorkflowAction extends BaseAction {
	private static final Log log = LogFactory.getLog(CaseTypeAction.class);

	@Autowired
	IssueWorkflowService issueWorkflowService;

	@Autowired
	CasePropertyService casePropertyService;

	@AuthPassport(validate = true)
	@RequestMapping(value = "list", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> list(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		Long orgId = userVo.getDefaultOrgId();

		List<CaseTypeVo> vos = issueWorkflowService.listVos(orgId);

		Map<String,Map<String,String>> casePropertyMap = casePropertyService.getMap(orgId);

        ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());

		return ret;
	}

	@AuthPassport(validate = true)
	@RequestMapping(value = "get", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

		Long id = json.getLong("id");
		if (id == null) {
			ret.put("data", new CaseTypeVo());
			ret.put("code", Constant.RespCode.SUCCESS.getCode());
			return ret;
		}

		TestCaseType po = (TestCaseType) issueWorkflowService.get(TestCaseType.class, id);
		CaseTypeVo vo = issueWorkflowService.genVo(po);
		ret.put("data", vo);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@AuthPassport(validate = true)
	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		Long orgId = userVo.getDefaultOrgId();

		CaseTypeVo vo = json.getObject("model", CaseTypeVo.class);

		TestCaseType po = issueWorkflowService.save(vo, orgId);
		CaseTypeVo projectVo = issueWorkflowService.genVo(po);

		Map<String,Map<String,String>> casePropertyMap = casePropertyService.getMap(orgId);
		ret.put("casePropertyMap", casePropertyMap);

        ret.put("data", projectVo);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@AuthPassport(validate = true)
	@RequestMapping(value = "delete", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		Long id = json.getLong("id");

		issueWorkflowService.delete(id);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@AuthPassport(validate = true)
	@RequestMapping(value = "setDefault", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> setDefault(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		Long orgId = userVo.getDefaultOrgId();
		Long id = json.getLong("id");

		boolean success = issueWorkflowService.setDefaultPers(id, orgId);

		List<CaseTypeVo> vos = issueWorkflowService.listVos(orgId);

        ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());

		return ret;
	}

	@AuthPassport(validate = true)
	@RequestMapping(value = "changeOrder", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> changeOrder(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		Long orgId = userVo.getDefaultOrgId();
		Long id = json.getLong("id");
		String act = json.getString("act");

		boolean success = issueWorkflowService.changeOrderPers(id, act, orgId);

		List<CaseTypeVo> vos = issueWorkflowService.listVos(orgId);

        ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());

		return ret;
	}

}
