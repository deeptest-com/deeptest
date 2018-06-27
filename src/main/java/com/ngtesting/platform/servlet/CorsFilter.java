package com.ngtesting.platform.servlet;

import com.ngtesting.platform.utils.WebUtils;

import javax.servlet.*;
import javax.servlet.FilterConfig;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;

public class CorsFilter implements Filter {
    @Override
    public void init(FilterConfig filterConfig) throws ServletException {

    }

    @Override
    public void doFilter(ServletRequest requ, ServletResponse resp, FilterChain chain) throws IOException, ServletException {
        HttpServletRequest req = (HttpServletRequest) requ;
        HttpServletResponse res = (HttpServletResponse) resp;

        String referer = req.getHeader("Origin");
        res = WebUtils.AddCorsSupport(res, referer);

        chain.doFilter(req, res);
    }

    @Override
    public void destroy() {
    }
}
