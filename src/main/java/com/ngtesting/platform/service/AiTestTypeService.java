package com.ngtesting.platform.service;

import com.ngtesting.platform.entity.AiTestType;
import com.ngtesting.platform.vo.AiTestTypeVo;

import java.util.List;

public interface AiTestTypeService extends BaseService {
    List<AiTestTypeVo> listTestTypeVo(Long projectId);

    List<AiTestTypeVo> genVos(List<AiTestType> pos);

    AiTestTypeVo genVo(AiTestType po);
}
