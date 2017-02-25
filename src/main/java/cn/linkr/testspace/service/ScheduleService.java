package cn.linkr.testspace.service;


import java.util.List;

import com.alibaba.fastjson.JSONObject;

import cn.linkr.testspace.entity.EvtClient;
import cn.linkr.testspace.entity.EvtEvent;
import cn.linkr.testspace.entity.EvtScheduleItem;
import cn.linkr.testspace.vo.ScheduleItemVo;

public interface ScheduleService extends BaseService {

	List<EvtScheduleItem> listScheduleItemsByDate(Long eventId);

	List<EvtScheduleItem> listScheduleItemsBySession(Long eventId);

	List<ScheduleItemVo> genVosBySession(List<EvtScheduleItem> scheduleItemsBySession, boolean isNest);

	List<ScheduleItemVo> genVosByDate(List<EvtScheduleItem> scheduleItemsByDate);

	EvtScheduleItem save(ScheduleItemVo req);

	EvtScheduleItem genPo(JSONObject vo);
}
