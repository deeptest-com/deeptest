package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.model.TstThread;
import com.ngtesting.platform.service.intf.ChatService;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class ChatServiceImpl extends BaseServiceImpl implements ChatService {

	@Override
	public List<TstThread> listByEvent(Integer eventId) {
//	    DetachedCriteria dc = DetachedCriteria.forClass(TstThread.class);
//        dc.save(Restrictions.eq("eventId", eventId));
//
//        dc.save(Restrictions.eq("deleted", Boolean.FALSE));
//        dc.save(Restrictions.eq("disabled", Boolean.FALSE));
//        dc.addOrder(Order.asc("id"));
//        Page listByPage = findPage(dc, 0, 10);
//
//        return listByPage.getItemsMap();

		return null;
	}

	@Override
	public TstThread save(Integer eventId, Integer parentId, Integer clientId, String content) {
//		TstThread thread = new TstThread(eventId, clientId, parentId, content);
//
//		saveOrUpdate(thread);
//		return thread;

		return null;
	}

	@Override
	public List<TstThread> enter(Integer eventId, Integer clientId) {
		List<TstThread> list = listByEvent(eventId);

		return list;
	}

}
