package com.ngtesting.platform.action.client;

import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.WelcomeService;
import com.ngtesting.platform.utils.AuthPassport;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.ResponseBody;

import javax.servlet.http.HttpServletRequest;
import java.util.HashMap;
import java.util.List;
import java.util.Map;


@Controller
@RequestMapping(Constant.API_PATH_CLIENT + "/")
public class WelcomeAction extends BaseAction {

    @Autowired
    WelcomeService welcomeService;

	@AuthPassport(validate=false)
	@RequestMapping(value = "test", method = RequestMethod.GET)
	@ResponseBody
	public Map<String, Object> test(HttpServletRequest request) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

        List ls = welcomeService.test();

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		ret.put("data", ls);
		ret.put("msg", ls.size() > 0? "pass": "fail");

		return ret;
	}

}
