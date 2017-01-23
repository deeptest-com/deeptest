package cn.linkr.events.action.admin;

import java.util.HashMap;
import java.util.Map;

import javax.servlet.http.HttpServletRequest;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.ResponseBody;

import cn.linkr.events.action.client.BaseAction;
import cn.linkr.events.constants.Constant;
import cn.linkr.events.entity.EvtClient;
import cn.linkr.events.entity.SysCompany;
import cn.linkr.events.entity.SysUser;
import cn.linkr.events.service.CompanyService;
import cn.linkr.events.service.UserService;
import cn.linkr.events.util.AuthPassport;
import cn.linkr.events.vo.CompanyVo;
import cn.linkr.events.vo.DocumentVo;
import cn.linkr.events.vo.UserVo;

import com.alibaba.fastjson.JSONObject;


@Controller
@RequestMapping(Constant.API_PATH_ADMIN + "company/")
public class CompanyAdmin extends BaseAction {
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
