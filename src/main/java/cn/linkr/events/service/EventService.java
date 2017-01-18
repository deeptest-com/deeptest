package cn.linkr.events.service;


import java.util.List;

import com.alibaba.fastjson.JSONObject;

import cn.linkr.events.entity.EvtClient;
import cn.linkr.events.entity.EvtEvent;
import cn.linkr.events.vo.EventVo;
import cn.linkr.events.vo.Page;

public interface EventService extends BaseService {

    public EvtEvent getDetail(Long eventId);

	EventVo genVo(EvtEvent po);

	List<EventVo> genVos(List<EvtEvent> pos);

	public Page list(Long companyId, String status, int startIndex, int itemsPerPage);

	public EvtEvent save(EventVo vo, EvtClient client);

	EvtEvent genPo(EventVo vo);

//	void updateStatus(EvtEvent event);

}
