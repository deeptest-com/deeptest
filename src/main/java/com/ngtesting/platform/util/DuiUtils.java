package com.ngtesting.platform.util;

import java.util.HashMap;

public class DuiUtils {

    public static void HResultsCall(String labPath, String recPath, String resultPath){
        String hresultsPath = "HResults";

        String os = System.getProperty("os.name").toLowerCase();
        String cmd;
        if (os.indexOf("win") > -1) {
            cmd = hresultsPath + " -t -e "+"\"一\" \"幺\""+" -A -D -T 1 -h -I " + labPath + " NUL "
                    + recPath;
        } else {
            cmd = hresultsPath + " -t -e "+"\"一\" \"幺\""+" -A -D -T 1 -h -I " + labPath + " /dev/null "
                    + recPath;
        }
        cmd += " > " + resultPath;
        System.out.println(cmd);

        ShellUtil.exec(cmd, new HashMap<String, String>());

    }

}
