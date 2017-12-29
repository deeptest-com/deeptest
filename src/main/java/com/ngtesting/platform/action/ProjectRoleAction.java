package com.ngtesting.platform.action;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.entity.TestProjectRoleForOrg;
import com.ngtesting.platform.service.ProjectPrivilegeService;
import com.ngtesting.platform.service.ProjectRoleService;
import com.ngtesting.platform.util.AuthPassport;
import com.ngtesting.platform.vo.ProjectPrivilegeDefineVo;
import com.ngtesting.platform.vo.ProjectRoleVo;
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
@RequestMapping(Constant.API_PATH_CLIENT + "project_role/")
public class ProjectRoleAction extends BaseAction {
	@Autowired
	ProjectRoleService projectRoleService;
	@Autowired
	ProjectPrivilegeService projectPrivilegeService;

	@AuthPassport(validate = true)
	@RequestMapping(value = "list", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> list(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		Long orgId = userVo.getDefaultOrgId();

		String keywords = json.getString("keywords");
		String disabled = json.getString("disabled");

		List<ProjectRoleVo> vos = projectRoleService.list(orgId, keywords, disabled);

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
		Long roleId = req.getLong("id");

		Map<String, List<ProjectPrivilegeDefineVo>> orgPrivileges =
				projectPrivilegeService.listPrivilegesByOrgAndProjectRole(orgId, roleId);
		if (roleId == null) {
			ret.put("projectRole", new ProjectRoleVo());
	        ret.put("projectPrivileges", orgPrivileges);
			ret.put("code", Constant.RespCode.SUCCESS.getCode());
			return ret;
		}

		TestProjectRoleForOrg po = (TestProjectRoleForOrg) projectRoleService.get(TestProjectRoleForOrg.class, roleId);
		ProjectRoleVo vo = projectRoleService.genVo(po);

        ret.put("projectRole", vo);
        ret.put("projectPrivileges", orgPrivileges);
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

		ProjectRoleVo projectRoleVo = JSON.parseObject(JSON.toJSONString(json.get("projectRole")), ProjectRoleVo.class);
		TestProjectRoleForOrg po = projectRoleService.save(projectRoleVo, orgId);

		Map<String, List<ProjectPrivilegeDefineVo>> projectPrivileges = (Map<String, List<ProjectPrivilegeDefineVo>>) json.get("projectPrivileges");
		boolean success = projectPrivilegeService.saveProjectPrivileges(po.getId(), projectPrivileges);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@AuthPassport(validate = true)
	@RequestMapping(value = "delete", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject to) {
		Map<String, Object> ret = new HashMap<String, Object>();

		boolean success = projectRoleService.delete(to.getLong("id"));

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
