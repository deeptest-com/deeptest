package com.ngtesting.platform.service;

import com.ngtesting.platform.model.AiAudioType;

import java.util.List;

public interface AiAudioTypeService extends BaseService {

	List<AiAudioType> listAudioTypeVo(Long projectId);

}
