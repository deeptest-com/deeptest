package com.ngtesting.platform.action.admin;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.CustomField;
import com.ngtesting.platform.model.CustomFieldOption;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.IssueCustomFieldOptionService;
import com.ngtesting.platform.service.intf.CustomFieldService;
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
@RequestMapping(Constant.API_PATH_ADMIN + "custom_field_option/")
public class CustomFieldOptionAdmin extends BaseAction {
	@Autowired
	IssueCustomFieldOptionService customFieldOptionService;

    @Autowired
    CustomFieldService customFieldService;

	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = user.getDefaultOrgId();

		Integer fieldId = json.getInteger("fieldId");

		if (fieldId == null) { // 未保存字段
            CustomField customField = JSON.parseObject(JSON.toJSONString(json.get("field")), CustomField.class);

            customFieldService.save(customField, orgId);
            fieldId = customField.getId();

            ret.put("field", customField);
        }

		CustomFieldOption option = JSON.parseObject(JSON.toJSONString(json.getJSONObject("model")), CustomFieldOption.class);
		option.setFieldId(fieldId);
		CustomFieldOption po = customFieldOptionService.save(option, orgId);
		if (po == null) { // 当所属fieldId不是默认org的，结果会返回空
			return authFail();
		}

		List<CustomFieldOption> vos = customFieldOptionService.list(fieldId, orgId);

		ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "delete", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = user.getDefaultOrgId();

		Integer fieldId = json.getInteger("fieldId");
		Integer id = json.getInteger("id");

		Boolean result = customFieldOptionService.delete(id, fieldId, orgId);
		if(!result) {  // 当找不到option对象、或option所属fieldId不是默认org的，结果会返回空
			return authFail();
		}

		List<CustomFieldOption> vos = customFieldOptionService.list(fieldId, orgId);

		ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}


	@RequestMapping(value = "changeOrder", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> changeOrder(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = user.getDefaultOrgId();

		Integer fieldId = json.getInteger("fieldId");
		Integer id = json.getInteger("id");
		String act = json.getString("act");

		Boolean result = customFieldOptionService.changeOrder(id, act, fieldId, orgId);
		if(!result) { // 当找不到option对象、或option所属fieldId不是默认org的，结果会返回空
			return authFail();
		}

		List<CustomFieldOption> vos = customFieldOptionService.list(fieldId, orgId);

		ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());

		return ret;
	}

	@RequestMapping(value = "setDefault", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> setDefault(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = user.getDefaultOrgId();

        Integer fieldId = json.getInteger("fieldId");
		Integer id = json.getInteger("id");

        customFieldOptionService.setDefault(id, fieldId, orgId);  // 涵盖项目设置WS推送消息

        List<CustomFieldOption> vos = customFieldOptionService.list(fieldId, orgId);

        ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());

		return ret;
	}

}
