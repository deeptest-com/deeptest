package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.SysRole;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.SysRoleService;
import com.ngtesting.platform.vo.Page;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.ResponseBody;

import javax.servlet.http.HttpServletRequest;
import java.util.HashMap;
import java.util.Map;


@Controller
@RequestMapping(Constant.API_PATH_CLIENT + "role/")
public class SysRoleAction extends BaseAction {
	@Autowired
    SysRoleService roleService;

	@RequestMapping(value = "list", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> list(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = json.getInteger("orgId");

		String keywords = json.getString("keywords");
		Boolean disabled = json.getBoolean("disabled");
		int page = json.getInteger("listByPage") == null? 0: json.getInteger("listByPage") - 1;
		int pageSize = json.getInteger("pageSize") == null? Constant.PAGE_SIZE: json.getInteger("pageSize");

		Page pageData = roleService.listByPage(orgId, keywords, disabled, page, pageSize);
//		List<SysRole> vos = roleService.genVos(pageData.getItems());

		ret.put("collectionSize", pageData.getTotal());
//        ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "get", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject req) {
		Map<String, Object> ret = new HashMap<String, Object>();
		String accountId = req.getString("id");

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

//		SysRole po = (SysRole) roleService.getDetail(SysRole.class, Integer.valueOf(accountId));
//		SysRole vo = roleService.genVo(po);
//
//        ret.put("data", vo);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

		Integer orgId = userVo.getDefaultOrgId();
		SysRole vo = json.getObject("role", SysRole.class);

		SysRole po = roleService.save(vo, orgId);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "delete", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject to) {
		Map<String, Object> ret = new HashMap<String, Object>();

		boolean success = roleService.delete(to.getInteger("id"));

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "disable", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> disable(HttpServletRequest request, @RequestBody JSONObject to) {
		Map<String, Object> ret = new HashMap<String, Object>();

		boolean success = roleService.disable(to.getInteger("id"));

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}
}
