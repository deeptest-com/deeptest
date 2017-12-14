package com.ngtesting.platform.bean.websocket;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant.RespCode;
import com.ngtesting.platform.config.WsConstant;
import com.ngtesting.platform.websocket.SystemWebSocketHandler;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.stereotype.Service;

import java.util.Map;
import java.util.concurrent.TimeoutException;

@Service
public class OptFacade {
    Log logger = LogFactory.getLog(OptFacade.class);

    @Autowired
    OptNews optNews;
    @Autowired
    OptQa optQa;
    @Autowired
    OptChat optChat;

    @Bean
    public SystemWebSocketHandler systemWebSocketHandler() {
        return new SystemWebSocketHandler();
    }

    public Map<String, Object> opt(JSONObject json) {
        Map<String, Object> ret = null;

        String act = json.getString("act");
        String trans = json.getString("trans");

        try {
            if (WsConstant.WS__OPT_ENTER_CHAT_ROOM.equals(act)) {
                ret = optChat.enter(json);
            } else if (WsConstant.WS_OPT_CHAT.equals(act)) {
                ret = optChat.chat(json);
            }
        } catch (Exception e) {
            if (e.getCause() instanceof TimeoutException) {
                ret.put("code", RespCode.BIZ_FAIL.getCode());
                ret.put("msg", "操作超时！");
            }else{
                ret.put("msg", "操作出错，请稍后重试");
            }
        }
        if(ret != null) {
        	ret.put("trans", trans);
        	ret.put("type", act);
        	ret.put("code", 1);
        }
        return ret;
    }

}
