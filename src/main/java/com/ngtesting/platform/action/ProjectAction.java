package com.ngtesting.platform.action;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.entity.TestHistory;
import com.ngtesting.platform.entity.TestPlan;
import com.ngtesting.platform.entity.TestProject;
import com.ngtesting.platform.entity.TestRelationProjectRoleEntity;
import com.ngtesting.platform.service.*;
import com.ngtesting.platform.util.AuthPassport;
import com.ngtesting.platform.vo.*;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
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
@RequestMapping(Constant.API_PATH_CLIENT + "project/")
public class ProjectAction extends BaseAction {
	private static final Log log = LogFactory.getLog(ProjectAction.class);

	@Autowired
    ProjectService projectService;
	@Autowired
	PlanService planService;
	@Autowired
	HistoryService historyService;

    @Autowired
    ProjectRoleService projectRoleService;
    @Autowired
    PushSettingsService pushSettingsService;

    @Autowired
    RelationProjectRoleEntityService relationProjectRoleEntityService;

	@AuthPassport(validate = true)
	@RequestMapping(value = "list", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> list(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		Long orgId = userVo.getDefaultOrgId();

		String keywords = json.getString("keywords");
		String disabled = json.getString("disabled");

		List<TestProjectVo> vos = projectService.listVos(orgId, keywords, disabled);

        ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());

		return ret;
	}

	@AuthPassport(validate = true)
	@RequestMapping(value = "get", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		Long orgId = userVo.getDefaultOrgId();

		Long projectId = json.getLong("id");

		if (projectId != null) {
			TestProject project = projectService.getDetail(projectId);
			TestProjectVo vo = projectService.genVo(project);

			if (TestProject.ProjectType.group.equals(project.getType())) {
				vo.setLastestProjectGroup(projectService.isLastestProjectGroup(orgId, projectId));
			}

			ret.put("data", vo);
		}

		List<TestProjectVo> groups = projectService.listProjectGroups(orgId);
        List<ProjectRoleVo> projectRoles = projectRoleService.list(orgId, null, null);

		List<TestRelationProjectRoleEntity> entityInRolesPos = relationProjectRoleEntityService.listByProject(projectId);
        List<RelationProjectRoleEntityVo> entityInRoles = relationProjectRoleEntityService.genVos(entityInRolesPos);

        ret.put("groups", groups);
        ret.put("projectRoles", projectRoles);
		ret.put("entityInRoles", entityInRoles);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@AuthPassport(validate = true)
	@RequestMapping(value = "view", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> view(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		Long id = json.getLong("id");

		TestProjectVo vo = projectService.viewPers(id, userVo);

        List<TestPlan> planPos = planService.list(id);
        List<TestPlanVo> planVos = planService.genVos(planPos);

        List<TestHistory> historyPos = historyService.list(id);
        Map<String, List<TestHistoryVo>> historyVos = historyService.genVosByDate(historyPos);

        pushSettingsService.pushRecentProjects(userVo);
        pushSettingsService.pushPrjSettings(userVo);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
        ret.put("project", vo);
        ret.put("plans", planVos);
        ret.put("histories", historyVos);

		return ret;
	}

	@AuthPassport(validate = true)
	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		Long orgId = userVo.getDefaultOrgId();
		Long userId = userVo.getId();

		TestProjectVo vo = json.getObject("model", TestProjectVo.class);

        TestProject po = projectService.save(vo, orgId, userVo);
        projectService.updateNameInHisotyPers(po.getId(), userId);
        pushSettingsService.pushRecentProjects(userVo);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@AuthPassport(validate = true)
	@RequestMapping(value = "saveMembers", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> saveMembers(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		Long orgId = userVo.getDefaultOrgId();

        List<TestRelationProjectRoleEntity> pos = relationProjectRoleEntityService.batchSavePers(json);
        List<RelationProjectRoleEntityVo> entityInRoles = relationProjectRoleEntityService.genVos(pos);

		ret.put("entityInRoles", entityInRoles);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@AuthPassport(validate = true)
	@RequestMapping(value = "changeRole", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> changeRole(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		Long orgId = userVo.getDefaultOrgId();

		List<TestRelationProjectRoleEntity> pos = relationProjectRoleEntityService.changeRolePers(json);
		List<RelationProjectRoleEntityVo> entityInRoles = relationProjectRoleEntityService.genVos(pos);

		ret.put("entityInRoles", entityInRoles);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@AuthPassport(validate = true)
	@RequestMapping(value = "delete", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		Long id = json.getLong("id");

		projectService.delete(id);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
