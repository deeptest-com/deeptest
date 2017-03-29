package com.ngtesting.platform.vo;

/**
 * 心跳信息
 *
 * @author xuxiang
 * @version $Id$
 * @see
 */
public class HeartBeat {

    private Long deviceId = 0L;

    private Long userId = 0L;

    private Long lastUpdate;

    private Long logId;

    public Long getDeviceId() {
        return deviceId;
    }

    public void setDeviceId(Long deviceId) {
        this.deviceId = deviceId;
    }

    public Long getUserId() {
        return userId;
    }

    public void setUserId(Long userId) {
        this.userId = userId;
    }

    public Long getLastUpdate() {
        return lastUpdate;
    }

    public void setLastUpdate(Long lastUpdate) {
        this.lastUpdate = lastUpdate;
    }

    public Long getLogId() {
        return logId;
    }

    public void setLogId(Long logId) {
        this.logId = logId;
    }

}
