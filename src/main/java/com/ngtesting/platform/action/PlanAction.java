package com.ngtesting.platform.action;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.bean.websocket.OptFacade;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.*;
import com.ngtesting.platform.service.TestEnvService;
import com.ngtesting.platform.service.TestPlanService;
import com.ngtesting.platform.service.TestSuiteService;
import com.ngtesting.platform.service.TestVerService;
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
@RequestMapping(Constant.API_PATH_CLIENT + "plan/")
public class PlanAction extends BaseAction {
	@Autowired
	private OptFacade optFacade;

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
	public Map<String, Object> query(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

        int page = json.getInteger("page") == null? 0: json.getInteger("page") - 1;
        int pageSize = json.getInteger("pageSize") == null? Constant.PAGE_SIZE: json.getInteger("pageSize");

        Integer projectId = json.getInteger("projectId");
        String status = json.getString("status");
        String keywords = json.getString("keywords");

        Page pageData = planService.page(projectId, status, keywords, page, pageSize);
        List<TstPlan> vos = planService.genVos(pageData.getItems());

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
		Integer projectId = json.getInteger("projectId");
        Integer id = json.getInteger("id");

		TstPlan vo = planService.getById(id);
		List<TstSuite> ls = suiteService.query(projectId, null);
		List<TstSuite> suites = suiteService.genVos(ls);

		List<TstVer> vers = verService.list(projectId, null, null);
		List<TstEnv> envs = envService.list(projectId, null, null);

        ret.put("data", vo);
		ret.put("suites", suites);
		ret.put("vers", vers);
		ret.put("envs", envs);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }


	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

		TstPlan po = planService.save(json, userVo);
		TstPlan vo = planService.genVo(po);

//		optFacade.opt(WsConstant.WS_TODO, userVo.getId().toString());

		ret.put("data", vo);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}


	@RequestMapping(value = "delete", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		Integer id = json.getInteger("id");

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

		TstPlan po = planService.delete(id, userVo.getId());

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
