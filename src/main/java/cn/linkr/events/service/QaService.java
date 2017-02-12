package cn.linkr.events.service;

import java.util.List;

import cn.linkr.events.entity.EvtQa;
import cn.linkr.events.vo.QaVo;



public interface QaService extends BaseService {

	List<EvtQa> list(Long eventId, Long clientId);

	void save(Long eventId, Long clientId, String content);

	List<QaVo> genVos(List<EvtQa> pos);

	QaVo genVo(EvtQa po);


}
