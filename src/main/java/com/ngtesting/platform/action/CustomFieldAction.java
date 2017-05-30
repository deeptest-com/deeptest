package com.ngtesting.platform.action;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

import javax.servlet.http.HttpServletRequest;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.ResponseBody;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.TestCustomField;
import com.ngtesting.platform.entity.TestOrgRole;
import com.ngtesting.platform.service.CustomFieldService;
import com.ngtesting.platform.service.TestProjectService;
import com.ngtesting.platform.util.AuthPassport;
import com.ngtesting.platform.util.Constant;
import com.ngtesting.platform.util.Constant.RespCode;
import com.ngtesting.platform.vo.CaseTypeVo;
import com.ngtesting.platform.vo.OrgPrivilegeVo;
import com.ngtesting.platform.vo.OrgRoleVo;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.CustomFieldVo;
import com.ngtesting.platform.vo.TestProjectVo;
import com.ngtesting.platform.vo.UserVo;


@Controller
@RequestMapping(Constant.API_PATH_CLIENT + "custom_field/")
public class CustomFieldAction extends BaseAction {
	@Autowired
	CustomFieldService customFieldService;
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "list", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> list(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		Long orgId = userVo.getDefaultOrgId();
		
		List<CustomFieldVo> vos = customFieldService.listVos(orgId);
		
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		ret.put("data", vos);
		return ret;
	}
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "get", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		Long orgId = userVo.getDefaultOrgId();
		
		Long customFieldId = json.getLong("id");
		
		CustomFieldVo vo;
		if (customFieldId == null) {
			vo = new CustomFieldVo();
		} else {
			TestCustomField po = (TestCustomField) customFieldService.get(TestCustomField.class, customFieldId);
			vo = customFieldService.genVo(po);
		}
		
		List<String> applyToList = customFieldService.listApplyTo();
		List<String> typeList = customFieldService.listType();
		List<String> formatList = customFieldService.listFormat();
		List<TestProjectVo> projectList = customFieldService.listProjectsForField(orgId, customFieldId);
        
        ret.put("data", vo);
        ret.put("applyToList", applyToList);
        ret.put("typeList", typeList);
        ret.put("formatList", formatList);
        ret.put("projects", projectList);
        
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
		
		CustomFieldVo customField = JSON.parseObject(JSON.toJSONString(json.get("model")), CustomFieldVo.class);
		List<TestProjectVo> projects = (List<TestProjectVo>) json.get("relations");
		
		TestCustomField po = customFieldService.save(customField, orgId);
		boolean success = customFieldService.saveRelationsProjects(po.getId(), projects);
		
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "delete", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		Long id = json.getLong("id");
		
		boolean success = customFieldService.delete(id);
		
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
		
		boolean success = customFieldService.changeOrderPers(id, act);
		
		List<CustomFieldVo> vos = customFieldService.listVos(orgId);
		
        ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		
		return ret;
	}
	
}
