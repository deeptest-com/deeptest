package com.ngtesting.platform.action.admin;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstProjectPrivilegeDefine;
import com.ngtesting.platform.model.TstProjectRole;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.ProjectPrivilegeService;
import com.ngtesting.platform.service.intf.ProjectRoleService;
import org.apache.shiro.SecurityUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

import javax.servlet.http.HttpServletRequest;
import java.util.HashMap;
import java.util.List;
import java.util.Map;


@RestController
@RequestMapping(Constant.API_PATH_ADMIN + "project_role/")
public class ProjectRoleAdmin extends BaseAction {
	@Autowired
    ProjectRoleService projectRoleService;
	@Autowired
    ProjectPrivilegeService projectPrivilegeService;


	@RequestMapping(value = "list", method = RequestMethod.POST)

	public Map<String, Object> list(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();
		Integer orgId = user.getDefaultOrgId();

		String keywords = json.getString("keywords");
		Boolean disabled = json.getBoolean("disabled");

		List<TstProjectRole> vos = projectRoleService.list(orgId, keywords, disabled); // 总是取当前用户的org，不需要再鉴权

        ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "get", method = RequestMethod.POST)

	public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();
		Integer orgId = user.getDefaultOrgId();
		Integer roleId = json.getInteger("id");

		TstProjectRole po = null;
		if (roleId == null) {
			po = new TstProjectRole();
		} else {
			po = projectRoleService.get(roleId, orgId);
			if (po == null) { // 当对象不是默认org的，此处为空
                return authorFail();
            }
		}

        Map<String, Map<String, TstProjectPrivilegeDefine>> projectPrivileges =
                projectPrivilegeService.listPrivilegesByOrgAndProjectRole(orgId, roleId);

        ret.put("projectRole", po);
        ret.put("projectPrivileges", projectPrivileges);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "save", method = RequestMethod.POST)

	public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();
		Integer orgId = user.getDefaultOrgId();

		TstProjectRole vo = JSON.parseObject(JSON.toJSONString(json.get("projectRole")), TstProjectRole.class);
		TstProjectRole po = projectRoleService.save(vo, orgId);
        if (po == null) { // 当对象不是默认org的，update的结果会返回空
            return authorFail();
        }

		Map<String, List<TstProjectPrivilegeDefine>> projectPrivileges =
				(Map<String, List<TstProjectPrivilegeDefine>>) json.get("projectPrivileges");

		projectPrivilegeService.saveProjectPrivileges(orgId, po.getId(), projectPrivileges);

		Map<String, Boolean> prjPrivileges = projectPrivilegeService.listByUser(user.getId(), user.getDefaultPrjId(), orgId);
		ret.put("prjPrivileges", prjPrivileges);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}


	@RequestMapping(value = "delete", method = RequestMethod.POST)

	public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();
		Integer orgId = user.getDefaultOrgId();

		Integer roleId = json.getInteger("id");

		Boolean result = projectRoleService.delete(roleId, orgId);
		if (!result) { // 当对象不是默认org的，结果会返回false
            return authorFail();
        }

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
