package com.ngtesting.platform.action.client;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;

@Controller
@RequestMapping("/")
public class IndexAction {
	@RequestMapping("/")
	public String index()  {
		return "forward:index.html";
	}
}
