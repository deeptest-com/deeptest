package com.ngtesting.platform.vo;

import java.io.Serializable;

public class JsonResult implements Serializable {

    private static final long serialVersionUID = 7311737007568787224L;

    private Integer code;

    private String result;

    private String type;

    public Integer getCode() {
        return code;
    }

    public void setCode(Integer code) {
        this.code = code;
    }

    public String getResult() {
        return result;
    }

    public void setResult(String result) {
        this.result = result;
    }

    public String getType() {
        return type;
    }

    public void setType(String type) {
        this.type = type;
    }

}
