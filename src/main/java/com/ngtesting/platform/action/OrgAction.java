package com.ngtesting.platform.action;

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

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.SysOrg;
import com.ngtesting.platform.service.OrgService;
import com.ngtesting.platform.util.AuthPassport;
import com.ngtesting.platform.util.Constant;
import com.ngtesting.platform.vo.OrgVo;
import com.ngtesting.platform.vo.UserVo;


@Controller
@RequestMapping(Constant.API_PATH_CLIENT + "org/")
public class OrgAction extends BaseAction {
	@Autowired
	OrgService orgService;
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "list", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> list(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		
		String keywords = json.getString("keywords");
		Boolean disabled = json.getBoolean("disabled");
		
		List<SysOrg> ls = orgService.list(keywords, disabled, userVo.getId());
		List<OrgVo> vos = orgService.genVos(ls);
        
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
		Long orgGroupId = json.getLong("orgGroupId");
		Long userId = json.getLong("userId");
		
		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		
		SysOrg po = (SysOrg) orgService.get(SysOrg.class, Long.valueOf(orgId));
		OrgVo vo = orgService.genVo(po);
		
        ret.put("data", vo);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody OrgVo vo) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		
		SysOrg po = orgService.save(vo, userVo.getId());
		
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "disable", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> disable(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		Long userId = json.getLong("id");
		Long orgId = json.getLong("orgId");
		
		boolean success = orgService.disable(orgId);
		
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "delete", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		Long userId = json.getLong("id");
		Long orgId = json.getLong("orgId");
		
		boolean success = orgService.delete(orgId);
		
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}
	
}
