package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSONObject;
import com.github.pagehelper.PageHelper;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.bean.websocket.WsFacade;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.config.WsConstant;
import com.ngtesting.platform.model.TstMsg;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.TestCustomFieldService;
import com.ngtesting.platform.service.MsgService;
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
@RequestMapping(Constant.API_PATH_CLIENT + "msg/")
public class MsgAction extends BaseAction {
    @Autowired
    private WsFacade optFacade;

	@Autowired
    MsgService msgService;

    @Autowired
    TestCustomFieldService customFieldService;

	@RequestMapping(value = "list", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> list(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

        String keywords = json.getString("keywords");
        Boolean isRead = json.getBoolean("isRead");
        Integer pageNum = json.getInteger("page");
        Integer pageSize = json.getInteger("pageSize");

        com.github.pagehelper.Page page = PageHelper.startPage(pageNum, pageSize);
        List<TstMsg> ls = msgService.list(user.getId(), isRead, keywords);

        ret.put("total", page.getTotal());
        ret.put("data", ls);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "delete", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		Integer id = json.getInteger("id");

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

		Boolean result = msgService.delete(id, userVo.getId());
		if (!result) {
            return authFail();
        }

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

    @RequestMapping(value = "markRead", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> markRead(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

        Integer id = json.getInteger("id");
        Boolean result = msgService.markRead(id, user.getId());
        if (!result) {
            return authFail();
        }

        optFacade.opt(WsConstant.WS_TODO, user);

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @RequestMapping(value = "markAllRead", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> markAllRead(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

		msgService.markAllRead(user.getId());
		optFacade.opt(WsConstant.WS_TODO, user);

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

}
