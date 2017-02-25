package cn.linkr.testspace.action;

import java.io.BufferedReader;
import java.io.IOException;
import java.util.HashMap;
import java.util.Map;

import javax.servlet.http.HttpServletRequest;
import javax.validation.Validator;

import org.apache.commons.lang.StringUtils;
import org.springframework.beans.factory.annotation.Autowired;

import cn.linkr.testspace.util.Constant.RespCode;

import com.alibaba.fastjson.JSONObject;

public class BaseAction {

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

	public boolean parameIsEmpty(String... params) {
		for (String p : params) {
			if (StringUtils.isEmpty(p)) {
				return true;
			}
		}

		return false;
	}

	public Map<String, Object> paramError() {
		Map<String, Object> ret = new HashMap<String, Object>();
		ret.put("code", RespCode.INTERFACE_FAIL.getCode());
		ret.put("msg", "parameters error");
		return ret;
	}

	public boolean pagingParamError(Map<String, String> json) {
		if (json.get("startIndex") == null || json.get("pageSize") == null) {
			return true;
		}

		return false;
	}

	public Map<String, Object> parameterError() {
		Map<String, Object> ret = new HashMap<String, Object>();

		ret.put("code", RespCode.INTERFACE_FAIL.getCode());
		ret.put("msg", "parameter error");
		return ret;
	}

}
