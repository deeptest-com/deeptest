package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.IsuPage;
import com.ngtesting.platform.model.IsuPageSolution;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.IssuePageService;
import com.ngtesting.platform.service.intf.IssuePageSolutionService;
import com.ngtesting.platform.servlet.PrivOrg;
import com.ngtesting.platform.servlet.PrivPrj;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.apache.shiro.SecurityUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

import javax.servlet.http.HttpServletRequest;
import java.util.HashMap;
import java.util.List;
import java.util.Map;


@RestController
@RequestMapping(Constant.API_PATH_CLIENT + "issue_page/")
public class IssuePageAction extends BaseAction {
	private static final Log log = LogFactory.getLog(IssuePageAction.class);

	@Autowired
	IssuePageSolutionService solutionService;

    @Autowired
    IssuePageService pageService;

	@RequestMapping(value = "get", method = RequestMethod.POST)

	@PrivOrg(perms = {"org_org:*"})
	public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();
        Integer orgId = user.getDefaultOrgId();

        Integer pageId = json.getInteger("id");
        IsuPage page = null;
        if (pageId == null) {
            page = new IsuPage();
        } else {
            page = pageService.get(pageId, orgId);
        }

        ret.put("data", page);

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
	}

	@RequestMapping(value = "getByProject", method = RequestMethod.POST)

	@PrivPrj
	public Map<String, Object> getByProject(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();
		Integer orgId = user.getDefaultOrgId();
		Integer prjId = user.getDefaultPrjId();

		IsuPageSolution solution = solutionService.getByProject(prjId, orgId);
		List<IsuPageSolution> solutions = solutionService.list(orgId);

		Map itemMap = solutionService.getItemsMap(solution.getId(), orgId);

		ret.put("model", solution);
		ret.put("models", solutions);
		ret.put("itemMap", itemMap);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "setByProject", method = RequestMethod.POST)

	@PrivOrg(perms = {"org_org:*"})
	public Map<String, Object> setByProject(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();
		Integer orgId = user.getDefaultOrgId();
		Integer prjId = user.getDefaultPrjId();

		Integer solutionId = json.getInteger("solutionId");

		solutionService.setByProject(solutionId, prjId, orgId);

		IsuPageSolution solution = solutionService.getByProject(prjId, orgId);
		List<IsuPageSolution> solutions = solutionService.list(orgId);

		ret.put("model", solution);
		ret.put("models", solutions);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
