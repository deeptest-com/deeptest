package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.bean.websocket.WsFacade;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.config.WsConstant;
import com.ngtesting.platform.model.TstTask;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.TestTaskService;
import com.ngtesting.platform.servlet.PrivPrj;
import org.apache.shiro.SecurityUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import javax.servlet.http.HttpServletRequest;
import java.util.HashMap;
import java.util.Map;


@RestController
@RequestMapping(Constant.API_PATH_CLIENT + "task/")
public class TaskAction extends BaseAction {
    @Autowired
    private WsFacade optFacade;

	@Autowired
	TestTaskService taskService;

    @RequestMapping(value = "get", method = RequestMethod.POST)
    @PrivPrj(perms = {"test_task-view"})
    public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
		TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();
		Integer prjId = user.getDefaultPrjId();

        Integer runId = json.getInteger("id");

        TstTask vo = taskService.getById(runId, prjId);

        ret.put("data", vo);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
    @PrivPrj(perms = {"test_plan-maintain"})
	public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();
        Integer prjId = user.getDefaultPrjId();

		TstTask po = taskService.save(json, user);
        if (po == null) {
            return authorFail();
        }

		TstTask vo = taskService.getById(po.getId(), prjId);

        optFacade.opt(WsConstant.WS_TODO, user);

		ret.put("data", vo);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "saveCases", method = RequestMethod.POST)
	@ResponseBody
    @PrivPrj(perms = {"test_plan-maintain"})
	public Map<String, Object> saveCases(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();
        Integer prjId = user.getDefaultPrjId();

		TstTask po = taskService.saveCases(json, user);
		TstTask caseVo = taskService.getById(po.getId(), prjId);

		ret.put("data", caseVo);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

    @RequestMapping(value = "delete", method = RequestMethod.POST)
    @PrivPrj(perms = {"test_plan-delete"})
    public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();
        Integer prjId = user.getDefaultPrjId();

        Integer id = json.getInteger("id");

        Boolean result = taskService.delete(id, prjId);
        if (!result) {
            return authorFail();
        }

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @RequestMapping(value = "close", method = RequestMethod.POST)
    @PrivPrj(perms = {"test_task-close"})
    public Map<String, Object> close(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();
        Integer prjId = user.getDefaultPrjId();

        Integer id = json.getInteger("id");

        Boolean result = taskService.close(id, prjId);
        if (!result) {
            return authorFail();
        }

        taskService.closePlanIfAllTaskClosed(id);
        TstTask vo = taskService.getById(id, prjId);

        ret.put("data", vo);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

}
