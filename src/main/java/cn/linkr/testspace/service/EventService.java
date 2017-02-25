package cn.linkr.testspace.service;


import java.util.List;

import com.alibaba.fastjson.JSONObject;

import cn.linkr.testspace.entity.EvtClient;
import cn.linkr.testspace.entity.EvtEvent;
import cn.linkr.testspace.entity.SysUser;
import cn.linkr.testspace.vo.EventVo;
import cn.linkr.testspace.vo.Page;

public interface EventService extends BaseService {

    public EvtEvent getDetail(Long eventId);

	EventVo genVo(EvtEvent po);

	List<EventVo> genVos(List<EvtEvent> pos);

	public Page list(Long companyId, String status, int startIndex, int itemsPerPage);

	public EvtEvent save(EventVo vo, Long userId, Long companyId);

	EvtEvent genPo(EventVo vo);

//	void updateStatus(EvtEvent event);

}
