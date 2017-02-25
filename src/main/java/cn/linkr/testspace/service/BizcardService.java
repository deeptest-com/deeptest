package cn.linkr.testspace.service;


import java.util.List;

import cn.linkr.testspace.entity.EvtBizcard;
import cn.linkr.testspace.vo.BizcardVo;

public interface BizcardService extends BaseService {

	EvtBizcard getMine(Long id);

	List<EvtBizcard> listByEvent(Long eventId, Long clientId);

	EvtBizcard getDetail(Long bizcardId, Long eventId, Long clientId);

	List<BizcardVo> genVos(List<EvtBizcard> pos);
	BizcardVo genVo(EvtBizcard po);

}
