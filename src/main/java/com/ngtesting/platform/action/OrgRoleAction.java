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
import com.ngtesting.platform.entity.SysOrgRole;
import com.ngtesting.platform.entity.SysRole;
import com.ngtesting.platform.entity.SysUser;
import com.ngtesting.platform.service.OrgPrivilegeService;
import com.ngtesting.platform.service.OrgRoleService;
import com.ngtesting.platform.service.RoleService;
import com.ngtesting.platform.util.AuthPassport;
import com.ngtesting.platform.util.Constant;
import com.ngtesting.platform.util.Constant.RespCode;
import com.ngtesting.platform.vo.OrgGroupVo;
import com.ngtesting.platform.vo.OrgPrivilegeVo;
import com.ngtesting.platform.vo.OrgRoleVo;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.RelationOrgGroupUserVo;
import com.ngtesting.platform.vo.RoleVo;
import com.ngtesting.platform.vo.UserVo;


@Controller
@RequestMapping(Constant.API_PATH_CLIENT + "org_role/")
public class OrgRoleAction extends BaseAction {
	@Autowired
	OrgRoleService orgRoleService;
	@Autowired
	OrgPrivilegeService orgPrivilegeService;
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "list", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> list(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		Long orgId = userVo.getDefaultOrgId();
		
		String keywords = json.getString("keywords");
		String disabled = json.getString("disabled");
		int currentPage = json.getInteger("currentPage") == null? 0: json.getInteger("currentPage") - 1;
		int itemsPerPage = json.getInteger("itemsPerPage") == null? Constant.PAGE_SIZE: json.getInteger("itemsPerPage");
		
		Page page = orgRoleService.listByPage(orgId, keywords, disabled, currentPage, itemsPerPage);
		List<OrgRoleVo> vos = orgRoleService.genVos(page.getItems());
        
		ret.put("totalItems", page.getTotal());
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
		
		List<OrgPrivilegeVo> orgPrivileges = orgPrivilegeService.listPrivilegesByOrg(orgId, orgRoleId);
		if (orgRoleId == null) {
			ret.put("orgRole", new OrgGroupVo());
	        ret.put("orgPrivileges", orgPrivileges);
			ret.put("code", Constant.RespCode.SUCCESS.getCode());
			return ret;
		}
		
		SysOrgRole po = (SysOrgRole) orgRoleService.get(SysOrgRole.class, orgRoleId);
		OrgRoleVo vo = orgRoleService.genVo(po);
        
        ret.put("orgRole", vo);
        ret.put("orgPrivileges", orgPrivileges);
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
		SysOrgRole po = orgRoleService.save(orgRoleVo, orgId);
		
		List<OrgPrivilegeVo> orgPrivileges = (List<OrgPrivilegeVo>) json.get("orgPrivileges");
		boolean success = orgPrivilegeService.saveOrgPrivileges(po.getId(), orgPrivileges);
		
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
