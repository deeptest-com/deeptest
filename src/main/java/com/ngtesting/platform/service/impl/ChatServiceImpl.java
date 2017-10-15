package com.ngtesting.platform.service.impl;

import java.util.List;

import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import com.ngtesting.platform.entity.TestThread;
import com.ngtesting.platform.service.ChatService;
import com.ngtesting.platform.vo.Page;

@Service
public class ChatServiceImpl extends BaseServiceImpl implements ChatService {

	@Override
	public List<TestThread> listByEvent(Long eventId) {
	    DetachedCriteria dc = DetachedCriteria.forClass(TestThread.class);
        dc.add(Restrictions.eq("eventId", eventId));

        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.addOrder(Order.asc("id"));
        Page page = findPage(dc, 0, 10);

        return page.getItems();
	}

	@Override
	public TestThread save(Long eventId, Long parentId, Long clientId, String content) {
		TestThread thread = new TestThread(eventId, clientId, parentId, content);

		saveOrUpdate(thread);
		return thread;
	}

	@Override
	public List<TestThread> enter(Long eventId, Long clientId) {
		List<TestThread> list = listByEvent(eventId);



		return list;
	}

}
