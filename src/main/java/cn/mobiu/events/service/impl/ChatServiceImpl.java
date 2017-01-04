package cn.mobiu.events.service.impl;

import java.util.List;

import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;
import cn.mobiu.events.entity.EvtThread;
import cn.mobiu.events.service.ChatService;
import cn.mobiu.events.vo.Page;

@Service
public class ChatServiceImpl extends BaseServiceImpl implements ChatService {

	@Override
	public List<EvtThread> listByEvent(Long eventId) {
	    DetachedCriteria dc = DetachedCriteria.forClass(EvtThread.class);
        dc.add(Restrictions.eq("eventId", eventId));
        
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.addOrder(Order.asc("id"));
        Page page = findPage(dc, 0, 10);
        
        return page.getItems();
	}
	
	@Override
	public EvtThread save(Long eventId, Long parentId, Long clientId, String content) {
		EvtThread thread = new EvtThread(eventId, clientId, parentId, content);
		
		saveOrUpdate(thread);
		return thread;
	}

	@Override
	public List<EvtThread> enter(Long eventId, Long clientId) {
		List<EvtThread> list = listByEvent(eventId);
		
		
		
		return list;
	}
    
}
