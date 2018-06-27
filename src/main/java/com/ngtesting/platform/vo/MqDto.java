package com.ngtesting.platform.vo;

import java.io.Serializable;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class MqDto implements Serializable {
    private static final long serialVersionUID = 9041855715915314046L;

    public static final String INSTRUCTION_SCREENSHOT = "screenshot";

    public static final String INSTRUCTION_INSTALL = "install";

    public static final String INSTRUCTION_REBOOT = "reboot";

    public static final String INSTRUCTION_UNINSTALL = "uninstall";

    public static final String INSTRUCTION_UNINSTALLBATCH = "uninstallBatch";

    public static final String INSTRUCTION_PUSHFILE = "pushfile";

    public static final String INSTRUCTION_STARTAPP = "startApp";

    public static final String INSTRUCTION_STOPAPP = "stopApp";

    public static final String INSTRUCTION_TEST = "test";

    public static final String MONITOR_VNC_RESTART = "vnc_restart";

    String msg;

    String sessionId;

    Long transactionId;
    String instruction;

    Long mechineId;
    Long mechineIp;

    String deviceId;
    String deviceIp;
    String deviceType;

    String packagePath;
    String packageName;

    String imagePath;

    Long requestTime;
    Long responseTime;

    String token;

    String hostIp;

    private Integer code = 1;

    private List<String> output = new ArrayList<String>();

    Map<String, Object> result = new HashMap<String, Object>();

    public Long getTransactionId() {
        return transactionId;
    }

    public void setTransactionId(Long transactionId) {
        this.transactionId = transactionId;
    }

    public String getInstruction() {
        return instruction;
    }

    public void setInstruction(String instruction) {
        this.instruction = instruction;
    }

    public Long getMechineId() {
        return mechineId;
    }

    public void setMechineId(Long mechineId) {
        this.mechineId = mechineId;
    }

    public Long getMechineIp() {
        return mechineIp;
    }

    public void setMechineIp(Long mechineIp) {
        this.mechineIp = mechineIp;
    }

    public String getDeviceId() {
        return deviceId;
    }

    public void setDeviceId(String deviceId) {
        this.deviceId = deviceId;
    }

    public String getDeviceIp() {
        return deviceIp;
    }

    public void setDeviceIp(String deviceIp) {
        this.deviceIp = deviceIp;
    }

    public String getDeviceType() {
        return deviceType;
    }

    public void setDeviceType(String deviceType) {
        this.deviceType = deviceType;
    }

    public String getPackagePath() {
        return packagePath;
    }

    public void setPackagePath(String packagePath) {
        this.packagePath = packagePath;
    }

    public String getPackageName() {
        return packageName;
    }

    public void setPackageName(String packageName) {
        this.packageName = packageName;
    }

    public String getImagePath() {
        return imagePath;
    }

    public void setImagePath(String imagePath) {
        this.imagePath = imagePath;
    }

    public Long getRequestTime() {
        return requestTime;
    }

    public void setRequestTime(Long requestTime) {
        this.requestTime = requestTime;
    }

    public Long getResponseTime() {
        return responseTime;
    }

    public void setResponseTime(Long responseTime) {
        this.responseTime = responseTime;
    }

    public String getToken() {
        return token;
    }

    public void setToken(String token) {
        this.token = token;
    }

    public Integer getCode() {
        return code;
    }

    public void setCode(Integer code) {
        this.code = code;
    }

    public List<String> getOutput() {
        return output;
    }

    public void setOutput(List<String> output) {
        this.output = output;
    }

    public String getMsg() {
        return msg;
    }

    public void setMsg(String msg) {
        this.msg = msg;
    }

    public String getSessionId() {
        return sessionId;
    }

    public void setSessionId(String sessionId) {
        this.sessionId = sessionId;
    }

    public boolean isSuccess() {
        return code > 0;
    }

    public String getHostIp() {
        return hostIp;
    }

    public void setHostIp(String hostIp) {
        this.hostIp = hostIp;
    }

    public Map<String, Object> getResult() {
        return result;
    }

    public void setResult(Map<String, Object> result) {
        this.result = result;
    }
}

