package cn.mobiu.events.service;


import java.util.List;

import cn.mobiu.events.entity.EvtAround;
import cn.mobiu.events.entity.EvtAround.AroundType;
import cn.mobiu.events.vo.AroundVo;

public interface AroundService extends BaseService {

	List<EvtAround> list(Long eventId, AroundType type);

	List<AroundVo> genVos(List<EvtAround> pos);

	AroundVo genVo(EvtAround po);

}
