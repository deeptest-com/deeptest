package com.ngtesting.platform.ctrl;

import com.ngtesting.platform.vo.Greeting;
import com.ngtesting.platform.vo.HelloMessage;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.messaging.handler.annotation.MessageMapping;
import org.springframework.messaging.handler.annotation.SendTo;
import org.springframework.stereotype.Controller;
import org.springframework.web.util.HtmlUtils;

@Controller
public class GreetingCtrl {

    @Value("${mail.smtp.host}")
    private String mailHost;

    @MessageMapping("/comm")
    @SendTo("/topic/comm")
    public Greeting greeting(HelloMessage message) throws Exception {
        Thread.sleep(1000); // simulated delay
        return new Greeting("Hello, " + mailHost + ", " + HtmlUtils.htmlEscape(message.getName()) + "!");
    }

}
