package com.ngtesting.platform.action.admin;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.github.pagehelper.Page;
import com.github.pagehelper.PageHelper;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstOrgGroup;
import com.ngtesting.platform.model.TstOrgGroupUserRelation;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.OrgGroupService;
import com.ngtesting.platform.service.OrgGroupUserRelationService;
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
@RequestMapping(Constant.API_PATH_ADMIN + "org_group/")
public class OrgGroupAdmin extends BaseAction {
	@Autowired
    OrgGroupService orgGroupService;

	@Autowired
	OrgGroupUserRelationService orgGroupUserService;

	@RequestMapping(value = "list", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> list(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = userVo.getDefaultOrgId();

		String keywords = json.getString("keywords");
		Boolean disabled = json.getBoolean("disabled");
		Integer pageNum = json.getInteger("page");
		Integer pageSize = json.getInteger("pageSize");

		Page page = PageHelper.startPage(pageNum, pageSize);
		List<TstOrgGroup> groups = // 总是取当前用户的org，不需要再鉴权
				orgGroupService.listByPage(orgId, keywords, disabled, pageNum, pageSize);

		ret.put("total", page.getTotal());
        ret.put("data", groups);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "get", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = userVo.getDefaultOrgId();
		Integer orgGroupId = json.getInteger("id");

		TstOrgGroup group;
		if (orgGroupId == null) {
			group = new TstOrgGroup();
            group.setOrgId(orgId);
		} else {
			group = orgGroupService.get(orgGroupId, orgId);
		}
		if (group == null) { // 当对象不是默认org的，此处为空
            return authFail();
        }

		List<TstOrgGroupUserRelation> relations = orgGroupUserService.listRelationsByGroup(orgId, orgGroupId);
		if (orgGroupId == null) {
			ret.put("group", new TstOrgGroup());
	        ret.put("relations", relations);
			ret.put("code", Constant.RespCode.SUCCESS.getCode());
			return ret;
		}

        ret.put("group", group);
        ret.put("relations", relations);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = user.getDefaultOrgId();

		TstOrgGroup group = JSON.parseObject(JSON.toJSONString(json.get("group")), TstOrgGroup.class);;
		List<TstOrgGroupUserRelation> relations = (List<TstOrgGroupUserRelation>) json.get("relations");

		TstOrgGroup po = orgGroupService.save(group, orgId);
        if (po == null) { // 当对象不是默认org的，update的结果会返回空
            return authFail();
        }

		orgGroupUserService.saveRelationsForGroup(orgId, po.getId(), relations);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "delete", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = user.getDefaultOrgId();

		Integer groupId = json.getInteger("id");

		Boolean result = orgGroupService.delete(groupId, orgId);
        if (!result) { // 当对象不是默认org的，结果会返回false
            return authFail();
        }

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
