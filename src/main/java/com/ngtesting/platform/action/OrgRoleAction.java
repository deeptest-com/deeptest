package com.ngtesting.platform.action;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.TestOrgRole;
import com.ngtesting.platform.service.OrgRolePrivilegeService;
import com.ngtesting.platform.service.OrgRoleService;
import com.ngtesting.platform.service.OrgRoleUserService;
import com.ngtesting.platform.util.AuthPassport;
import com.ngtesting.platform.util.Constant;
import com.ngtesting.platform.vo.*;
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
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "list", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> list(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		Long orgId = userVo.getDefaultOrgId();
		
		String keywords = json.getString("keywords");
		String disabled = json.getString("disabled");
		
		List ls = orgRoleService.list(orgId, keywords, disabled);
		List<OrgRoleVo> vos = orgRoleService.genVos(ls);

        ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "get", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject req) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		Long orgId = userVo.getDefaultOrgId();
		Long orgRoleId = req.getLong("id");
		
		List<OrgPrivilegeVo> orgRolePrivileges = orgRolePrivilegeService.listPrivilegesByOrgRole(orgId, orgRoleId);
        List<UserVo> orgRoleUsers = orgRoleUserService.listUserByOrgRole(orgRoleId);

		if (orgRoleId == null) {
			ret.put("orgRole", new OrgGroupVo());
	        ret.put("orgRolePrivileges", orgRolePrivileges);
			ret.put("orgRoleUsers", orgRoleUsers);
			ret.put("code", Constant.RespCode.SUCCESS.getCode());
			return ret;
		}
		
		TestOrgRole po = (TestOrgRole) orgRoleService.get(TestOrgRole.class, orgRoleId);
		OrgRoleVo vo = orgRoleService.genVo(po);
        
        ret.put("orgRole", vo);
        ret.put("orgRolePrivileges", orgRolePrivileges);
        ret.put("orgRoleUsers", orgRoleUsers);
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
		
		OrgRoleVo orgRoleVo = JSON.parseObject(JSON.toJSONString(json.get("orgRole")), OrgRoleVo.class);
		TestOrgRole po = orgRoleService.save(orgRoleVo, orgId);
		
		List<OrgPrivilegeVo> orgPrivileges = (List<OrgPrivilegeVo>) json.get("orgPrivileges");
		boolean success = orgRolePrivilegeService.saveOrgRolePrivileges(po.getId(), orgPrivileges);

        List<UserVo> orgRoleUsers = (List<UserVo>) json.get("orgRoleUsers");
        success = orgRoleUserService.saveOrgRoleUsers(po.getId(), orgRoleUsers);
		
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "delete", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject to) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		boolean success = orgRoleService.delete(to.getLong("id"));
		
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
