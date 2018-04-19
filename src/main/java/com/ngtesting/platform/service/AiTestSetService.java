package com.ngtesting.platform.service;

import com.ngtesting.platform.entity.AiTestSet;
import com.ngtesting.platform.vo.AiTestSetVo;

import java.util.List;

public interface AiTestSetService extends BaseService {
    List<AiTestSetVo> listTestSetVo(Long projectId);

    List<AiTestSetVo> genVos(List<AiTestSet> pos);

    AiTestSetVo genVo(AiTestSet po);
}
