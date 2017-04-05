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

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.SysUser;
import com.ngtesting.platform.service.RelationOrgGroupUserService;
import com.ngtesting.platform.service.UserService;
import com.ngtesting.platform.util.AuthPassport;
import com.ngtesting.platform.util.Constant;
import com.ngtesting.platform.util.Constant.RespCode;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.RelationOrgGroupUserVo;
import com.ngtesting.platform.vo.UserVo;


@Controller
@RequestMapping(Constant.API_PATH_CLIENT + "user/")
public class UserAction extends BaseAction {
	@Autowired
	UserService userService;
	@Autowired
	RelationOrgGroupUserService orgGroupUserService;
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "list", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> list(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		Long orgId = userVo.getDefaultOrgId();
		
		String keywords = json.getString("keywords");
		String disabled = json.getString("disabled");
		int currentPage = json.getInteger("currentPage") == null? 0: json.getInteger("currentPage") - 1;
		int itemsPerPage = json.getInteger("itemsPerPage") == null? Constant.PAGE_SIZE: json.getInteger("itemsPerPage");
		
		Page page = userService.listByPage(orgId, keywords, disabled, currentPage, itemsPerPage);
		List<UserVo> vos = userService.genVos(page.getItems());
        
		ret.put("totalItems", page.getTotal());
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
		Long userId = json.getLong("id");
		
		List<RelationOrgGroupUserVo> relations = orgGroupUserService.listRelationsByUser(orgId, userId);
		
		if (userId == null) {
			ret.put("user", new UserVo());
	        ret.put("relations", relations);
			ret.put("code", Constant.RespCode.SUCCESS.getCode());
			return ret;
		}
		
		SysUser po = (SysUser) userService.get(SysUser.class, Long.valueOf(userId));
		UserVo vo = userService.genVo(po);
		
        ret.put("user", vo);
        ret.put("relations", relations);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		Long orgId = userVo.getDefaultOrgId();
		
		UserVo user = JSON.parseObject(JSON.toJSONString(json.get("user")), UserVo.class);
		SysUser po = userService.save(user, orgId);
		
		if (po == null) {
			ret.put("code", RespCode.BIZ_FAIL.getCode());
			ret.put("msg", "邮箱已存在");
			return ret;
		} 

		List<RelationOrgGroupUserVo> relations = (List<RelationOrgGroupUserVo>) json.get("relations");
		boolean success = orgGroupUserService.saveRelations(relations);
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
		
		boolean success = userService.disable(json.getLong("id"), orgId);
		
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
		
		boolean success = userService.remove(userId, orgId);
		
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}
	
}
