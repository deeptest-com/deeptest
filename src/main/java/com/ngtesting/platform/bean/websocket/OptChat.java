package com.ngtesting.platform.bean.websocket;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.service.intf.ChatService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.HashMap;
import java.util.Map;

@Service
public class OptChat {

    @Autowired
    ChatService chatService;

	public Map<String, Object> enter(JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		Long eventId = json.getLong("eventId");
        Long clientId = json.getLong("clientId");

//		List<TstThread> pos = chatService.enter(eventId, clientId);
//        List<TstThread> vos = new LinkedList<TstThread>();
//        for (TstThread po: pos) {
//            TstThread vo = new ThreadVo();
//            TstThread.copyProperties(vo, po);
//        	vos.save(vo);
//        }

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
//		ret.put("threads", vos);

		return ret;
	}

    public Map<String, Object> chat(JSONObject json){
        Map<String, Object> ret = new HashMap<String, Object>();

        Long eventId = json.getLong("eventId");
        Long parentId = json.getLong("parentId");
        Long clientId = json.getLong("clientId");
        String content = json.getString("content");

//        TestThread thread = chatService.save(eventId, parentId, clientId, content);
//    	ThreadVo vo = new ThreadVo();
//    	BeanUtilEx.copyProperties(vo, thread);

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
//        ret.put("data", vo);

        return ret;
    }

}
