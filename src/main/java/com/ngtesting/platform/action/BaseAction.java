package com.ngtesting.platform.action;

import java.io.BufferedReader;
import java.io.IOException;
import java.util.HashMap;
import java.util.Map;

import javax.servlet.http.HttpServletRequest;

import org.apache.commons.lang.StringUtils;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.util.Constant;
import com.ngtesting.platform.util.Constant.RespCode;
import com.ngtesting.platform.vo.UserVo;

public class BaseAction {
	
	public UserVo genRequest(HttpServletRequest request, JSONObject json) {
		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		
		if (json.getLong("orgId") == null) {
			json.put("orgId", userVo.getDefaultOrgId());
		}
		if (json.getLong("projectId") == null) {
			json.put("projectId", userVo.getDefaultProjectId());
		}

		return userVo;
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
