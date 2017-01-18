package cn.linkr.events.service;

import cn.linkr.events.entity.EvtClient;

public interface ClientService extends BaseService {
    EvtClient getByToken(String token);
}
