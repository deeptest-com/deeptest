package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSONObject;
import com.github.pagehelper.PageHelper;
import com.itfsw.query.builder.support.model.JsonRule;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.IsuQuery;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.IssueQueryService;
import com.ngtesting.platform.servlet.PrivPrj;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

import javax.servlet.http.HttpServletRequest;
import java.util.HashMap;
import java.util.List;
import java.util.Map;


@Controller
@RequestMapping(Constant.API_PATH_CLIENT + "issue_query/")
public class IssueQueryAction extends BaseAction {
	private static final Log log = LogFactory.getLog(IssueQueryAction.class);

	@Autowired
	IssueQueryService queryService;

    @ResponseBody
    @PostMapping("/list")
    @PrivPrj
    public Object list(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer prjId = user.getDefaultPrjId();

        String keywords = json.getString("keywords");
        Integer pageNum = json.getInteger("page");
        Integer pageSize = json.getInteger("pageSize");

        com.github.pagehelper.Page page = PageHelper.startPage(pageNum, pageSize);
        List<IsuQuery> vos = queryService.list(prjId, user.getId(), keywords);

        ret.put("total", page.getTotal());
        ret.put("data", vos);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());

        return ret;
    }

    @ResponseBody
    @PostMapping("/get")
    @PrivPrj
    public Object get(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

        Integer id = json.getInteger("id");

        IsuQuery vo = queryService.get(id, user.getId());

        ret.put("data", vo);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());

        return ret;
    }

	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
    @PrivPrj
	public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

        String queryName = json.getString("queryName");
		JsonRule rule = json.getObject("rule", JsonRule.class);

		queryService.save(queryName, rule, user);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

    @RequestMapping(value = "update", method = RequestMethod.POST)
    @ResponseBody
    @PrivPrj
    public Map<String, Object> update(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

        IsuQuery vo = json.getObject("model", IsuQuery.class);

        queryService.update(vo, user);

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @RequestMapping(value = "delete", method = RequestMethod.POST)
    @ResponseBody
    @PrivPrj
    public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

        Integer id = json.getInteger("id");
        queryService.delete(id, user);

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

}
