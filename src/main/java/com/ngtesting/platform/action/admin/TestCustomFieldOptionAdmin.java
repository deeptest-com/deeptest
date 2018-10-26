package com.ngtesting.platform.action.admin;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstCustomFieldOption;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.TestCustomFieldOptionService;
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
@RequestMapping(Constant.API_PATH_ADMIN + "test_custom_field_option/")
public class TestCustomFieldOptionAdmin extends BaseAction {
	@Autowired
	TestCustomFieldOptionService customFieldOptionService;

	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = user.getDefaultOrgId();

		Integer fieldId = json.getInteger("fieldId");

		TstCustomFieldOption option = JSON.parseObject(JSON.toJSONString(json.getJSONObject("model")), TstCustomFieldOption.class);
        option.setFieldId(fieldId);
        TstCustomFieldOption po = customFieldOptionService.save(option, orgId);
		if (po == null) { // 当所属fieldId不是默认org的，结果会返回空
			return authFail();
		}

        List<TstCustomFieldOption> vos = customFieldOptionService.listVos(fieldId);

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

		Boolean result = customFieldOptionService.delete(id, orgId);
        if(!result) {  // 当找不到option对象、或option所属fieldId不是默认org的，结果会返回空
            return authFail();
        }

        List<TstCustomFieldOption> vos = customFieldOptionService.listVos(fieldId);

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

        List<TstCustomFieldOption> vos = customFieldOptionService.listVos(fieldId);

        ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());

		return ret;
	}

}
