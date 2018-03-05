package com.ngtesting.platform.servlet;

import com.ngtesting.platform.util.WebUtils;

import javax.servlet.*;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;

public class CorsFilter implements Filter {
    @Override
    public void doFilter(ServletRequest requ, ServletResponse resp, FilterChain chain) throws IOException, ServletException {
        HttpServletRequest req = (HttpServletRequest) requ;
        HttpServletResponse res = (HttpServletResponse) resp;

        String referer = req.getHeader("Origin");
//        if (Constant.CLIENT_URL_LIST.contains(referer)) {
            res = WebUtils.AddCorsSupport(res, referer);
//        }

        chain.doFilter(req, res);
    }

    @Override
    public void init(FilterConfig arg0) throws ServletException {
    }

    @Override
    public void destroy() {
    }
}
