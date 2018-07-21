package com.ngtesting.platform.action;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstOrgGroup;
import com.ngtesting.platform.model.TstOrgGroupUserRelation;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.OrgGroupService;
import com.ngtesting.platform.service.OrgGroupUserRelationService;
import com.ngtesting.platform.vo.Page;
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
@RequestMapping(Constant.API_PATH_CLIENT + "org_group/")
public class OrgGroupAction extends BaseAction {
	@Autowired
    OrgGroupService orgGroupService;

	@Autowired
	OrgGroupUserRelationService orgGroupUserService;

	@RequestMapping(value = "list", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> list(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		Integer orgId = userVo.getDefaultOrgId();

		String keywords = json.getString("keywords");
		String disabled = json.getString("disabled");
		int page = json.getInteger("page") == null? 0: json.getInteger("page") - 1;
		int pageSize = json.getInteger("pageSize") == null? Constant.PAGE_SIZE: json.getInteger("pageSize");

		Page pageData = orgGroupService.listByPage(orgId, keywords, disabled, page, pageSize);
		List<TstOrgGroup> vos = orgGroupService.genVos(pageData.getItems());

		ret.put("collectionSize", pageData.getTotal());
        ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "get", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		Integer orgId = userVo.getDefaultOrgId();
		Integer orgGroupId = json.getInteger("id");

		List<TstOrgGroupUserRelation> relations = orgGroupUserService.listRelationsByGroup(orgId, orgGroupId);
		if (orgGroupId == null) {

			ret.put("group", new TstOrgGroup());
	        ret.put("relations", relations);
			ret.put("code", Constant.RespCode.SUCCESS.getCode());
			return ret;
		}

//		TstOrgGroup po = (TstOrgGroup) orgGroupService.get(TstOrgGroup.class, Integer.valueOf(orgGroupId));
//		TstOrgGroup group = orgGroupService.genVo(po);
//
//        ret.put("group", group);
        ret.put("relations", relations);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		Integer orgId = userVo.getDefaultOrgId();

		TstOrgGroup group = JSON.parseObject(JSON.toJSONString(json.get("group")), TstOrgGroup.class);;
		List<TstOrgGroupUserRelation> relations = (List<TstOrgGroupUserRelation>) json.get("relations");

		TstOrgGroup po = orgGroupService.save(group, orgId);
		boolean success = orgGroupUserService.saveRelations(relations);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "delete", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject to) {
		Map<String, Object> ret = new HashMap<String, Object>();

		Integer id = to.getInteger("id");

		boolean success = orgGroupService.delete(id);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
