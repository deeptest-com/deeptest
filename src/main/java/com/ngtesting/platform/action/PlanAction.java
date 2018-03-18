package com.ngtesting.platform.action;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.bean.websocket.OptFacade;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.config.WsConstant;
import com.ngtesting.platform.entity.TestPlan;
import com.ngtesting.platform.entity.TestSuite;
import com.ngtesting.platform.service.PlanService;
import com.ngtesting.platform.service.SuiteService;
import com.ngtesting.platform.util.AuthPassport;
import com.ngtesting.platform.vo.TestPlanVo;
import com.ngtesting.platform.vo.TestSuiteVo;
import com.ngtesting.platform.vo.UserVo;
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

	@AuthPassport(validate = true)
	@RequestMapping(value = "query", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> query(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		List<TestPlan> ls = planService.query(json);

		List<TestPlanVo> vos = planService.genVos(ls);

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
		List<TestSuite> ls = suiteService.query(projectId, null, null);
		List<TestSuiteVo> suites = suiteService.genVos(ls);

        ret.put("data", vo);
		ret.put("suites", suites);
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
