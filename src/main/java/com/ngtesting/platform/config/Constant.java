package com.ngtesting.platform.config;

import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public final class Constant {

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
    public static final String API_PATH_CLIENT = "api/client/v1/";

    public static final String API_PACKAGE_FOR_CLIENT = "com.ngtesting.platform.action";

    // 上传目录
    public static final String FTP_UPLOAD_DIR = "upload/";

    public static String GetUploadDir() {
        return Constant.WORK_DIR + FTP_UPLOAD_DIR;
    }

    public static final int PAGE_SIZE = 10;

    public static enum RespCode {
        SUCCESS(1), BIZ_FAIL(101), BIZ_FAIL_2(102), INTERFACE_FAIL(-10), NOT_LOGIN(-100),
        RELOAD(100);

        private RespCode(int code) {
            this.code = code;
        }

        private int code;

        public int getCode() {
            return code;
        }
    }

    public static final String HTTP_SESSION_USER_KEY = "http_session_user";

    public static final String KEY_TESTCASE_DESIGN = "TC-";
    public static final String KEY_TESTCASE_EXE = "TE-";

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

    public enum AlertType {
        run_start("run_start", 1),
        run_end("run_end", 1);

        AlertType(String code, Integer remindDay) {
            this.code = code;
            this.remindDay = remindDay;
        }

        public String code;
        public Integer remindDay;
        public String toString() {
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

        upload_attachment("upload_attachment", "上传附件"),
        delete_attachment("delete_attachment", "删除附件");

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
