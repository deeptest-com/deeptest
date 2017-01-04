package cn.mobiu.events.service.impl;

import java.util.List;

import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import cn.mobiu.events.entity.EvtClient;
import cn.mobiu.events.service.ClientService;

@Service
public class ClientServiceImpl extends BaseServiceImpl implements ClientService {

    @Override
    public EvtClient getByToken(String token) {
        DetachedCriteria dc = DetachedCriteria.forClass(EvtClient.class);
        dc.add(Restrictions.eq("token", token));

        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.addOrder(Order.asc("id"));
        List<EvtClient> ls = findAllByCriteria(dc);

        EvtClient client = null;
        if (ls.size() > 0) {
        	client = ls.get(0);
        }

        return client;
    }
}
