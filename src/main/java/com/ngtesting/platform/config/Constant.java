package com.ngtesting.platform.config;

import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class Constant {
    public static final String HTTP_SESSION_USER_PROFILE = "http_session_user_profile";

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

    public enum MsgType {
        create("create", "创建"),
        update("update", "更新"),
        update_case("update_case", "更新用例为");

        MsgType(String code, String msg) {
            this.code = code;
            this.msg = msg;
        }

        public String code;
        public String msg;
        public String toString() {
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

    public enum CaseAct {
        create("create", "创建"),
        rename("rename", "改名"),
        update("update", "更新"),
        move("move", "移动"),
        copy("copy", "复制"),
        delete("delete", "删除"),

        attachment_upload("attachment_upload", "上传附件"),
        attachment_delete("attachment_delete", "删除附件"),

        comments_add("comments_add", "新增注释"),
        comments_update("comments_update", "修改注释"),
        comments_delete("comments_delete", "删除注释"),

        exe_result("exe_result", "标注执行结果");

        CaseAct(String code, String msg) {
            this.code = code;
            this.msg = msg;
        }

        public String code;
        public String msg;
        public String toString() {
            return code;
        }
    }
}
