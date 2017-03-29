package com.ngtesting.platform.service;

import java.util.List;

import com.ngtesting.platform.entity.EvtQa;
import com.ngtesting.platform.vo.QaVo;



public interface QaService extends BaseService {

	List<EvtQa> list(Long eventId, Long clientId);

	void save(Long eventId, Long clientId, String content);

	List<QaVo> genVos(List<EvtQa> pos);

	QaVo genVo(EvtQa po);


}
