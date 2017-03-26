package cn.linkr.testspace.action;

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

import cn.linkr.testspace.entity.EvtBanner;
import cn.linkr.testspace.entity.EvtClient;
import cn.linkr.testspace.entity.EvtDocument;
import cn.linkr.testspace.entity.EvtEvent;
import cn.linkr.testspace.entity.EvtGuest;
import cn.linkr.testspace.entity.EvtOrganizer;
import cn.linkr.testspace.entity.SysUser;
import cn.linkr.testspace.entity.TestCase;
import cn.linkr.testspace.entity.TestProject;
import cn.linkr.testspace.service.GuestService;
import cn.linkr.testspace.service.TestCaseService;
import cn.linkr.testspace.service.TestProjectService;
import cn.linkr.testspace.service.impl.TestProjectServiceImpl;
import cn.linkr.testspace.util.AuthPassport;
import cn.linkr.testspace.util.Constant;
import cn.linkr.testspace.util.StringUtil;
import cn.linkr.testspace.vo.BannerVo;
import cn.linkr.testspace.vo.DocumentVo;
import cn.linkr.testspace.vo.EventVo;
import cn.linkr.testspace.vo.GuestVo;
import cn.linkr.testspace.vo.OrganizerVo;
import cn.linkr.testspace.vo.Page;
import cn.linkr.testspace.vo.TestCaseTreeVo;
import cn.linkr.testspace.vo.TestCaseVo;
import cn.linkr.testspace.vo.TestProjectVo;
import cn.linkr.testspace.vo.UserVo;

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
		
		String isActive = json.getString("isActive");
		
		Long t1 = new Date().getTime();

		Map<String, Object> out = projectService.listCache(userVo.getCompanyId(), isActive);
		
		Long t2 = new Date().getTime();
		log.debug("获取项目信息花了" + (t1 - t2) + "毫秒");
		
        ret.put("data", out);
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
		
		Map<String, Object> out = projectService.listCache(userVo.getCompanyId(), "true");
		LinkedList<TestProjectVo> vos = projectService.removeChildren((LinkedList<TestProjectVo>)out.get("models"), vo);
        
        ret.put("data", vo);
        ret.put("projects", vos);
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
