package cn.linkr.testspace.action;

import java.util.HashMap;
import java.util.Map;

import javax.servlet.http.HttpServletRequest;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.ResponseBody;


import cn.linkr.testspace.entity.EvtClient;
import cn.linkr.testspace.entity.SysCompany;
import cn.linkr.testspace.entity.SysUser;
import cn.linkr.testspace.service.CompanyService;
import cn.linkr.testspace.service.UserService;
import cn.linkr.testspace.util.AuthPassport;
import cn.linkr.testspace.util.Constant;
import cn.linkr.testspace.vo.CompanyVo;
import cn.linkr.testspace.vo.DocumentVo;
import cn.linkr.testspace.vo.UserVo;

import com.alibaba.fastjson.JSONObject;


@Controller
@RequestMapping(Constant.API_PATH_CLIENT + "company/")
public class CompanyAction extends BaseAction {
	@Autowired
	UserService userService;
	
	@Autowired
	CompanyService companyService;
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "get", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject req) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		UserVo vo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		long companyId = vo.getCompanyId();
		
		SysCompany po = (SysCompany) companyService.get(SysCompany.class, companyId);
		CompanyVo eventVo = companyService.genVo(po);
        
        ret.put("data", eventVo);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody CompanyVo vo) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		SysCompany doc = companyService.save(vo);
		
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
