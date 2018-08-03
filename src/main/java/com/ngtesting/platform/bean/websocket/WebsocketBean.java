package com.ngtesting.platform.bean.websocket;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.WsConstant;
import com.ngtesting.platform.model.TstUser;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.messaging.handler.annotation.MessageMapping;
import org.springframework.messaging.simp.SimpMessageHeaderAccessor;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestBody;

import java.util.HashMap;
import java.util.Map;

@Controller
public class WebsocketBean {

    @Autowired
    private WsFacade optFacade;

    @Value("${mail.smtp.host}")
    private String mailHost;

    @MessageMapping("/comm") // 加上前缀send后，客户端使用/send/comm发消息
//    @SendTo("/topic/comm") // 返回的目标，客户端使用/topic/comm订阅消息
//    @SendToUser(value = "/topic/comm", broadcast = false)
    public Map topicComm(@RequestBody JSONObject json, SimpMessageHeaderAccessor headerAccessor) throws Exception {
        Map map = headerAccessor.getSessionAttributes();
        TstUser user = (TstUser)map.get(WsConstant.WS_USER_KEY);
        optFacade.opt(json, user);

        Map ret = new HashMap();
        ret.put("code", 1);
        ret.put("from", "topicComm method");

        return ret;
    }

}
