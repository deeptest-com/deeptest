package com.ngtesting.platform.action;

import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.WelcomeService;
import com.ngtesting.platform.utils.AuthPassport;
import org.apache.shiro.SecurityUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.core.io.FileSystemResource;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

import javax.servlet.http.HttpServletRequest;
import java.util.HashMap;
import java.util.List;
import java.util.Map;


@RestController
@RequestMapping(Constant.API_PATH_CLIENT + "/")
public class WelcomeAction extends BaseAction {

    @Autowired
    WelcomeService welcomeService;

	@AuthPassport(validate=false)
	@RequestMapping(value = "test", method = RequestMethod.GET)

	public Map<String, Object> test(HttpServletRequest request) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();

        List ls = welcomeService.test();

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		ret.put("data", ls);
		ret.put("msg", ls.size() > 0? "pass": "fail");

		return ret;
	}

	@AuthPassport(validate=false)
	@RequestMapping(value = "stream", method = RequestMethod.GET)

	public FileSystemResource stream(HttpServletRequest request) {
		return new FileSystemResource("/Users/aaron/Downloads/test.mkv");
	}

}
