package com.ngtesting.platform.service;


import java.util.List;

import com.alibaba.fastjson.JSONObject;

import com.ngtesting.platform.entity.EvtClient;
import com.ngtesting.platform.entity.EvtEvent;
import com.ngtesting.platform.entity.SysUser;
import com.ngtesting.platform.vo.EventVo;
import com.ngtesting.platform.vo.Page;

public interface EventService extends BaseService {

    public EvtEvent getDetail(Long eventId);

	EventVo genVo(EvtEvent po);

	List<EventVo> genVos(List<EvtEvent> pos);

	public Page list(Long companyId, String status, int startIndex, int itemsPerPage);

	public EvtEvent save(EventVo vo, Long userId, Long companyId);

	EvtEvent genPo(EventVo vo);

//	void updateStatus(EvtEvent event);

}
