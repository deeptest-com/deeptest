package com.ngtesting.platform.service.inf;

import com.ngtesting.platform.model.TstAlert;
import com.ngtesting.platform.model.TstRun;

import java.util.List;

public interface AlertService extends BaseService {

    List<TstAlert> list(Integer userId, Boolean isRead);

    List<TstAlert> scanTestAlert(Integer userId);

    void saveAlert(TstRun run);

    TstAlert getByRun(Integer id);

    void markAllReadPers(String ids);

    List<TstAlert> genVos(List<TstAlert> pos);

    TstAlert genVo(TstAlert po);
}
