package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.bean.websocket.WsFacade;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.config.WsConstant;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.AlertService;
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
@RequestMapping(Constant.API_PATH_CLIENT + "alert/")
public class AlertAction extends BaseAction {
    @Autowired
    private WsFacade optFacade;

	@Autowired
    AlertService alertService;

    @RequestMapping(value = "markAllRead", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> markAllRead(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

		alertService.markAllRead(json.getString("ids"), user.getId());
        optFacade.opt(WsConstant.WS_TODO, user);

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

}
