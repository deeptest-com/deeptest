package cn.mobiu.events.service;


import java.util.List;

import cn.mobiu.events.entity.EvtBizcard;
import cn.mobiu.events.vo.BizcardVo;

public interface BizcardService extends BaseService {

	EvtBizcard getMine(Long id);

	List<EvtBizcard> listByEvent(Long eventId, Long clientId);

	EvtBizcard getDetail(Long bizcardId, Long eventId, Long clientId);

	List<BizcardVo> genVos(List<EvtBizcard> pos);
	BizcardVo genVo(EvtBizcard po);

}
