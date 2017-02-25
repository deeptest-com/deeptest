package cn.linkr.testspace.websocket;

import java.util.Date;
import java.util.Map;
import java.util.concurrent.ConcurrentSkipListSet;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.socket.CloseStatus;
import org.springframework.web.socket.TextMessage;
import org.springframework.web.socket.WebSocketHandler;
import org.springframework.web.socket.WebSocketMessage;
import org.springframework.web.socket.WebSocketSession;

import cn.linkr.testspace.bean.ApplicationScopeBean;
import cn.linkr.testspace.bean.websocket.OptFacade;
import cn.linkr.testspace.util.Constant;

import com.alibaba.fastjson.JSONObject;

public class SystemWebSocketHandler implements WebSocketHandler {

    private static final Logger logger = LoggerFactory.getLogger(SystemWebSocketHandler.class);
    @Autowired
    private OptFacade optFacade;
    
    @Autowired
    private ApplicationScopeBean scopeBean;

    @Override
    public void handleMessage(WebSocketSession session, WebSocketMessage<?> message) throws Exception {
        logger.debug("hander ws message......" + message.getPayload().toString());
        
        String str = message.getPayload().toString();
        JSONObject json = (JSONObject) JSONObject.parse(str);
        Map<String, Object> ret = optFacade.opt(json);
        if (ret != null) {
        	String clientId = (String) session.getAttributes().get(Constant.WEBSOCKET_CLIENT_KEY);
            scopeBean.sendMessageToClient(clientId, new TextMessage(JSONObject.toJSONString(ret)));
        }
    }

    @Override
    public void afterConnectionEstablished(WebSocketSession session) throws Exception {
        logger.debug("connect to the websocket success......");

        Long eventId = (Long) session.getAttributes().get(Constant.WEBSOCKET_EVENT_KEY);
        String clientId = (String) session.getAttributes().get(Constant.WEBSOCKET_CLIENT_KEY);

        if (!session.getAttributes().containsKey(Constant.WEBSOCKET_TIMESNAP)) {
            session.getAttributes().put(Constant.WEBSOCKET_TIMESNAP, new Date().getTime());
        }

        if (clientId != null) { // eventId != null
        	scopeBean.getOnlines().put(clientId, session);
        	
//        	if (!scopeBean.getChatroom().containsKey(eventId)) {
//        		scopeBean.getChatroom().put(eventId, new ConcurrentSkipListSet<Long>());
//        	}
//        	if (!scopeBean.getChatroom().get(eventId).contains(clientId)) {
//        		scopeBean.getChatroom().get(eventId).add(clientId);
//        	}
        }
    }

    @Override
    public void handleTransportError(WebSocketSession session, Throwable exception) throws Exception {
        if (session.isOpen()) {
            session.close();
        }
        
        logger.debug("websocket connection closed when error......");
        if (session.getAttributes().get(Constant.WEBSOCKET_CLIENT_KEY) != null) {
        	String clientId = (String) session.getAttributes().get(Constant.WEBSOCKET_CLIENT_KEY);
            if (clientId != null && scopeBean.getOnlines().get(clientId) != null) {
                scopeBean.getOnlines().remove(clientId);
            }
        }
    }

    @Override
    public void afterConnectionClosed(WebSocketSession session, CloseStatus closeStatus) throws Exception {
        logger.debug("websocket connection closed......");
        Long clientId = (Long) session.getAttributes().get(Constant.WEBSOCKET_CLIENT_KEY);
        if (clientId != null && scopeBean.getOnlines().get(clientId) != null) {
        	scopeBean.getOnlines().remove(clientId);
        }
    }

    @Override
    public boolean supportsPartialMessages() {
        return false;
    }
}