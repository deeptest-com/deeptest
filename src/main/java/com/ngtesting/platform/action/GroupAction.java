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
import com.ngtesting.platform.entity.SysGroup;
import com.ngtesting.platform.entity.SysUser;
import com.ngtesting.platform.service.GroupService;
import com.ngtesting.platform.service.UserService;
import com.ngtesting.platform.util.AuthPassport;
import com.ngtesting.platform.util.Constant;
import com.ngtesting.platform.vo.GroupVo;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.UserVo;


@Controller
@RequestMapping(Constant.API_PATH_CLIENT + "group/")
public class GroupAction extends BaseAction {
	@Autowired
	UserService userService;
	
	@Autowired
	GroupService groupService;
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "list", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> list(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		UserVo vo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		long companyId = vo.getCompanyId();
		
		String keywords = json.getString("keywords");
		String disabled = json.getString("disabled");
		int currentPage = json.getInteger("currentPage") == null? 0: json.getInteger("currentPage") - 1;
		int itemsPerPage = json.getInteger("itemsPerPage") == null? Constant.PAGE_SIZE: json.getInteger("itemsPerPage");
		
		Page page = groupService.listByPage(companyId, keywords, disabled, currentPage, itemsPerPage);
		List<GroupVo> vos = groupService.genVos(page.getItems());
        
		ret.put("totalItems", page.getTotal());
        ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "get", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject req) {
		Map<String, Object> ret = new HashMap<String, Object>();

		UserVo user = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		
		Long groupId = req.getLong("id");
		if (groupId == null) {
			List<UserVo> users = userService.listByGroup(user.getCompanyId(), groupId);
			ret.put("group", new SysGroup());
	        ret.put("users", users);
			ret.put("code", Constant.RespCode.SUCCESS.getCode());
			return ret;
		}
		
		SysGroup po = (SysGroup) groupService.get(SysGroup.class, Long.valueOf(groupId));
		GroupVo group = groupService.genVo(po);
		
		List<UserVo> users = userService.listByGroup(user.getCompanyId(), Long.valueOf(groupId));
        
        ret.put("group", group);
        ret.put("users", users);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		
		GroupVo group = JSON.parseObject(JSON.toJSONString(json.get("group")), GroupVo.class);;
		List<UserVo> users = (List<UserVo>) json.get("users");
		
		SysGroup po = groupService.save(group, userVo.getCompanyId());
		boolean success = userService.saveUsersByGroup(users, group.getCompanyId(), group.getId());
		
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "delete", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject to) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		boolean success = groupService.delete(to.getLong("id"));
		
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@AuthPassport(validate = true)
	@RequestMapping(value = "disable", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> disable(HttpServletRequest request, @RequestBody JSONObject to) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		boolean success = groupService.disable(to.getLong("id"));
		
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
