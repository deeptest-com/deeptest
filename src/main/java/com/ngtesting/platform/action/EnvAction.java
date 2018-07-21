package com.ngtesting.platform.action;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.bean.websocket.OptFacade;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstEnv;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.TestEnvService;
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
@RequestMapping(Constant.API_PATH_CLIENT + "env/")
public class EnvAction extends BaseAction {
	@Autowired
	private OptFacade optFacade;

	@Autowired
	TestEnvService envService;

	@RequestMapping(value = "list", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> List(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
        TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		Integer projectId = json.getInteger("projectId");
        String keywords = json.getString("keywords");
        String disabled = json.getString("disabled");

		List<TstEnv> ls = envService.list(projectId, keywords, disabled);

		List<TstEnv> vos = envService.genVos(ls);

        ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}


    @RequestMapping(value = "get", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
        TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

		Integer id = json.getInteger("id");

        TstEnv vo = envService.getById(id);

        ret.put("data", vo);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }


	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

		TstEnv po = envService.save(json, userVo);
		TstEnv vo = envService.genVo(po);

//		optFacade.opt(WsConstant.WS_TODO, userVo.getId().toString());

		ret.put("data", vo);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}


	@RequestMapping(value = "delete", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
        TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

		Integer id = json.getInteger("id");

		TstEnv po = envService.delete(id, userVo.getId());

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}


    @RequestMapping(value = "changeOrder", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> changeOrder(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		Integer projectId = json.getInteger("projectId");
		Integer id = json.getInteger("id");
        String act = json.getString("act");

        boolean success = envService.changeOrderPers(id, act, projectId);
        List<TstEnv> vos = envService.listVos(projectId);

        ret.put("data", vos);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());

        return ret;
    }

}
