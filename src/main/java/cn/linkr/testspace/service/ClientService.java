package cn.linkr.testspace.service;

import cn.linkr.testspace.entity.EvtClient;

public interface ClientService extends BaseService {
    EvtClient getByToken(String token);
}
