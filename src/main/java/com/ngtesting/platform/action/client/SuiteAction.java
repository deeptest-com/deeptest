package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.github.pagehelper.PageHelper;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.bean.websocket.WsFacade;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstSuite;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.TestSuiteService;
import com.ngtesting.platform.servlet.PrivPrj;
import org.apache.shiro.SecurityUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import javax.servlet.http.HttpServletRequest;
import java.util.HashMap;
import java.util.List;
import java.util.Map;


@RestController
@RequestMapping(Constant.API_PATH_CLIENT + "suite/")
public class SuiteAction extends BaseAction {
	@Autowired
	private WsFacade optFacade;

	@Autowired
	TestSuiteService suiteService;

	@RequestMapping(value = "query", method = RequestMethod.POST)
	@ResponseBody
	@PrivPrj(perms = {"test_suite-view"})
	public Map<String, Object> query(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();
        Integer prjId = user.getDefaultPrjId();

		String keywords = json.getString("keywords");
        Boolean disabled = json.getBoolean("disabled");
		Integer pageNum = json.getInteger("page");
		Integer pageSize = json.getInteger("pageSize");

		com.github.pagehelper.Page page = PageHelper.startPage(pageNum, pageSize);
		List ls = suiteService.listByPage(prjId, keywords,disabled);

        ret.put("total", page.getTotal());
        ret.put("data", ls);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

    @RequestMapping(value = "get", method = RequestMethod.POST)
	@PrivPrj(perms = {"test_suite-view"})
    public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
		TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();
        Integer prjId = user.getDefaultPrjId();

        Integer id = json.getInteger("id");

		TstSuite vo = suiteService.get(id, prjId);

        ret.put("data", vo);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	@PrivPrj(perms = {"test_suite-maintain"})
	public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();

		TstSuite po = suiteService.save(json, user);
		if (po == null) {
		  return authorFail();
        }

		ret.put("data", po);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

    @RequestMapping(value = "saveCases", method = RequestMethod.POST)
	@PrivPrj(perms = {"test_suite-maintain"})
    public Map<String, Object> saveCases(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
		TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();
        Integer prjId = user.getDefaultPrjId();

        Integer caseProjectId = json.getInteger("caseProjectId");
        Integer suiteId = json.getInteger("suiteId");
        List<Integer> ids = JSON.parseArray(json.getString("cases"), Integer.class);

        TstSuite suite = suiteService.get(suiteId, prjId);
        if (suite == null) { // suite和project不匹配
            return authorFail();
        }

        TstSuite po = suiteService.saveCases(prjId, caseProjectId, suiteId, ids, user);

        ret.put("data", po);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

	@RequestMapping(value = "delete", method = RequestMethod.POST)
	@ResponseBody
	@PrivPrj(perms = {"test_suite-delete"})
	public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();
		Integer prjId = user.getDefaultPrjId();

		Integer id = json.getInteger("id");

		Boolean result = suiteService.delete(id, prjId);
        if (!result) {
            return authorFail();
        }

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
