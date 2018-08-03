package com.ngtesting.platform.action;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.bean.websocket.WsFacade;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstMsg;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.CustomFieldService;
import com.ngtesting.platform.service.MsgService;
import com.ngtesting.platform.vo.Page;
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
@RequestMapping(Constant.API_PATH_CLIENT + "msg/")
public class MsgAction extends BaseAction {
    @Autowired
    private WsFacade optFacade;

	@Autowired
    MsgService msgService;

    @Autowired
    CustomFieldService customFieldService;

	@RequestMapping(value = "list", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> list(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

        String keywords = json.getString("keywords");
        String isRead = json.getString("isRead");
        int page = json.getInteger("listByPage") == null? 0: json.getInteger("listByPage") - 1;
        int pageSize = json.getInteger("pageSize") == null? Constant.PAGE_SIZE: json.getInteger("pageSize");

        Page pageDate = msgService.listByPage(userVo.getId(), isRead, keywords, page, pageSize);
//        List<TstUser> vos = msgService.genVos(pageDate.getItems());

        ret.put("collectionSize", pageDate.getTotal());
//        ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

    @RequestMapping(value = "get", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        Integer msgId = json.getInteger("id");

        TstMsg vo = msgService.getById(msgId);

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

		msgService.delete(id, userVo.getId());

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

    @RequestMapping(value = "markRead", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> markRead(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

        Integer id = json.getInteger("id");
        TstMsg msg = msgService.markReadPers(id, userVo.getId());
//        TestMsgVo vo = msgService.genVo(msg);

//        optFacade.opt(WsConstant.WS_TODO, userVo.getId().toString());

//        ret.put("data", vo);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @RequestMapping(value = "markAllRead", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> markAllRead(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

		msgService.markAllReadPers(userVo.getId());
//		optFacade.opt(WsConstant.WS_TODO, userVo.getId().toString());

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

}
