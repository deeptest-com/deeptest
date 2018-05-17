package com.ngtesting.platform.action;

import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.util.AuthPassport;
import org.springframework.core.io.FileSystemResource;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.ResponseBody;

import javax.servlet.http.HttpServletRequest;

@Controller
@RequestMapping(Constant.API_PATH_CLIENT + "test/")
public class TestAction extends BaseAction {

	@AuthPassport(validate = false)
	@RequestMapping(value = "stream", method = RequestMethod.GET)
	@ResponseBody
	public FileSystemResource stream(HttpServletRequest request) {
		return new FileSystemResource("/Users/aaron/Downloads/test.mkv");
	}

}
