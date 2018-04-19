package com.ngtesting.platform.service;

import com.ngtesting.platform.vo.AiTestTaskVo;

public interface JenkinsService extends BaseService {

    String genRunJsonStr(AiTestTaskVo task);

    String execute(AiTestTaskVo vo);

}
