package com.ngtesting.platform.action;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.*;
import com.ngtesting.platform.service.OrgRolePrivilegeRelationService;
import com.ngtesting.platform.service.OrgRoleService;
import com.ngtesting.platform.service.OrgRoleUserRelationService;
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
    OrgRolePrivilegeRelationService orgRolePrivilegeRelationService;
    @Autowired
    OrgRoleUserRelationService orgRoleUserRelationService;

	@RequestMapping(value = "list", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> list(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		Integer orgId = userVo.getDefaultOrgId();

		String keywords = json.getString("keywords");
		Boolean disabled = json.getBoolean("disabled");

		List<TstOrgRole> ls = orgRoleService.list(orgId, keywords, disabled);

        ret.put("data", ls);
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

        TstOrgRole po = orgRoleService.get(orgRoleId);

		List<TstOrgRolePrivilegeRelation> privileges =
                orgRolePrivilegeRelationService.listRelationsByOrgRole(orgId, orgRoleId);
        List<TstOrgRoleUserRelation> users =
                orgRoleUserRelationService.listRelationsByOrgRole(orgId, orgRoleId);

		if (orgRoleId == null) {
			ret.put("orgRole", new TstOrgRole());
		} else {
            ret.put("orgRole", po);
        }

        ret.put("privileges", privileges);
        ret.put("users", users);
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

		List<TstOrgRolePrivilegeRelation> privileges = (List<TstOrgRolePrivilegeRelation>) json.get("privileges");
		boolean success = orgRolePrivilegeRelationService.saveRelationsForRole(orgId, po.getId(), privileges);

        List<TstOrgRoleUserRelation> users = (List<TstOrgRoleUserRelation>) json.get("users");
        success = orgRoleUserRelationService.saveRelationsForRole(orgId, po.getId(), users);

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
