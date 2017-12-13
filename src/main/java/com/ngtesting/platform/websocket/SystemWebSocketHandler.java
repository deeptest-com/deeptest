package com.ngtesting.platform.websocket;

import com.ngtesting.platform.bean.ApplicationScopeBean;
import com.ngtesting.platform.bean.websocket.OptFacade;
import com.ngtesting.platform.util.Constant;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.socket.*;

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

        String str = message.getPayload().toString();

            String clientId = (String) session.getAttributes().get(Constant.WEBSOCKET_USER_KEY);
            scopeBean.sendMessageToClient(clientId, new TextMessage(str));

//            JSONObject json = (JSONObject) JSONObject.parse(str);
//            Map<String, Object> ret = optFacade.opt(json);
//            if (ret != null) {
//                String clientId = (String) session.getAttributes().get(Constant.WEBSOCKET_USER_KEY);
//                scopeBean.sendMessageToClient(clientId, new TextMessage(JSONObject.toJSONString(ret)));
//            }
    }

    @Override
    public void afterConnectionEstablished(WebSocketSession session) throws Exception {
        logger.debug("connect to the websocket success......");

        String clientId = (String) session.getAttributes().get(Constant.WEBSOCKET_USER_KEY);

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
        if (session.getAttributes().get(Constant.WEBSOCKET_USER_KEY) != null) {
        	String clientId = (String) session.getAttributes().get(Constant.WEBSOCKET_USER_KEY);
            if (clientId != null && scopeBean.getOnlines().get(clientId) != null) {
                scopeBean.getOnlines().remove(clientId);
            }
        }
    }

    @Override
    public void afterConnectionClosed(WebSocketSession session, CloseStatus closeStatus) throws Exception {
        logger.debug("websocket connection closed......");
        String clientId = (String) session.getAttributes().get(Constant.WEBSOCKET_USER_KEY);
        if (clientId != null && scopeBean.getOnlines().get(clientId) != null) {
        	scopeBean.getOnlines().remove(clientId);
        }
    }

    @Override
    public boolean supportsPartialMessages() {
        return false;
    }
}
