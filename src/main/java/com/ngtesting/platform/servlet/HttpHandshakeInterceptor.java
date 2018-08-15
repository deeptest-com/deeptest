package com.ngtesting.platform.servlet;

import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.config.WsConstant;
import com.ngtesting.platform.model.TstUser;
import org.springframework.http.server.ServerHttpRequest;
import org.springframework.http.server.ServerHttpResponse;
import org.springframework.http.server.ServletServerHttpRequest;
import org.springframework.stereotype.Component;
import org.springframework.web.socket.WebSocketHandler;
import org.springframework.web.socket.server.HandshakeInterceptor;

import javax.servlet.http.HttpSession;
import java.util.Map;

@Component
public class HttpHandshakeInterceptor implements HandshakeInterceptor {

    @Override
    public boolean beforeHandshake(ServerHttpRequest request, ServerHttpResponse response, WebSocketHandler wsHandler,
                                   Map attributes) throws Exception {

        if (request instanceof ServletServerHttpRequest) {
            ServletServerHttpRequest servletRequest = (ServletServerHttpRequest) request;
            HttpSession httpSession = servletRequest.getServletRequest().getSession(true);

            String test = (String) httpSession.getAttribute("TEST");

            TstUser user = null;
            if (httpSession.getAttribute(Constant.HTTP_SESSION_USER_PROFILE) != null) {
                user = (TstUser) httpSession.getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
            }

            if (user != null) {
                String userId = user.getId().toString();
                attributes.put(WsConstant.WS_USER_KEY, user);

                return true;
            }
        }

        return false;

    }

    public void afterHandshake(ServerHttpRequest request, ServerHttpResponse response, WebSocketHandler wsHandler,
                               Exception ex) {
    }
}
