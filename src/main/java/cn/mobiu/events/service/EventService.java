package cn.mobiu.events.service;


import java.util.List;

import com.alibaba.fastjson.JSONObject;

import cn.mobiu.events.entity.EvtClient;
import cn.mobiu.events.entity.EvtEvent;
import cn.mobiu.events.vo.EventVo;
import cn.mobiu.events.vo.Page;

public interface EventService extends BaseService {

    public EvtEvent getDetail(Long eventId);

	EventVo genVo(EvtEvent po);

	List<EventVo> genVos(List<EvtEvent> pos);

	public Page list(Long companyId, String status, int startIndex, int itemsPerPage);

	public EvtEvent save(EventVo vo, EvtClient client);

	EvtEvent genPo(EventVo vo);

//	void updateStatus(EvtEvent event);

}
