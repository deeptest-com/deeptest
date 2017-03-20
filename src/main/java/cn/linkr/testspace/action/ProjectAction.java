package cn.linkr.testspace.action;

import java.util.Date;
import java.util.HashMap;
import java.util.HashSet;
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
		String keywords = json.getString("keywords");
		
		Long t1 = new Date().getTime();
		
		List<TestProject> pos = projectService.list(isActive, keywords, userVo.getCompanyId());
		Map<String, Integer> param = new HashMap<String, Integer>();
		TestProjectVo vos = projectService.genVos(pos, param);
		
		Long t2 = new Date().getTime();
		log.debug("获取项目信息花了" + (t1 - t2) + "毫秒");
			
        ret.put("data", vos);
        ret.put("maxLevel", param.get("maxLevel"));
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		
		return ret;
	}
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "get", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> get(HttpServletRequest request) {
		Map<String, Object> ret = new HashMap<String, Object>();
		JSONObject req = reqJson(request);
		String id = req.getString("id");
		
		UserVo user = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		
		TestProject project = projectService.getDetail(Long.valueOf(id));
		TestProjectVo vo = projectService.genVo(project);
        
        ret.put("data", vo);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@AuthPassport(validate = true)
	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		Long id = json.getLong("id");
		String value = json.getString("value");
		Integer type = json.getInteger("type");
		Long pid = json.getLong("pid");
		
		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		
		TestProject po = projectService.save(id, value, type, pid, userVo.getId());
		TestProjectVo caseVo = projectService.genVo(po);
        
        ret.put("data", caseVo);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "delete", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		Long id = json.getLong("id");
		
		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		
		TestProject po = projectService.delete(id, userVo.getId());
		TestProjectVo caseVo = projectService.genVo(po);
        
        ret.put("data", caseVo);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}
}
