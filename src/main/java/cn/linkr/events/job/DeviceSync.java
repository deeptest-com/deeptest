package cn.linkr.events.job;

import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Component;

import cn.linkr.events.service.HttpClientService;

@Component
public class DeviceSync {

    Log logger = LogFactory.getLog(DeviceSync.class);

    @Autowired
    private HttpClientService httpClient;

    @Scheduled(cron = "0 0 3 * * ?")
    //@Scheduled(cron = "0/5 * *  * * ? ")
    private void syncDevice() {

    }
}
