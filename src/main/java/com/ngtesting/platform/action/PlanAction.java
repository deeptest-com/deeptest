package com.ngtesting.platform.action;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.bean.websocket.OptFacade;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.config.WsConstant;
import com.ngtesting.platform.entity.TestPlan;
import com.ngtesting.platform.entity.TestSuite;
import com.ngtesting.platform.service.EnvService;
import com.ngtesting.platform.service.PlanService;
import com.ngtesting.platform.service.SuiteService;
import com.ngtesting.platform.service.VerService;
import com.ngtesting.platform.util.AuthPassport;
import com.ngtesting.platform.vo.*;
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
	private OptFacade optFacade;

	@Autowired
	PlanService planService;
	@Autowired
	SuiteService suiteService;
	@Autowired
	VerService verService;
	@Autowired
	EnvService envService;

	@AuthPassport(validate = true)
	@RequestMapping(value = "query", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> query(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

        int page = json.getInteger("page") == null? 0: json.getInteger("page") - 1;
        int pageSize = json.getInteger("pageSize") == null? Constant.PAGE_SIZE: json.getInteger("pageSize");

        Long projectId = json.getLong("projectId");
        String status = json.getString("status");
        String keywords = json.getString("keywords");

        Page pageData = planService.page(projectId, status, keywords, page, pageSize);
        List<TestPlanVo> vos = planService.genVos(pageData.getItems());

        ret.put("collectionSize", pageData.getTotal());
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
		Long projectId = json.getLong("projectId");
        Long id = json.getLong("id");

        TestPlanVo vo = planService.getById(id);
		List<TestSuite> ls = suiteService.query(projectId, null);
		List<TestSuiteVo> suites = suiteService.genVos(ls);

		List<TestVerVo> vers = verService.listVos(projectId);
		List<TestEnvVo> envs = envService.listVos(projectId);

        ret.put("data", vo);
		ret.put("suites", suites);
		ret.put("vers", vers);
		ret.put("envs", envs);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

	@AuthPassport(validate = true)
	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

		TestPlan po = planService.save(json, userVo);
		TestPlanVo vo = planService.genVo(po);

		optFacade.opt(WsConstant.WS_TODO, userVo.getId().toString());

		ret.put("data", vo);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@AuthPassport(validate = true)
	@RequestMapping(value = "delete", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		Long id = json.getLong("id");

		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

		TestPlan po = planService.delete(id, userVo.getId());

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
