package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.github.pagehelper.PageHelper;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.bean.websocket.WsFacade;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.config.WsConstant;
import com.ngtesting.platform.model.*;
import com.ngtesting.platform.service.intf.TestEnvService;
import com.ngtesting.platform.service.intf.TestPlanService;
import com.ngtesting.platform.service.intf.TestSuiteService;
import com.ngtesting.platform.service.intf.TestVerService;
import com.ngtesting.platform.servlet.PrivPrj;
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
@RequestMapping(Constant.API_PATH_CLIENT + "plan/")
public class PlanAction extends BaseAction {
	@Autowired
	private WsFacade optFacade;

	@Autowired
	TestPlanService planService;
	@Autowired
	TestSuiteService suiteService;
	@Autowired
	TestVerService verService;
	@Autowired
	TestEnvService envService;


	@RequestMapping(value = "query", method = RequestMethod.POST)
	@ResponseBody
    @PrivPrj
	public Map<String, Object> query(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer prjId = user.getDefaultPrjId();

		String keywords = json.getString("keywords");
		String status = json.getString("status");
		Integer pageNum = json.getInteger("page");
		Integer pageSize = json.getInteger("pageSize");

		com.github.pagehelper.Page page = PageHelper.startPage(pageNum, pageSize);
        List<TstPlan> pos = planService.listByPage(prjId, keywords, status);
        planService.genVos(pos);

		ret.put("total", page.getTotal());
        ret.put("data", pos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

    @RequestMapping(value = "get", method = RequestMethod.POST)
    @ResponseBody
    @PrivPrj(perms = {"test_plan-view"})
    public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer prjId = user.getDefaultPrjId();

        Integer id = json.getInteger("id");

		TstPlan vo = planService.getById(id, prjId);
		if (vo == null) {
            return authFail();
        }

		List<TstSuite> suites = suiteService.listForImport(prjId);

		List<TstVer> vers = verService.list(prjId, null, null);
		List<TstEnv> envs = envService.list(prjId, null, null);

        ret.put("data", vo);
		ret.put("suites", suites);
		ret.put("vers", vers);
		ret.put("envs", envs);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
    @PrivPrj(perms = {"test_plan-maintain"})
	public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer prjId = user.getDefaultPrjId();

        TstPlan vo = JSON.parseObject(JSON.toJSONString(json), TstPlan.class);

		TstPlan po = planService.save(vo, user, prjId);
        if (po == null) {
            return authFail();
        }

		planService.genVo(po);

		optFacade.opt(WsConstant.WS_TODO, user);

		ret.put("data", po);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "delete", method = RequestMethod.POST)
	@ResponseBody
    @PrivPrj(perms = {"test_plan-delete"})
	public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer prjId = user.getDefaultPrjId();

		Integer id = json.getInteger("id");

		Boolean result = planService.delete(id, prjId);
        if (!result) {
            return authFail();
        }

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
