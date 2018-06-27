package com.ngtesting.platform.utils;

public class Constants {
    public static String SESSION_KEY = "session_key";

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
}
