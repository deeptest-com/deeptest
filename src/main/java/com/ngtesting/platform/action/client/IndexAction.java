package com.ngtesting.platform.action.client;

import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/")
public class IndexAction {
	@RequestMapping("/")
	public String index()  {
		return "forward:index.html";
	}
}
