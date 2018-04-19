package com.ngtesting.platform.service;

import com.ngtesting.platform.entity.AiAudioType;
import com.ngtesting.platform.vo.AiAudioTypeVo;

import java.util.List;

public interface AiAudioTypeService extends BaseService {

	List<AiAudioTypeVo> listAudioTypeVo(Long projectId);

	List<AiAudioTypeVo> genVos(List<AiAudioType> pos);

	AiAudioTypeVo genVo(AiAudioType po);
}
