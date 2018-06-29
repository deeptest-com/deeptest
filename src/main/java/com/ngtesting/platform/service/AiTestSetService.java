package com.ngtesting.platform.service;


import com.ngtesting.platform.model.AiTestSet;

import java.util.List;

public interface AiTestSetService extends BaseService {
    List<AiTestSet> listTestSetVo(Integer projectId);
}
