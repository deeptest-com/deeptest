package com.ngtesting.platform.service;

import com.ngtesting.platform.model.TstAlert;
import com.ngtesting.platform.model.TstTask;

import java.util.List;

public interface AlertService extends BaseService {

    List<TstAlert> list(Integer userId, Boolean isRead);

    List<TstAlert> scanAlerts(Integer userId);

    void create(TstTask task);

    void markAllRead(String ids, Integer userId);

    List<TstAlert> genVos(List<TstAlert> pos);

    TstAlert genVo(TstAlert po);
}
