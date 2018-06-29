package com.ngtesting.platform.service.inf;

import com.ngtesting.platform.model.AiAsrLangModel;

import java.util.List;

public interface AiAsrLangModelService extends BaseService {

    List<AiAsrLangModel> listAsrLangModelVo(Long projectId);

    List<AiAsrLangModel> genVos(List<AiAsrLangModel> pos);

    AiAsrLangModel genVo(AiAsrLangModel po);
}
