package cn.mobiu.events.action.client;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;

import cn.mobiu.events.constants.Constant;
import cn.mobiu.events.service.ChatService;


@Controller
@RequestMapping(Constant.API_PATH_CLIENT + "chat/")
public class ChatAction extends BaseAction {
	@Autowired
	ChatService chatService;
	
	
}
