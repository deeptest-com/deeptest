package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstHistory;
import com.ngtesting.platform.model.TstOrg;
import com.ngtesting.platform.model.TstPlan;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.OrgService;
import com.ngtesting.platform.service.intf.ProjectHistoryService;
import com.ngtesting.platform.service.intf.TestPlanService;
import com.ngtesting.platform.servlet.PrivOrg;
import org.apache.shiro.SecurityUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import javax.servlet.http.HttpServletRequest;
import java.util.HashMap;
import java.util.List;
import java.util.Map;


@RestController
@RequestMapping(Constant.API_PATH_CLIENT + "org/")
public class OrgAction extends BaseAction {
	@Autowired
    OrgService orgService;

	@Autowired
    TestPlanService planService;
	@Autowired
    ProjectHistoryService historyService;

	@RequestMapping(value = "view", method = RequestMethod.POST)
	@PrivOrg(perms = {"belongs_to:org"})
	public Map<String, Object> view(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();
		Integer orgId = json.getInteger("orgId");

		TstOrg po = orgService.get(orgId, user);

		List<TstPlan> planPos = planService.listByOrg(orgId);
		planService.genVos(planPos);

		List<TstHistory> historyPos = historyService.listByOrg(orgId);
		Map<String, List<TstHistory>> historyVos = historyService.genVosByDate(historyPos);

		ret.put("org", po);
		ret.put("plans", planPos);
		ret.put("histories", historyVos);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	// 来源于前端上下文的变化
	@RequestMapping(value = "changeContext", method = RequestMethod.POST)
	@PrivOrg(perms = {"belongs_to:org"})
	public Map<String, Object> changeContext(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();
		Integer orgId = json.getInteger("orgId");

		orgService.changeDefaultOrg(user, orgId); // 涵盖项目设置WS推送消息

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		ret.put("projectId", user.getDefaultPrjId());

		return ret;
	}

}
