package cn.linkr.testspace.vo;

import java.io.Serializable;

/**
 * json字符串封装工具类
 *
 * @author zhangzhi
 */
public class JsonBean implements Serializable {

    private static final long serialVersionUID = 3462100289096119626L;

    private String type;

    private Long startTime;

    private Long endTime;

    private String number;

    private String content;

    private String longitude;

    private String latitude;

    public JsonBean() {
    }

    public JsonBean(String type) {
        if (type != null) {
            this.type = type;
        }
    }

    public JsonBean(String type, String content) {
        this(type);
        if (content != null) {
            this.content = content;
        }
    }

    public JsonBean(String type, String number, String content) {
        this(type, content);
        if (number != null) {
            this.number = number;
        }
    }

    public String getType() {
        return type;
    }

    public void setType(String type) {
        this.type = type;
    }

    public Long getStartTime() {
        return startTime;
    }

    public void setStartTime(Long startTime) {
        this.startTime = startTime;
    }

    public Long getEndTime() {
        return endTime;
    }

    public void setEndTime(Long endTime) {
        this.endTime = endTime;
    }

    public String getNumber() {
        return number;
    }

    public void setNumber(String number) {
        this.number = number;
    }

    public String getContent() {
        return content;
    }

    public void setContent(String content) {
        this.content = content;
    }

    public String getLongitude() {
        return longitude;
    }

    public void setLongitude(String longitude) {
        this.longitude = longitude;
    }

    public String getLatitude() {
        return latitude;
    }

    public void setLatitude(String latitude) {
        this.latitude = latitude;
    }

}
