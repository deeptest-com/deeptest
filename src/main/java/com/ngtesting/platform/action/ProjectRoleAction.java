package com.ngtesting.platform.action;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstProjectPrivilegeDefine;
import com.ngtesting.platform.model.TstProjectRole;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.ProjectPrivilegeService;
import com.ngtesting.platform.service.ProjectRoleService;
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
@RequestMapping(Constant.API_PATH_CLIENT + "project_role/")
public class ProjectRoleAction extends BaseAction {
	@Autowired
    ProjectRoleService projectRoleService;
	@Autowired
    ProjectPrivilegeService projectPrivilegeService;


	@RequestMapping(value = "list", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> list(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		Integer orgId = userVo.getDefaultOrgId();

		String keywords = json.getString("keywords");
		Boolean disabled = json.getBoolean("disabled");

		List<TstProjectRole> vos = projectRoleService.list(orgId, keywords, disabled);

        ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}


	@RequestMapping(value = "get", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject req) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		Integer orgId = userVo.getDefaultOrgId();
		Integer roleId = req.getInteger("id");

		Map<String, Map<String, TstProjectPrivilegeDefine>> orgPrivileges =
				projectPrivilegeService.listPrivilegesByOrgAndProjectRole(orgId, roleId);

		TstProjectRole po = null;
		if (roleId == null) {
			po = new TstProjectRole();
		} else {
			po = projectRoleService.get(roleId);
		}

        ret.put("projectRole", po);
        ret.put("projectPrivileges", orgPrivileges);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		Integer orgId = userVo.getDefaultOrgId();

		TstProjectRole projectRoleVo = JSON.parseObject(JSON.toJSONString(json.get("projectRole")), TstProjectRole.class);
		TstProjectRole po = projectRoleService.save(projectRoleVo, orgId);

		Map<String, List<TstProjectPrivilegeDefine>> projectPrivileges =
				(Map<String, List<TstProjectPrivilegeDefine>>) json.get("projectPrivileges");

		boolean success = projectPrivilegeService.saveProjectPrivileges(orgId, po.getId(), projectPrivileges);

		Map<String, Boolean> prjPrivileges = projectPrivilegeService.listByUser(userVo.getId(), userVo.getDefaultPrjId(), orgId);
		ret.put("prjPrivileges", prjPrivileges);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}


	@RequestMapping(value = "delete", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject to) {
		Map<String, Object> ret = new HashMap<String, Object>();

		boolean success = projectRoleService.delete(to.getInteger("id"));

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
