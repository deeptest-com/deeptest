package com.ngtesting.platform.action;

import java.util.Date;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import javax.servlet.http.HttpServletRequest;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.ResponseBody;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.TestProject;
import com.ngtesting.platform.entity.TestProjectAccessHistory;
import com.ngtesting.platform.service.TestProjectService;
import com.ngtesting.platform.util.AuthPassport;
import com.ngtesting.platform.util.Constant;
import com.ngtesting.platform.vo.TestProjectAccessHistoryVo;
import com.ngtesting.platform.vo.TestProjectVo;
import com.ngtesting.platform.vo.UserVo;


@Controller
@RequestMapping(Constant.API_PATH_CLIENT + "project/")
public class ProjectAction extends BaseAction {
	private static final Log log = LogFactory.getLog(ProjectAction.class);
	
	@Autowired
	TestProjectService projectService;
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "list", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> list(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		UserVo userVo = genRequest(request, json);
		Long orgId = json.getLong("orgId");
		String keywords = json.getString("keywords");
		String disabled = json.getString("disabled");
		
		Long t1 = new Date().getTime();

		List<TestProjectVo> vos = projectService.list(orgId, keywords, disabled);
		
		Long t2 = new Date().getTime();
		log.debug("获取项目信息花了" + (t1 - t2) + "毫秒");
		
        ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		
		return ret;
	}
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "get", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		Long orgId = json.getLong("orgId");
		TestProjectVo vo = json.getObject("vo", TestProjectVo.class);
		Long id = vo.getId();
		
		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		
		TestProject project = projectService.getDetail(id);
		TestProjectVo vo2 = projectService.genVo(project);
		
		List<TestProjectVo> vos = projectService.listProjectGroups(orgId);
        
        ret.put("data", vo2);
        ret.put("groups", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@AuthPassport(validate = true)
	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		Long orgId = json.getLong("orgId");
		TestProjectVo vo = json.getObject("vo", TestProjectVo.class);
		Long id = vo.getId();
		
		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		
		TestProject po = projectService.save(vo, orgId);
		TestProjectVo projectVo = projectService.genVo(po);
        
        ret.put("data", projectVo);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "delete", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> delete(HttpServletRequest request, @RequestBody TestProjectVo json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		Long id = json.getId();
		
		projectService.delete(id, userVo.getId());
        
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "view", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> view(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		UserVo userVo = genRequest(request, json);
		Long orgId = json.getLong("orgId");
		Long projectId = json.getLong("projectId");
		
		TestProjectVo vo = projectService.viewPers(orgId, userVo, projectId);
		
		List<TestProjectAccessHistoryVo> recentProjects = projectService.listRecentProjectVo(orgId, userVo.getId());
		
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		ret.put("project", vo);
		ret.put("recentProjects", recentProjects);
		return ret;
	}
}
