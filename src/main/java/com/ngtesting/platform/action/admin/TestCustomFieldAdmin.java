package com.ngtesting.platform.action.admin;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstCustomField;
import com.ngtesting.platform.model.TstCustomFieldProjectRelation;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.TestCustomFieldProjectRelationService;
import com.ngtesting.platform.service.TestCustomFieldService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.ResponseBody;

import javax.servlet.http.HttpServletRequest;
import java.util.HashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;

@Controller
@RequestMapping(Constant.API_PATH_ADMIN + "test_custom_field/")
public class TestCustomFieldAdmin extends BaseAction {
	@Autowired
	TestCustomFieldService customFieldService;
	@Autowired
	TestCustomFieldProjectRelationService customFieldProjectRelationService;

	@RequestMapping(value = "list", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> list(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = user.getDefaultOrgId();

		List<TstCustomField> vos = customFieldService.list(orgId); // 总是取当前活动org的，不需要再鉴权

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		ret.put("data", vos);
		return ret;
	}

	@RequestMapping(value = "get", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = user.getDefaultOrgId();

		Integer id = json.getInteger("id");

		TstCustomField vo = null;
		if (id == null) {
			vo = new TstCustomField();
			vo.setMyColumn(customFieldService.getLastUnusedColumn(orgId));
//			vo.setCode(UUID.randomUUID().toString());
		} else {
			vo = customFieldService.get(id, orgId);
		}

		if (vo == null) { // 当对象不是默认org的，此处为空
			return authFail();
		}

		if (vo.getMyColumn() == null) {
            ret.put("code", Constant.RespCode.BIZ_FAIL.getCode());
            ret.put("msg", "自定义字段不能超过20个");
        }

		List<String> applyToList = customFieldService.listApplyTo();
		List<String> typeList = customFieldService.listType();
		List<String> formatList = customFieldService.listFormat();
		List<TstCustomFieldProjectRelation> relations =
                customFieldProjectRelationService.listRelationsByField(orgId, id);

        ret.put("data", vo);
        ret.put("applyToList", applyToList);
        ret.put("typeList", typeList);
        ret.put("formatList", formatList);
        ret.put("relations", relations);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = userVo.getDefaultOrgId();

		TstCustomField customField = JSON.parseObject(JSON.toJSONString(json.get("model")), TstCustomField.class);

		TstCustomField po = customFieldService.save(customField, orgId);
        if (po == null) { // 当对象不是默认org的，update的结果会返回空
            return authFail();
        }

		List<TstCustomFieldProjectRelation> relations = (List<TstCustomFieldProjectRelation>) json.get("relations");
        if (po.getGlobal()) {
			relations = new LinkedList<>();
		}
		customFieldProjectRelationService.saveRelationsByField(orgId, po.getId(), relations);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "delete", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = user.getDefaultOrgId();

		Integer id = json.getInteger("id");

		Boolean result = customFieldService.delete(id, orgId);
        if (!result) { // 当对象不是默认org的，结果会返回false
            return authFail();
        }

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "changeOrder", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> changeOrder(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = user.getDefaultOrgId();

		Integer id = json.getInteger("id");
		String act = json.getString("act");

        Boolean result = customFieldService.changeOrderPers(id, act, orgId);
        if (!result) { // 当对象不是默认org的，结果会返回false
            return authFail();
        }

		List<TstCustomField> vos = customFieldService.list(orgId);

        ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());

		return ret;
	}

}
