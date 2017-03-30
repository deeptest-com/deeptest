package com.ngtesting.platform.util;

import java.util.Arrays;
import java.util.List;

public final class Constant {
    // 配置项，初始化在 PropertyConfig.processProperties()
    public static String WORK_DIR;
    public static String API_BASE;

    // WEB根，初始化在 SystemInterceptor.preHandle()
    public static String WEB_ROOT = null;

    // 跨域白名单
    public static final List<String> CLIENT_URL_LIST = Arrays.asList("http://localhost:3000");
    public static final String API_PATH_CLIENT = "api/client/v1/";
    
    public static final String API_PACKAGE_FOR_CLIENT = "com.ngtesting.platform.action";

    // 上传目录
    public static final String FTP_UPLOAD_DIR = "upload/";

    public static String GetUploadDir() {
        return Constant.WORK_DIR + FTP_UPLOAD_DIR;
    }

    public static final int PAGE_SIZE = 10;

    public static enum RespCode {
        SUCCESS(1), BIZ_FAIL(101), BIZ_FAIL_2(102), INTERFACE_FAIL(-10), NOT_LOGIN(-100);

        private RespCode(int code) {
            this.code = code;
        }

        private int code;

        public int getCode() {
            return code;
        }
    }
    
    public static final String HTTP_SESSION_USER_KEY = "http_session_user";
    
    public static final String WEBSOCKET_CLIENT_KEY = "clientId";
    public static final String WEBSOCKET_EVENT_KEY = "eventId";
    public static final String WEBSOCKET_TIMESNAP = "websocket_timesnap";
    
    public static final String WEBSCOKET_OPT_ENTER_CHAT_ROOM = "enter_chat_room";
	public static final String WEBSCOKET_OPT_CHAT = "chat";
	
    public static enum TreeNodeType {
    	root("root"),
        branch("branch"),
        leaf("leaf");

        private TreeNodeType(String textVal) {
            this.textVal = textVal;
        }

        private String textVal;

        public String toString() {
            return textVal;
        }
    }
	
}
