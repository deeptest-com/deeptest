package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.bean.websocket.WsFacade;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.dao.TestEnvDao;
import com.ngtesting.platform.model.TstEnv;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.TestEnvService;
import com.ngtesting.platform.servlet.PrivPrj;
import org.apache.shiro.SecurityUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

import javax.servlet.http.HttpServletRequest;
import java.util.HashMap;
import java.util.List;
import java.util.Map;


@RestController
@RequestMapping(Constant.API_PATH_CLIENT + "env/")
public class ProjectEnvAction extends BaseAction {
	@Autowired
	private WsFacade optFacade;

	@Autowired
	TestEnvService envService;
    @Autowired
    TestEnvDao envDao;

    @RequestMapping(value = "list", method = RequestMethod.POST)
    @PrivPrj(perms = {"project:*"})
    public Map<String, Object> list(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();
        Integer prjId = user.getDefaultPrjId();

        String keywords = json.getString("keywords");
        Boolean disabled = json.getBoolean("disabled");

        List<TstEnv> ls = envService.list(prjId, keywords, disabled);

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

        TstEnv vo = envService.getById(id, prjId);

        ret.put("data", vo);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @RequestMapping(value = "save", method = RequestMethod.POST)
    @PrivPrj(perms = {"project:*"})
    public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();

        TstEnv po = envService.save(json, user);
        if(po == null) {
            return authorFail();
        }

        ret.put("data", po);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @RequestMapping(value = "delete", method = RequestMethod.POST)
    @PrivPrj(perms = {"project:*"})
    public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();
        Integer prjId = user.getDefaultPrjId();

        Integer id = json.getInteger("id");

        Boolean result = envService.delete(id, prjId);
        if (!result) {
            return authorFail();
        }

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @RequestMapping(value = "changeOrder", method = RequestMethod.POST)
    @PrivPrj(perms = {"project:*"})
    public Map<String, Object> changeOrder(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();
        Integer prjId = user.getDefaultPrjId();

        Integer id = json.getInteger("id");
        String act = json.getString("act");

        String keywords = json.getString("keywords");
        Boolean disabled = json.getBoolean("disabled");

        Boolean result = envService.changeOrder(id, act, prjId);
        if (!result) {
            return authorFail();
        }

        List<TstEnv> vos = envService.list(prjId, keywords, disabled);

        ret.put("data", vos);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());

        return ret;
    }

}
