package com.ngtesting.platform.utils;

import com.ngtesting.platform.config.Constant;

import javax.servlet.http.Cookie;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;
import java.util.Map;

public class WebUtils {

    /**
     * 响应正常状态值200
     */
    private static final int RESPONSE_STATUS_NORMAL = 200;

    /**
     * COOKIE 默认过期时间
     */
    private static final long COOKIE_DEFAULT_EXPIRE_TIME = 1000L;


    /**
     * COOKIE MAX_AGE
     */
    private static final int COOKIE_MAX_AGE = 86400;

    public static void InitWebContext(HttpServletRequest request, String workDir) {
        if (Constant.WEB_ROOT == null) {
            Constant.WEB_ROOT = request.getScheme() + "://" + request.getServerName() +
                    (request.getServerPort() != 80? ":" + request.getServerPort() : "")
                    + request.getContextPath() + "/";
        }
        if (Constant.WORK_DIR == null) {
            Constant.WORK_DIR = workDir;
        }
    }

    /**
     * 获得系统servelet路径
     *
     * @param httpServletRequest 请求
     * @return 返回路径字符串
     */
    public static String getServetPath(HttpServletRequest httpServletRequest) {
        String serverPath = httpServletRequest.getScheme() + "://"
                + httpServletRequest.getServerName() + ":" + httpServletRequest.getServerPort()
                + httpServletRequest.getContextPath();
        return serverPath;
    }

    /**
     * render
     *
     * @param response    响应
     * @param text        文本
     * @param contentType render 内容的类型
     * @return 返回字符串
     */
    public static String render(HttpServletResponse response, String text, String contentType) {
        try {
            response.setContentType(contentType);
            response.getWriter().write(text);
        } catch (IOException localIOException) {
            localIOException.printStackTrace();
        }
        return null;
    }

    /**
     * render 文本
     *
     * @param response 响应
     * @param text     文本
     * @return 返回字符串
     */
    public static String renderText(HttpServletResponse response, String text) {
        return render(response, text, "text/plain;charset=UTF-8");
    }

    /**
     * render html
     *
     * @param response 响应
     * @param html     html
     * @return 返回字符串
     */
    public static String renderHtml(HttpServletResponse response, String html) {
        return render(response, html, "text/html;charset=UTF-8");
    }

    /**
     * render xml
     *
     * @param response 响应
     * @param xml      xml
     * @return 返回字符串
     */
    public static String renderXML(HttpServletResponse response, String xml) {
        return render(response, xml, "text/xml;charset=UTF-8");
    }

    /**
     * render json
     *
     * @param response 响应
     * @param json     json
     * @return 返回字符串
     */
    public static String renderJson(HttpServletResponse response, String json) {
        return render(response, json, "application/json;charset=UTF-8");
    }

    /**
     * render js
     *
     * @param response 响应
     * @param jsText   jsText
     * @return 返回字符串
     */
    public static String renderJs(HttpServletResponse response, String jsText) {
        return render(response, jsText, "application/x-javascript;charset=utf-8");
    }

    /**
     * render js 到 html上
     *
     * @param response 响应
     * @param jsText   jsText
     * @return 返回字符串
     */
    public static String renderJsToHtml(HttpServletResponse response, String jsText) {
        WebUtils.setNoCacheHeader(response);
        return render(response, "<script type='text/javascript'>" + jsText + "</script>", "text/html;charset=utf-8");
    }

    /**
     * 设置 ExpiresHeader
     *
     * @param response       响应
     * @param expiresSeconds expiresSeconds
     */
    public static void setExpiresHeader(HttpServletResponse response, long expiresSeconds) {
        response.setDateHeader("Expires", System.currentTimeMillis() + expiresSeconds * COOKIE_DEFAULT_EXPIRE_TIME);
        response.setHeader("Cache-Control", "max-age=" + expiresSeconds);
    }

    /**
     * 设置 NoCacheHeader
     *
     * @param response 响应
     */
    public static void setNoCacheHeader(HttpServletResponse response) {
        response.setDateHeader("Expires", 0L);
        response.setHeader("Cache-Control", "no-cache");
        response.setHeader("Pragma", "no-cache");
        response.setHeader("Cache-control", "private, no-cache, no-store");
        response.setHeader("Expires", "0");
        response.setStatus(RESPONSE_STATUS_NORMAL);
    }

    /**
     * 设置 LastModifiedHeader
     *
     * @param response         响应
     * @param lastModifiedDate 最后修改时间
     */
    public static void setLastModifiedHeader(HttpServletResponse response, long lastModifiedDate) {
        response.setDateHeader("Last-Modified", lastModifiedDate);
    }

    /**
     * 设置 DownloadableHeader
     *
     * @param response 响应
     * @param fileName 文件名
     */
    public static void setDownloadableHeader(HttpServletResponse response, String fileName) {
        response.setHeader("Content-Disposition", "attachment; filename=\"" + fileName + "\"");
    }

    /**
     * 获得指定前缀的请求参数
     *
     * @param request 请求
     * @param prefix  前缀
     * @return 返回map
     */
    public static Map<String, Object> getParametersStartingWith(HttpServletRequest request, String prefix) {
        return org.springframework.web.util.WebUtils.getParametersStartingWith(request, prefix);
    }

    /**
     * 设置cookie
     *
     * @param cKey     cookies key
     * @param value    值
     * @param response 响应
     */
    public static void setCookie(String cKey, String value, HttpServletResponse response) {
        Cookie c1 = new Cookie(cKey, value);
        c1.setMaxAge(COOKIE_MAX_AGE);
        response.addCookie(c1);
    }

    /**
     * 获得指定cookie
     *
     * @param cKey    cookies key
     * @param request 请求
     * @return 返回字符串
     */
    public static String getCookie(String cKey, HttpServletRequest request) {
        Cookie[] cookies = request.getCookies();
        if (cookies != null) {
            for (Cookie cookie : cookies) {
                if (cookie.getName().equals(cKey)) {
                    return cookie.getValue();
                }
            }
        }
        return null;
    }


    public static HttpServletResponse AddCorsSupport(HttpServletResponse res, String ref) {
        res.addHeader("Access-Control-Allow-Origin", ref);
        res.addHeader("Access-Control-Allow-Credentials", "true");

        res.addHeader("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, token, Authorization");

        return res;
    }
}
