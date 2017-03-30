package com.ngtesting.platform.service;


import java.util.List;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.EvtScheduleItem;
import com.ngtesting.platform.vo.ScheduleItemVo;

public interface ScheduleService extends BaseService {

	List<EvtScheduleItem> listScheduleItemsByDate(Long eventId);

	List<EvtScheduleItem> listScheduleItemsBySession(Long eventId);

	List<ScheduleItemVo> genVosBySession(List<EvtScheduleItem> scheduleItemsBySession, boolean isNest);

	List<ScheduleItemVo> genVosByDate(List<EvtScheduleItem> scheduleItemsByDate);

	EvtScheduleItem save(ScheduleItemVo req);

	EvtScheduleItem genPo(JSONObject vo);
}
