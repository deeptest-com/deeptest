package com.ngtesting.platform.action;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstOrgPrivilegeDefine;
import com.ngtesting.platform.model.TstOrgRole;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.OrgRolePrivilegeService;
import com.ngtesting.platform.service.OrgRoleService;
import com.ngtesting.platform.service.OrgRoleUserService;
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
@RequestMapping(Constant.API_PATH_CLIENT + "org_role/")
public class OrgRoleAction extends BaseAction {
	@Autowired
    OrgRoleService orgRoleService;
	@Autowired
    OrgRolePrivilegeService orgRolePrivilegeService;
    @Autowired
    OrgRoleUserService orgRoleUserService;

	@RequestMapping(value = "list", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> list(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		Integer orgId = userVo.getDefaultOrgId();

		String keywords = json.getString("keywords");
		String disabled = json.getString("disabled");

		List ls = orgRoleService.list(orgId, keywords, disabled);
//		List<TstOrgRole> vos = orgRoleService.genVos(ls);
//
//        ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "get", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject req) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		Integer orgId = userVo.getDefaultOrgId();
		Integer orgRoleId = req.getInteger("id");

		List<TstOrgPrivilegeDefine> orgRolePrivileges = orgRolePrivilegeService.listPrivilegesByOrgRole(orgId, orgRoleId);
        List<TstUser> orgRoleUsers = orgRoleUserService.listUserByOrgRole(orgId, orgRoleId);

		if (orgRoleId == null) {
			ret.put("orgRole", new TstOrgRole());
	        ret.put("orgRolePrivileges", orgRolePrivileges);
			ret.put("orgRoleUsers", orgRoleUsers);
			ret.put("code", Constant.RespCode.SUCCESS.getCode());
			return ret;
		}

//		TstOrgRole po = (TstOrgRole) orgRoleService.get(TstOrgRole.class, orgRoleId);
//		TstOrgRole vo = orgRoleService.genVo(po);
//
//        ret.put("orgRole", vo);
        ret.put("orgRolePrivileges", orgRolePrivileges);
        ret.put("orgRoleUsers", orgRoleUsers);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		Integer orgId = userVo.getDefaultOrgId();

		TstOrgRole orgRoleVo = JSON.parseObject(JSON.toJSONString(json.get("orgRole")), TstOrgRole.class);
		TstOrgRole po = orgRoleService.save(orgRoleVo, orgId);

		List<TstOrgPrivilegeDefine> orgPrivileges = (List<TstOrgPrivilegeDefine>) json.get("orgPrivileges");
		boolean success = orgRolePrivilegeService.saveOrgRolePrivileges(po.getId(), orgPrivileges);

        List<TstUser> orgRoleUsers = (List<TstUser>) json.get("orgRoleUsers");
        success = orgRoleUserService.saveOrgRoleUsers(po.getId(), orgRoleUsers);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "delete", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject to) {
		Map<String, Object> ret = new HashMap<String, Object>();

		boolean success = orgRoleService.delete(to.getInteger("id"));

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
