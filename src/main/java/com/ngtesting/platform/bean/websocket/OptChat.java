package com.ngtesting.platform.bean.websocket;

import java.util.HashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.stereotype.Service;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.TestThread;
import com.ngtesting.platform.service.ChatService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.util.Constant;
import com.ngtesting.platform.vo.ThreadVo;
import com.ngtesting.platform.websocket.SystemWebSocketHandler;

@Service
public class OptChat {

    @Autowired
    ChatService chatService;

    @Bean
    public SystemWebSocketHandler systemWebSocketHandler() {
        return new SystemWebSocketHandler();
    }

	public Map<String, Object> enter(JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		Long eventId = json.getLong("eventId");
        Long clientId = json.getLong("clientId");

		List<TestThread> pos = chatService.enter(eventId, clientId);
        List<ThreadVo> vos = new LinkedList<ThreadVo>();
        for (TestThread po: pos) {
        	ThreadVo vo = new ThreadVo();
        	BeanUtilEx.copyProperties(vo, po);
        	vos.add(vo);
        }

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		ret.put("threads", vos);

		return ret;
	}

    public Map<String, Object> chat(JSONObject json){
        Map<String, Object> ret = new HashMap<String, Object>();

        Long eventId = json.getLong("eventId");
        Long parentId = json.getLong("parentId");
        Long clientId = json.getLong("clientId");
        String content = json.getString("content");

        TestThread thread = chatService.save(eventId, parentId, clientId, content);
    	ThreadVo vo = new ThreadVo();
    	BeanUtilEx.copyProperties(vo, thread);

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        ret.put("data", vo);

        return ret;
    }

}
