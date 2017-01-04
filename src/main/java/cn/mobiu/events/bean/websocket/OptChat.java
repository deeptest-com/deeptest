package cn.mobiu.events.bean.websocket;

import java.util.HashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.stereotype.Service;

import cn.mobiu.events.constants.Constant;
import cn.mobiu.events.entity.EvtThread;
import cn.mobiu.events.service.ChatService;
import cn.mobiu.events.util.BeanUtilEx;
import cn.mobiu.events.vo.ThreadVo;
import cn.mobiu.events.websocket.SystemWebSocketHandler;

import com.alibaba.fastjson.JSONObject;

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
        
		List<EvtThread> pos = chatService.enter(eventId, clientId);
        List<ThreadVo> vos = new LinkedList<ThreadVo>();
        for (EvtThread po: pos) {
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
        
        EvtThread thread = chatService.save(eventId, parentId, clientId, content);
    	ThreadVo vo = new ThreadVo();
    	BeanUtilEx.copyProperties(vo, thread);
        
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        ret.put("data", vo);
        
        return ret;
    }

}
