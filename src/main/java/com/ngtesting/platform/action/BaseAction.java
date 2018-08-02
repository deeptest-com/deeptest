package com.ngtesting.platform.action;

public class BaseAction {

//	public TstUser genRequest(HttpServletRequest request, JSONObject json) {
//		TstUser TstUser = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
//
//		if (json.getInteger("orgId") == null) {
//			json.put("orgId", TstUser.getDefaultOrgId());
//		}
//		if (json.getInteger("projectId") == null) {
//			json.put("projectId", TstUser.getDefaultPrjId());
//		}
//
//		return TstUser;
//	}

//	public boolean parameIsEmpty(String... params) {
//		for (String p : params) {
//			if (StringUtils.isEmpty(p)) {
//				return true;
//			}
//		}
//
//		return false;
//	}
//
//	public Map<String, Object> paramError() {
//		Map<String, Object> ret = new HashMap<String, Object>();
//		ret.put("code", RespCode.INTERFACE_FAIL.getCode());
//		ret.put("msg", "parameters error");
//		return ret;
//	}
//
//	public boolean pagingParamError(Map<String, String> json) {
//		if (json.getDetail("startIndex") == null || json.getDetail("pageSize") == null) {
//			return true;
//		}
//
//		return false;
//	}
//
//	public Map<String, Object> parameterError() {
//		Map<String, Object> ret = new HashMap<String, Object>();
//
//		ret.put("code", RespCode.INTERFACE_FAIL.getCode());
//		ret.put("msg", "parameter error");
//		return ret;
//	}

}
