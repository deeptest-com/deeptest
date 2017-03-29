package com.ngtesting.platform.service;


import java.util.List;

import com.ngtesting.platform.entity.EvtBizcard;
import com.ngtesting.platform.vo.BizcardVo;

public interface BizcardService extends BaseService {

	EvtBizcard getMine(Long id);

	List<EvtBizcard> listByEvent(Long eventId, Long clientId);

	EvtBizcard getDetail(Long bizcardId, Long eventId, Long clientId);

	List<BizcardVo> genVos(List<EvtBizcard> pos);
	BizcardVo genVo(EvtBizcard po);

}
