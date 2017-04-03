package com.ngtesting.platform.util;

import java.io.BufferedReader;
import java.io.IOException;

import javax.servlet.http.HttpServletRequest;

import com.alibaba.fastjson.JSONObject;

public class HttpUtils {

	public JSONObject reqJson(HttpServletRequest request) { // for wechat
		StringBuffer jb = new StringBuffer();
		String line = null;
		JSONObject jsonObject = null;
		BufferedReader reader;
		try {
			reader = request.getReader();
			while ((line = reader.readLine()) != null) {
				jb.append(line);
			}
			jsonObject = JSONObject.parseObject(jb.toString());
		} catch (IOException e) {
			e.printStackTrace();
		}
		return jsonObject;
	}
    
}
