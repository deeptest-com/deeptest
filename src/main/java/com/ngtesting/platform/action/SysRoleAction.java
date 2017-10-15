package com.ngtesting.platform.action;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.SysRole;
import com.ngtesting.platform.service.SysRoleService;
import com.ngtesting.platform.util.AuthPassport;
import com.ngtesting.platform.util.Constant;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.RoleVo;
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
@RequestMapping(Constant.API_PATH_CLIENT + "role/")
public class SysRoleAction extends BaseAction {
	@Autowired
    SysRoleService roleService;

	@AuthPassport(validate = true)
	@RequestMapping(value = "list", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> list(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		Long orgId = json.getLong("orgId");

		String keywords = json.getString("keywords");
		String disabled = json.getString("disabled");
		int page = json.getInteger("page") == null? 0: json.getInteger("page") - 1;
		int pageSize = json.getInteger("pageSize") == null? Constant.PAGE_SIZE: json.getInteger("pageSize");

		Page pageData = roleService.listByPage(orgId, keywords, disabled, page, pageSize);
		List<RoleVo> vos = roleService.genVos(pageData.getItems());

		ret.put("collectionSize", pageData.getTotal());
        ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@AuthPassport(validate = true)
	@RequestMapping(value = "get", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject req) {
		Map<String, Object> ret = new HashMap<String, Object>();
		String accountId = req.getString("id");

		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

		SysRole po = (SysRole) roleService.get(SysRole.class, Long.valueOf(accountId));
		RoleVo vo = roleService.genVo(po);

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
		RoleVo vo = json.getObject("role", RoleVo.class);

		SysRole po = roleService.save(vo, orgId);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@AuthPassport(validate = true)
	@RequestMapping(value = "delete", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject to) {
		Map<String, Object> ret = new HashMap<String, Object>();

		boolean success = roleService.delete(to.getLong("id"));

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@AuthPassport(validate = true)
	@RequestMapping(value = "disable", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> disable(HttpServletRequest request, @RequestBody JSONObject to) {
		Map<String, Object> ret = new HashMap<String, Object>();

		boolean success = roleService.disable(to.getLong("id"));

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}
}
