package com.ngtesting.platform.bean.websocket;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.bean.ApplicationScopeBean;
import com.ngtesting.platform.config.Constant.RespCode;
import com.ngtesting.platform.config.WsConstant;
import com.ngtesting.platform.service.AlertService;
import com.ngtesting.platform.service.MsgService;
import com.ngtesting.platform.util.ExceptionUtil;
import com.ngtesting.platform.vo.TestAlertVo;
import com.ngtesting.platform.vo.TestMsgVo;
import com.ngtesting.platform.websocket.SystemWebSocketHandler;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.stereotype.Service;
import org.springframework.web.socket.TextMessage;

import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.concurrent.TimeoutException;

@Service
public class OptFacade {
    Log logger = LogFactory.getLog(OptFacade.class);

    @Autowired
    private ApplicationScopeBean scopeBean;

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

    @Bean
    public SystemWebSocketHandler systemWebSocketHandler() {
        return new SystemWebSocketHandler();
    }

    public void opt(JSONObject json, String userId) {
        Map<String, Object> ret = new HashMap<>();
        String type = json.getString("type");

        try {
            if (WsConstant.WS_OPEN.equals(type) || WsConstant.WS_TODO.equals(type)) {
                List<TestMsgVo> msgs = msgService.list(Long.valueOf(userId), false);
                List<TestAlertVo> alerts = alertService.list(Long.valueOf(userId), false);

                ret.put("msgs", msgs);
                ret.put("alerts", alerts);

                ret.put("type", WsConstant.WS_MSG_AND_ALERT_LASTEST);
            } else if (WsConstant.WS_OPT_ENTER_CHAT_ROOM.equals(type)) {
                ret = optChat.enter(json);
            } else if (WsConstant.WS_OPT_CHAT.equals(type)) {
                ret = optChat.chat(json);
            }
        } catch (Exception e) {
            logger.error(ExceptionUtil.GetExceptionInfo(e));
            if (e.getCause() instanceof TimeoutException) {
                ret.put("code", RespCode.BIZ_FAIL.getCode());
                ret.put("msg", "操作超时！");
            }else{
                ret.put("msg", "操作出错，请稍后重试");
            }
        }

        ret.put("code", 1);
        if (ret.get("type") != null) {
            scopeBean.sendMessageToClient(userId, new TextMessage(JSON.toJSONString(ret)));
        }
    }

    public void opt(String act, String userId) {
        JSONObject json = new JSONObject();
        json.put("type", act);
        opt(json, userId);
    }

}
