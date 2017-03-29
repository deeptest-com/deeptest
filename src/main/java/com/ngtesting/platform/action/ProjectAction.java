package com.ngtesting.platform.action;

import java.util.Date;
import java.util.HashMap;
import java.util.HashSet;
import java.util.LinkedList;
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

import com.ngtesting.platform.entity.EvtBanner;
import com.ngtesting.platform.entity.EvtClient;
import com.ngtesting.platform.entity.EvtDocument;
import com.ngtesting.platform.entity.EvtEvent;
import com.ngtesting.platform.entity.EvtGuest;
import com.ngtesting.platform.entity.EvtOrganizer;
import com.ngtesting.platform.entity.SysUser;
import com.ngtesting.platform.entity.TestCase;
import com.ngtesting.platform.entity.TestProject;
import com.ngtesting.platform.service.GuestService;
import com.ngtesting.platform.service.TestCaseService;
import com.ngtesting.platform.service.TestProjectService;
import com.ngtesting.platform.service.impl.TestProjectServiceImpl;
import com.ngtesting.platform.util.AuthPassport;
import com.ngtesting.platform.util.Constant;
import com.ngtesting.platform.util.StringUtil;
import com.ngtesting.platform.vo.BannerVo;
import com.ngtesting.platform.vo.DocumentVo;
import com.ngtesting.platform.vo.EventVo;
import com.ngtesting.platform.vo.GuestVo;
import com.ngtesting.platform.vo.OrganizerVo;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.TestCaseTreeVo;
import com.ngtesting.platform.vo.TestCaseVo;
import com.ngtesting.platform.vo.TestProjectVo;
import com.ngtesting.platform.vo.UserVo;

import com.alibaba.fastjson.JSONObject;


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
		
		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		
		String keywords = json.getString("keywords");
		String disabled = json.getString("disabled");
		
		Long t1 = new Date().getTime();

		List<TestProjectVo> vos = projectService.list(userVo.getCompanyId(), keywords, disabled);
		
		Long t2 = new Date().getTime();
		log.debug("获取项目信息花了" + (t1 - t2) + "毫秒");
		
        ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		
		return ret;
	}
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "get", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> get(HttpServletRequest request, @RequestBody TestProjectVo json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		Long id = json.getId();
		
		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		
		TestProject project = projectService.getDetail(id);
		TestProjectVo vo = projectService.genVo(project);
		
		List<TestProjectVo> vos = projectService.listGroups(userVo.getCompanyId());
        
        ret.put("data", vo);
        ret.put("groups", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@AuthPassport(validate = true)
	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody TestProjectVo vo) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		
		TestProject po = projectService.save(vo, userVo);
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
}
