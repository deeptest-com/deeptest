package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.bean.websocket.WsFacade;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.config.WsConstant;
import com.ngtesting.platform.model.TstTask;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.TestTaskService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.ResponseBody;

import javax.servlet.http.HttpServletRequest;
import java.util.HashMap;
import java.util.Map;


@Controller
@RequestMapping(Constant.API_PATH_CLIENT + "task/")
public class TaskAction extends BaseAction {
    @Autowired
    private WsFacade optFacade;

	@Autowired
	TestTaskService taskService;

    @RequestMapping(value = "get", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer projectId = user.getDefaultPrjId();

        Integer runId = json.getInteger("id");

        TstTask vo = taskService.getById(runId, projectId);

        ret.put("data", vo);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @RequestMapping(value = "delete", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer projectId = user.getDefaultPrjId();

        Integer id = json.getInteger("id");

        Boolean result = taskService.delete(id, projectId);
        if (!result) {
            return authFail();
        }

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

	@RequestMapping(value = "close", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> close(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer projectId = user.getDefaultPrjId();

		Integer id = json.getInteger("id");

		Boolean result = taskService.close(id, projectId);
		if (!result) {
		    return authFail();
        }

		taskService.closePlanIfAllTaskClosed(id);
		TstTask vo = taskService.getById(id, projectId);

        ret.put("data", vo);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer projectId = user.getDefaultPrjId();

		TstTask po = taskService.save(json, user);
        if (po == null) {
            return authFail();
        }

		TstTask vo = taskService.getById(po.getId(), projectId);

        optFacade.opt(WsConstant.WS_TODO, user);

		ret.put("data", vo);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "saveCases", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> saveCases(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer projectId = user.getDefaultPrjId();

		TstTask po = taskService.saveCases(json, user);
		TstTask caseVo = taskService.getById(po.getId(), projectId);

		ret.put("data", caseVo);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
