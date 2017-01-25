package cn.linkr.events.action.client;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;

import cn.linkr.events.service.ChatService;
import cn.linkr.events.util.Constant;


@Controller
@RequestMapping(Constant.API_PATH_CLIENT + "chat/")
public class ChatAction extends BaseAction {
	@Autowired
	ChatService chatService;
	
	
}
