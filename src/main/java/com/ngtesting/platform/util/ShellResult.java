package com.ngtesting.platform.util;

import java.util.HashMap;
import java.util.Map;

public class ShellResult {
    public static final int COMMAND_EXIT_VALUE_SUCCESS = 1;
    public static final int COMMAND_EXIT_VALUE_EXCEPTION = -101;
    public static final int COMMAND_EXIT_VALUE_TIMEOUT = -102;
    public static final int COMMAND_EXIT_VALUE_ERROR_IN_OUTPUT = -103;

    private Map<String, Object> regexMap = new HashMap<String, Object>();

    private int exitValue = COMMAND_EXIT_VALUE_SUCCESS;
    private String errorOutput;
    private String infoOutput;

    public String getErrorOutput() {
        return errorOutput;
    }

    public void setErrorOutput(String errorOutput) {
        this.errorOutput = errorOutput;
    }

    public int getExitValue() {
        return exitValue;
    }

    public void setExitValue(int exitValue) {
        this.exitValue = exitValue;
    }

    public String getInfoOutput() {
        return infoOutput;
    }

    public void setInfoOutput(String infoOutput) {
        this.infoOutput = infoOutput;
    }

    public Map<String, Object> getRegexMap() {
        return regexMap;
    }

    public void setRegexMap(Map<String, Object> regexMap) {
        this.regexMap = regexMap;
    }

}
