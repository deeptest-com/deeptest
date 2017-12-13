package com.ngtesting.platform.bean;

import org.springframework.context.annotation.Scope;
import org.springframework.stereotype.Component;
import org.springframework.web.socket.TextMessage;
import org.springframework.web.socket.WebSocketSession;

import java.io.IOException;
import java.util.Map;
import java.util.Set;
import java.util.concurrent.ConcurrentHashMap;

@Component
@Scope("singleton")
public class ApplicationScopeBean {

    // 在线用户，clientId -> session
    private final Map<String, WebSocketSession> onlines = new ConcurrentHashMap<String, WebSocketSession>();

    // 聊天室用户，eventId -> clientIds
    private final Map<String, Set<String>> chatroom = new ConcurrentHashMap<String, Set<String>>();

    public Map<String, WebSocketSession> getOnlines() {
        return onlines;
    }

    public Map<String, Set<String>> getChatroom() {
        return chatroom;
    }

    // 给所有用户发送消息
    public void sendMessageToAllClient(TextMessage message) {
        for (String clientId : getOnlines().keySet()) {
        	WebSocketSession session = getOnlines().get(clientId);

            if (session.isOpen()) {
                try {
                    session.sendMessage(message);
                } catch (IOException e) {
                    e.printStackTrace();
                }
            }
        }
    }

    // 给聊天的用户发送消息
    public void sendMessageToChatroom(Long eventId, TextMessage message) {
        for (String clientId : getChatroom().get(eventId)) {
        	WebSocketSession session = getOnlines().get(clientId);

        	// 离线，从聊天室移除
        	if (session == null) {
        		getChatroom().get(eventId).remove(clientId);
        		return;
        	}

            if (session.isOpen()) {
                try {
                    session.sendMessage(message);
                } catch (IOException e) {
                    e.printStackTrace();
                }
            }
        }
    }

    // 给某个用户发送消息
    public void sendMessageToClient(String clientId, TextMessage message) {
    	WebSocketSession session = getOnlines().get(clientId);

        if (session == null) {
            return;
        }

        if (session.isOpen()) {
            try {
                session.sendMessage(message);
            } catch (IOException e) {
                e.printStackTrace();
            }
        }
    }

}
