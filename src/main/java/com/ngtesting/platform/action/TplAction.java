package com.ngtesting.platform.action;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.bean.websocket.WsFacade;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.dao.TestVerDao;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.TestVerService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.ResponseBody;

import javax.servlet.http.HttpServletRequest;
import java.util.HashMap;
import java.util.Map;


@Controller
@RequestMapping(Constant.API_PATH_CLIENT + "tql/")
public class TplAction extends BaseAction {
	@Autowired
	private WsFacade optFacade;

	@Autowired
	TestVerService verService;
	@Autowired
	TestVerDao verDao;

	@RequestMapping(value = "getAllFilters", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> getAllFilters(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

		String tql = json.getString("tql");

//		List<TstVer> ls = verService.list(projectId, keywords, disabled);

//        ret.put("data", ls);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
