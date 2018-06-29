package com.ngtesting.platform.service;

import com.ngtesting.platform.vo.AiRun;
import com.ngtesting.platform.model.AiTestTask;

public interface JenkinsService extends BaseService {

    AiRun genRunVo(AiTestTask task);

    String execute(AiTestTask vo);

}
