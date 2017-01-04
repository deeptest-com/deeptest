package cn.mobiu.events.util;

import java.io.File;
import java.util.UUID;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class CommonUtil {
    private static final Logger logger = LoggerFactory
            .getLogger(CommonUtil.class);

    public static String RandomNameFromOldName(String name) {
        String postfix = name.substring(name.lastIndexOf(".") + 1);
        return RandomName(postfix);
    }

    public static String RandomName(String extName) {
        return UUID.randomUUID().toString() + "." + extName;
    }

    public static void CreateDirIfNeeded(String path) {
        File dir = new File(path);
        if (!dir.exists()) {
            dir.mkdirs();
        }
    }

    public static String GetFileName(String remotePath) {
        return remotePath.substring(remotePath.lastIndexOf("/"), remotePath.length());
    }
}