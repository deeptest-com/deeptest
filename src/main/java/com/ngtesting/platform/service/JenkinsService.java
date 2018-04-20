package com.ngtesting.platform.service;

import com.ngtesting.platform.vo.AiRunVo;
import com.ngtesting.platform.vo.AiTestTaskVo;

public interface JenkinsService extends BaseService {

    AiRunVo genRunVo(AiTestTaskVo task);

    String execute(AiTestTaskVo vo);

}
