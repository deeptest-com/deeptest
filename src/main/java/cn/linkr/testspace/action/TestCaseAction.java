package cn.linkr.testspace.action;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

import javax.servlet.http.HttpServletRequest;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.ResponseBody;

import cn.linkr.testspace.entity.EvtGuest;
import cn.linkr.testspace.entity.SysUser;
import cn.linkr.testspace.entity.TestCase;
import cn.linkr.testspace.service.GuestService;
import cn.linkr.testspace.service.TestCaseService;
import cn.linkr.testspace.util.AuthPassport;
import cn.linkr.testspace.util.Constant;
import cn.linkr.testspace.vo.GuestVo;
import cn.linkr.testspace.vo.Page;
import cn.linkr.testspace.vo.TestCaseTreeVo;
import cn.linkr.testspace.vo.TestCaseVo;

import com.alibaba.fastjson.JSONObject;


@Controller
@RequestMapping(Constant.API_PATH_ADMIN + "testcase/")
public class TestCaseAction extends BaseAction {
	@Autowired
	TestCaseService caseService;
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "query", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> query(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		Long projectId = json.getLong("projectId");
		Long moduleId = json.getLong("moduleId");
		String keywords = json.getString("keywords");
		
		List<TestCase> ls = caseService.query(projectId, moduleId, keywords);
		
		TestCaseTreeVo tree = caseService.buildTree(ls);
		
        ret.put("data", tree);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}


}
