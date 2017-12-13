package com.ngtesting.platform.job;

import com.ngtesting.platform.entity.TestAlert;
import com.ngtesting.platform.entity.TestMsg;
import com.ngtesting.platform.service.AlertService;
import com.ngtesting.platform.service.MailService;
import com.ngtesting.platform.service.MsgService;
import com.ngtesting.platform.util.DateUtils;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Component;

import java.util.Date;
import java.util.List;

@Component
public class MsgJob {
	Log logger = LogFactory.getLog(MsgJob.class);

	@Autowired
	private MsgService msgService;

	@Autowired
	private AlertService alertService;

    @Autowired
    private MailService mailService;

    @Scheduled(cron="0 0/1 * * * ?") // 1分钟测试
	// @Scheduled(cron="0 0 0/1 * * ?") // 1小时
    private void sendMsg() {
    	String time = DateUtils.FormatDate(new Date(), "yyyy-MM-dd HH");
    	System.out.println("开始定时任务-发送消息 @" + time);


        List<TestMsg> msgs = msgService.list();

        List<TestAlert> alerts = alertService.list();



    }
}
