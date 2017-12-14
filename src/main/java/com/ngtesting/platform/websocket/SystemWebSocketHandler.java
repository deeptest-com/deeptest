package com.ngtesting.platform.websocket;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.bean.ApplicationScopeBean;
import com.ngtesting.platform.bean.websocket.OptFacade;
import com.ngtesting.platform.config.WsConstant;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.socket.CloseStatus;
import org.springframework.web.socket.WebSocketHandler;
import org.springframework.web.socket.WebSocketMessage;
import org.springframework.web.socket.WebSocketSession;

import java.util.Date;

public class SystemWebSocketHandler implements WebSocketHandler {

    private static final Logger logger = LoggerFactory.getLogger(SystemWebSocketHandler.class);
    @Autowired
    private OptFacade optFacade;

    @Autowired
    private ApplicationScopeBean scopeBean;

    @Override
    public void handleMessage(WebSocketSession session, WebSocketMessage<?> message) throws Exception {
        logger.info("hander ws message......" + message.getPayload().toString());
        String userId = (String) session.getAttributes().get(WsConstant.WS_USER_KEY);

        String str = message.getPayload().toString();
        JSONObject json = (JSONObject) JSONObject.parse(str);

        optFacade.opt(json, userId);
    }

    @Override
    public void afterConnectionEstablished(WebSocketSession session) throws Exception {
        logger.debug("connect to the websocket success......");

        String userId = (String) session.getAttributes().get(WsConstant.WS_USER_KEY);

        if (!session.getAttributes().containsKey(WsConstant.WS_TIMESNAP)) {
            session.getAttributes().put(WsConstant.WS_TIMESNAP, new Date().getTime());
        }

        if (userId != null) {
        	scopeBean.getOnlines().put(userId, session);

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
        if (session.getAttributes().get(WsConstant.WS_USER_KEY) != null) {
        	String userId = (String) session.getAttributes().get(WsConstant.WS_USER_KEY);
            if (userId != null && scopeBean.getOnlines().get(userId) != null) {
                scopeBean.getOnlines().remove(userId);
            }
        }
    }

    @Override
    public void afterConnectionClosed(WebSocketSession session, CloseStatus closeStatus) throws Exception {
        logger.debug("websocket connection closed......");
        String clientId = (String) session.getAttributes().get(WsConstant.WS_USER_KEY);
        if (clientId != null && scopeBean.getOnlines().get(clientId) != null) {
        	scopeBean.getOnlines().remove(clientId);
        }
    }

    @Override
    public boolean supportsPartialMessages() {
        return false;
    }
}
