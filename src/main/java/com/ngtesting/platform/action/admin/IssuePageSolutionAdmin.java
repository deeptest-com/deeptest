package com.ngtesting.platform.action.admin;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.IsuCustomField;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.IssueCustomFieldService;
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
@RequestMapping(Constant.API_PATH_ADMIN + "issue_solution/")
public class IssuePageSolutionAdmin extends BaseAction {
	@Autowired
    IssueCustomFieldService customFieldService;

	@RequestMapping(value = "load", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> load(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = userVo.getDefaultOrgId();

		List<IsuCustomField> vos = customFieldService.list(orgId);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		ret.put("data", vos);
		return ret;
	}

	@RequestMapping(value = "get", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = userVo.getDefaultOrgId();

		Integer customFieldId = json.getInteger("id");

		IsuCustomField vo = null;
		if (customFieldId == null) {
			vo = new IsuCustomField();
		} else {
			vo = customFieldService.get(customFieldId, orgId);
		}

		if (vo.getMyColumn() == null) {
            ret.put("code", Constant.RespCode.BIZ_FAIL.getCode());
            ret.put("msg", "自定义字段不能超过20个");
        }

		List<String> typeList = customFieldService.listType();
		List<String> formatList = customFieldService.listFormat();

        ret.put("data", vo);
        ret.put("typeList", typeList);
        ret.put("formatList", formatList);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = userVo.getDefaultOrgId();

		IsuCustomField customField = JSON.parseObject(JSON.toJSONString(json.get("model")), IsuCustomField.class);
//		List<TestProjectVo> projects = (List<TestProjectVo>) json.getDetail("relations");
//
//		IsuCustomField po = customFieldService.save(customField, orgId);
//		boolean success = customFieldService.saveRelationsByField(po.getCode(), projects);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "delete", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = userVo.getDefaultOrgId();

		Integer id = json.getInteger("id");

		boolean success = customFieldService.delete(id, orgId);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "changeOrder", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> changeOrder(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = userVo.getDefaultOrgId();
		Integer id = json.getInteger("id");
		String act = json.getString("act");

		boolean success = customFieldService.changeOrderPers(id, act, orgId);

		List<IsuCustomField> vos = customFieldService.list(orgId);

        ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());

		return ret;
	}

}
