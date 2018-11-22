package com.ngtesting.platform.bean.websocket;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant.RespCode;
import com.ngtesting.platform.config.WsConstant;
import com.ngtesting.platform.model.TstAlert;
import com.ngtesting.platform.model.TstMsg;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.AlertService;
import com.ngtesting.platform.service.intf.MsgService;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.messaging.simp.SimpMessagingTemplate;
import org.springframework.stereotype.Service;
import org.springframework.web.socket.TextMessage;

import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.concurrent.TimeoutException;

@Service
public class WsFacade {
    Log logger = LogFactory.getLog(WsFacade.class);

    @Autowired
    private SimpMessagingTemplate simpMessagingTemplate;

    @Autowired
    OptNews optNews;
    @Autowired
    OptQa optQa;
    @Autowired
    OptChat optChat;

    @Autowired
    MsgService msgService;
    @Autowired
    AlertService alertService;

    public void opt(String act, TstUser user) {
        JSONObject json = new JSONObject();
        json.put("type", act);
        opt(json, user);
    }

    public void opt(JSONObject json, TstUser user) {
        Map<String, Object> ret = new HashMap<>();
        String type = json.getString("type");

        try {
            if (WsConstant.WS_OPEN.equals(type) || WsConstant.WS_TODO.equals(type)) {
                List<TstMsg> msgs = msgService.list(user.getId(), false, null);
                List<TstAlert> alerts = alertService.list(user.getId(), false);

                ret.put("msgs", msgs);
                ret.put("alerts", alerts);

                ret.put("type", WsConstant.WS_MSG_AND_ALERT_LASTEST);
            } else if (WsConstant.WS_OPT_ENTER_CHAT_ROOM.equals(type)) {
                ret = optChat.enter(json);
            } else if (WsConstant.WS_OPT_CHAT.equals(type)) {
                ret = optChat.chat(json);
            }
        } catch (Exception e) {
            e.printStackTrace();
            if (e.getCause() instanceof TimeoutException) {
                ret.put("code", RespCode.BIZ_FAIL.getCode());
                ret.put("msg", "操作超时！");
            }else{
                ret.put("msg", "操作出错，请稍后重试");
            }
        }

        ret.put("code", 1);
        if (ret.get("type") != null) {
            simpMessagingTemplate.convertAndSendToUser(user.getToken(), "/notification",
                    new TextMessage(JSON.toJSONString(ret)));
        }
    }

}
