package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.model.TstThread;

import java.util.List;


public interface ChatService extends BaseService {

	List<TstThread> listByEvent(Integer eventId);

	TstThread save(Integer eventId, Integer parentId, Integer clientId, String content);

	List<TstThread> enter(Integer eventId, Integer clientId);


}
