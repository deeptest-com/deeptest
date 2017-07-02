package com.ngtesting.platform.action;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.TestSuite;
import com.ngtesting.platform.service.SuiteService;
import com.ngtesting.platform.util.AuthPassport;
import com.ngtesting.platform.util.Constant;
import com.ngtesting.platform.vo.TestSuiteTreeVo;
import com.ngtesting.platform.vo.TestSuiteVo;
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
@RequestMapping(Constant.API_PATH_CLIENT + "suite/")
public class SuiteAction extends BaseAction {
	@Autowired
	SuiteService suiteService;
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "query", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> query(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		Long projectId = json.getLong("projectId");
		
		List<TestSuite> ls = suiteService.query(projectId);
		
		TestSuiteTreeVo tree = suiteService.buildTree(ls);

        ret.put("data", tree);
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
        Long suiteId = json.getLong("id");

        TestSuiteVo vo = suiteService.getById(suiteId);

        ret.put("data", vo);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

	@AuthPassport(validate = true)
	@RequestMapping(value = "create", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> create(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		Long id = json.getLong("id");
		String value = json.getString("value");
		String type = json.getString("type");
		Long pid = json.getLong("pid");
		
		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		
		TestSuite po = suiteService.create(id, value, type, pid, userVo.getId());
		TestSuiteVo caseVo = suiteService.genVo(po);
        
        ret.put("data", caseVo);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "move", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> move(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		Long id = json.getLong("id");
		Long pid = json.getLong("pid");
		Long prePid = json.getLong("prePid");
		
		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		
		TestSuite po = suiteService.move(id, pid, prePid, userVo.getId());
		TestSuiteVo suiteVo = po == null? null: suiteService.genVo(po);
        
        ret.put("data", suiteVo);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "rename", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> rename(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		Long id = json.getLong("id");
		String value = json.getString("value");
		
		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		
		TestSuite po = suiteService.rename(id, value, userVo.getId());
		TestSuiteVo caseVo = suiteService.genVo(po);
        
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
		
		TestSuite po = suiteService.delete(id, userVo.getId());
		TestSuiteVo caseVo = suiteService.genVo(po);
        
        ret.put("data", caseVo);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@AuthPassport(validate = true)
	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TestSuite po = suiteService.save(json);
		TestSuiteVo caseVo = suiteService.genVo(po);

		ret.put("data", caseVo);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}
}
