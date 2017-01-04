package cn.mobiu.events.filter;

import java.io.IOException;

import javax.servlet.Filter;
import javax.servlet.FilterChain;
import javax.servlet.FilterConfig;
import javax.servlet.ServletException;
import javax.servlet.ServletRequest;
import javax.servlet.ServletResponse;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import cn.mobiu.events.constants.Constant;
import cn.mobiu.events.util.WebUtils;

public class CorsFilter implements Filter {
    @Override
    public void doFilter(ServletRequest requ, ServletResponse resp, FilterChain chain) throws IOException, ServletException {
        HttpServletRequest req = (HttpServletRequest) requ;
        HttpServletResponse res = (HttpServletResponse) resp;

        String referer = req.getHeader("Origin");
        if (Constant.CLIENT_URL_LIST.contains(referer)) {
            res = WebUtils.AddCorsSupport(res, referer);
        }

        chain.doFilter(req, res);
    }

    @Override
    public void init(FilterConfig arg0) throws ServletException {
    }

    @Override
    public void destroy() {
    }
}
