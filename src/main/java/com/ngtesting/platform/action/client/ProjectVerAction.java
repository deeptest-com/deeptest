package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.bean.websocket.WsFacade;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.dao.TestVerDao;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.model.TstVer;
import com.ngtesting.platform.service.intf.TestVerService;
import com.ngtesting.platform.servlet.PrivOrg;
import com.ngtesting.platform.servlet.PrivPrj;
import org.apache.shiro.SecurityUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import javax.servlet.http.HttpServletRequest;
import java.util.HashMap;
import java.util.List;
import java.util.Map;


@RestController
@RequestMapping(Constant.API_PATH_CLIENT + "ver/")
public class ProjectVerAction extends BaseAction {
	@Autowired
	private WsFacade optFacade;

	@Autowired
	TestVerService verService;
	@Autowired
	TestVerDao verDao;

	@RequestMapping(value = "list", method = RequestMethod.POST)
	@PrivPrj(perms = {"project:*"})
	public Map<String, Object> list(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();
		Integer prjId = user.getDefaultPrjId();

		String keywords = json.getString("keywords");
		Boolean disabled = json.getBoolean("disabled");

		List<TstVer> ls = verService.list(prjId, keywords, disabled);

        ret.put("data", ls);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

    @RequestMapping(value = "get", method = RequestMethod.POST)
	@PrivPrj(perms = {"project:*"})
    public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
		TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();
		Integer prjId = user.getDefaultPrjId();

        Integer id = json.getInteger("id");

		TstVer vo = verService.getById(id, prjId);

        ret.put("data", vo);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	@PrivOrg(perms = {"project:*"})
	public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();

		TstVer po = verService.save(json, user);
        if(po == null) {
            return authorFail();
        }

		ret.put("data", po);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "delete", method = RequestMethod.POST)
	@ResponseBody
	@PrivOrg(perms = {"project:*"})
	public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();
		Integer prjId = user.getDefaultPrjId();

		Integer id = json.getInteger("id");

		Boolean result = verService.delete(id, prjId);
        if (!result) {
            return authorFail();
        }

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "changeOrder", method = RequestMethod.POST)
	@ResponseBody
	@PrivOrg(perms = {"project:*"})
	public Map<String, Object> changeOrder(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();
        Integer prjId = user.getDefaultPrjId();

        Integer id = json.getInteger("id");
        String act = json.getString("act");

        String keywords = json.getString("keywords");
        Boolean disabled = json.getBoolean("disabled");

        Boolean result = verService.changeOrder(id, act, prjId);
        if (!result) {
            return authorFail();
        }

		List<TstVer> vos = verService.list(prjId, keywords, disabled);

		ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());

		return ret;
	}

}
