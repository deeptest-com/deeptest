package cn.mobiu.events.service;

import java.util.List;

import cn.mobiu.events.entity.EvtQa;
import cn.mobiu.events.vo.QaVo;



public interface QaService extends BaseService {

	List<EvtQa> list(Long eventId, Long clientId);

	void save(Long eventId, String content);

	List<QaVo> genVos(List<EvtQa> pos);

	QaVo genVo(EvtQa po);


}
