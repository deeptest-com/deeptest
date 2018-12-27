package com.ngtesting.platform.action.admin;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstOrg;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.OrgService;
import com.ngtesting.platform.service.intf.PushSettingsService;
import com.ngtesting.platform.servlet.PrivCommon;
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
@RequestMapping(Constant.API_PATH_ADMIN + "org/")
public class OrgAdmin extends BaseAction {
	@Autowired
	OrgService orgService;

	@Autowired
	PushSettingsService pushSettingsService;


	@RequestMapping(value = "list", method = RequestMethod.POST)
	@ResponseBody
	@PrivCommon(check="false")
	public Map<String, Object> list(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

		String keywords = json.getString("keywords");
		Boolean disabled = json.getBoolean("disabled");

		List<TstOrg> vos = orgService.list(user.getId(), keywords, disabled); // 总是取当前用户的org，不需要再鉴权

        ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "get", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = json.getInteger("orgId");

		TstOrg po = orgService.get(orgId); //

		ret.put("data", po);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
    @PrivCommon(check="false")
	public Map<String, Object> save(HttpServletRequest request, @RequestBody TstOrg vo) {
		Map<String, Object> ret = new HashMap<String, Object>();
		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer orgId = vo.getId();

        TstOrg org = orgService.save(vo, user);

        if (user.getDefaultOrgId() == null) {
            orgService.changeDefaultOrg(user, org.getId());
        } else if (user.getDefaultOrgId().intValue() == org.getId().intValue() &&
                !org.getName().equals(user.getDefaultOrgName())) { // 修改当前组织名称
            user.setDefaultOrgName(vo.getName());
            pushSettingsService.pushOrgSettings(user);
        }

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

    @RequestMapping(value = "update", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> update(HttpServletRequest request, @RequestBody TstOrg vo) {
        Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer orgId = vo.getId();

        TstOrg org = orgService.update(vo, user);

        if (user.getDefaultOrgId() == null) {
            orgService.changeDefaultOrg(user, org.getId());
        } else if (user.getDefaultOrgId().intValue() == org.getId().intValue() &&
                !org.getName().equals(user.getDefaultOrgName())) { // 修改当前组织名称
            user.setDefaultOrgName(vo.getName());
            pushSettingsService.pushOrgSettings(user);
        }

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

	@RequestMapping(value = "delete", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = json.getInteger("orgId");

		Boolean result = orgService.delete(orgId, user);

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

    // 用户在列表页，设置默认组织
	@RequestMapping(value = "setDefault", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> setDefault(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

		Integer orgId = json.getInteger("orgId");
		String keywords = json.getString("keywords");
		Boolean disabled = json.getBoolean("disabled");

		orgService.changeDefaultOrg(user, orgId);  // 涵盖项目设置WS推送消息

        List<TstOrg> vos = orgService.list(user.getId(), keywords, disabled);

        ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());

		return ret;
	}

}
