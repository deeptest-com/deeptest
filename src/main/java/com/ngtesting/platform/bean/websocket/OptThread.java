package com.ngtesting.platform.bean.websocket;

import java.util.HashMap;
import java.util.Map;

import org.springframework.context.annotation.Bean;
import org.springframework.stereotype.Service;

import com.ngtesting.platform.websocket.SystemWebSocketHandler;

import com.alibaba.fastjson.JSONObject;

@Service
public class OptThread {

    @Bean
    public SystemWebSocketHandler systemWebSocketHandler() {
        return new SystemWebSocketHandler();
    }


    public Map<String, Object> dosomething(JSONObject json){
        Map<String, Object> ret = new HashMap<String, Object>();

        return ret;
    }
}
