package com.ngtesting.platform.utils;

import com.alibaba.fastjson.JSONObject;

public class CustomFieldUtil {
    public static Object GetFieldVal(String type, JSONObject json) {
        Object value = null;
        if (type.equals("string") ) {
            value = json.getString("value");
        } else if (type.equals("integer") ) {
            value = json.getInteger("value");
        } else if (type.equals("double") ) {
            value = json.getString("value");
        } else if (type.equals("date") ) {
            value = json.getString("value");
        } else if (type.equals("time") ) {
            value = json.getString("value");
        } else if (type.equals("datetime") ) {
            value = json.getString("value");
        }

        return value;
    }
}
