package cn.linkr.events.service;


import java.util.List;

import com.alibaba.fastjson.JSONObject;

import cn.linkr.events.entity.EvtClient;
import cn.linkr.events.entity.EvtEvent;
import cn.linkr.events.entity.EvtScheduleItem;
import cn.linkr.events.vo.ScheduleItemVo;

public interface ScheduleService extends BaseService {

	List<EvtScheduleItem> listScheduleItemsByDate(Long eventId);

	List<EvtScheduleItem> listScheduleItemsBySession(Long eventId);

	List<ScheduleItemVo> genVosBySession(List<EvtScheduleItem> scheduleItemsBySession, boolean isNest);

	List<ScheduleItemVo> genVosByDate(List<EvtScheduleItem> scheduleItemsByDate);

	EvtScheduleItem save(ScheduleItemVo req);

	EvtScheduleItem genPo(JSONObject vo);
}
