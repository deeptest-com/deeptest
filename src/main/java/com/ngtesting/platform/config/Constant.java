package com.ngtesting.platform.config;

import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class Constant {
    public static final String HTTP_SESSION_USER_PROFILE = "http_session_user_profile";
    public static final String HTTP_SESSION_USER_PERMISSION = "http_session_user_permission";

    // 配置项，初始化在 PropertyConfig.processProperties()
    public static String WORK_DIR;
    public static String API_BASE;

    // WEB根，初始化在 SystemInterceptor.preHandle()
    public static String WEB_ROOT = null;

    // 跨域白名单
    public static final List<String> CLIENT_URL_LIST = Arrays.asList(
            "http://localhost:4200",
            "http://lab.dui.ai/",
            "http://116.62.17.31");
    public static final String API_PATH = "/api/v1/";
    public static final String API_PATH_CLIENT = API_PATH + "client/";
    public static final String API_PATH_ADMIN = API_PATH + "admin/";

    public static final String API_PACKAGE_FOR_CLIENT = "com.ngtesting.platform.action";

    // 上传目录
    public static final String FTP_UPLOAD_DIR = "upload/";

    public static final int PAGE_SIZE = 20;

    public static enum RespCode {
        SUCCESS(1),

        BIZ_FAIL(101),
        BIZ_FAIL_2(102),
        INTERFACE_FAIL(-10),
        NOT_LOGIN(-100),
        AUTH_FAIL(-110),

        RELOAD(100);

        private RespCode(int code) {
            this.code = code;
        }

        private int code;

        public int getCode() {
            return code;
        }
    }

    public static Map<String, String> ExeStatus = new HashMap() {{
        put("pass", "成功");
        put("fail", "失败");
        put("block", "阻塞");
        put("untest", "未执行");
    }};

    public static Map<String, String> JenkinsTask = new HashMap() {{
        put("asr", "ngtesting-asr");
    }};
}
