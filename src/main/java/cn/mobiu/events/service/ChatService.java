package cn.mobiu.events.service;

import java.util.List;

import cn.mobiu.events.entity.EvtThread;



public interface ChatService extends BaseService {

	List<EvtThread> listByEvent(Long eventId);

	EvtThread save(Long eventId, Long parentId, Long clientId, String content);

	List<EvtThread> enter(Long eventId, Long clientId);


}
