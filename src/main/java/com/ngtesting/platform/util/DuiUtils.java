package com.ngtesting.platform.util;

import org.apache.commons.collections.map.HashedMap;

import java.util.Map;

public class DuiUtils {

    public static String HResultsCall(String labPath, String recPath, String resultPath){
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
        cmd += " | tee " + resultPath;
        System.out.println(cmd);

        Map<String, String> regex = new HashedMap();
        regex.put("output", "(HTK Results Analysis at .*?)\\s*|$");
        ShellResult res = ShellUtil.exec(cmd, regex);

        System.out.println(res.getRegexMap().get("output"));

        return res.getRegexMap().get("output") != null? res.getRegexMap().get("output").toString(): "";
    }

}
