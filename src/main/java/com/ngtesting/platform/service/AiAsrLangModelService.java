package com.ngtesting.platform.service;

import com.ngtesting.platform.entity.AiAsrLangModel;
import com.ngtesting.platform.vo.AiAsrLangModelVo;

import java.util.List;

public interface AiAsrLangModelService extends BaseService {

    List<AiAsrLangModelVo> listAsrLangModelVo(Long projectId);

    List<AiAsrLangModelVo> genVos(List<AiAsrLangModel> pos);

    AiAsrLangModelVo genVo(AiAsrLangModel po);
}
