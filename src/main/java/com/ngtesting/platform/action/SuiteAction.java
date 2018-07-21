package com.ngtesting.platform.action;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.bean.websocket.OptFacade;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstSuite;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.TestSuiteService;
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
@RequestMapping(Constant.API_PATH_CLIENT + "suite/")
public class SuiteAction extends BaseAction {
	@Autowired
	private OptFacade optFacade;

	@Autowired
	TestSuiteService suiteService;

	@RequestMapping(value = "query", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> query(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		int page = json.getInteger("page") == null? 0: json.getInteger("page") - 1;
		int pageSize = json.getInteger("pageSize") == null? Constant.PAGE_SIZE: json.getInteger("pageSize");
        Integer projectId = json.getInteger("projectId");
        String keywords = json.getString("keywords");

		Page pageData = suiteService.page(projectId, keywords, page, pageSize);
		List<TstSuite> vos = suiteService.genVos(pageData.getItems());

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
        Integer id = json.getInteger("id");

		TstSuite vo = suiteService.getById(id, false);

        ret.put("data", vo);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

		TstSuite po = suiteService.save(json, userVo);
		TstSuite vo = suiteService.genVo(po);

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

		TstSuite po = suiteService.delete(id, userVo.getId());

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "saveCases", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> saveCases(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

		TstSuite po = suiteService.saveCases(json, userVo);
		TstSuite vo = suiteService.genVo(po);

		ret.put("data", vo);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
