package com.ngtesting.platform.action;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.bean.websocket.WsFacade;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.model.TstVer;
import com.ngtesting.platform.service.IsuFilterService;
import com.ngtesting.platform.service.IsuTqlService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.ResponseBody;

import javax.servlet.http.HttpServletRequest;
import java.io.UnsupportedEncodingException;
import java.net.URLEncoder;
import java.util.HashMap;
import java.util.List;
import java.util.Map;


@Controller
@RequestMapping(Constant.API_PATH_CLIENT + "tql/")
public class TplAction extends BaseAction {
	@Autowired
	private WsFacade optFacade;

	@Autowired
	IsuTqlService isuTqlService;
	@Autowired
	IsuFilterService isuFilterService;

	@RequestMapping(value = "getFilters", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> getFilters(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

		String tql = json.getString("tql");
		if ("all".equals(tql)) {
            try {
                tql = URLEncoder.encode("project=350 AND status='in_progress' ORDER BY status ASC", "UTF-8"); // sample
            } catch (UnsupportedEncodingException e) {
                e.printStackTrace();
            }
        }

		List<TstVer> ls = isuTqlService.getFilters(tql);

        ret.put("data", ls);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
