package cn.mobiu.events.service;

import cn.mobiu.events.entity.EvtClient;

public interface ClientService extends BaseService {
    EvtClient getByToken(String token);
}
