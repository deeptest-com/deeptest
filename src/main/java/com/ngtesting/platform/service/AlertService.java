package com.ngtesting.platform.service;

import com.ngtesting.platform.entity.TestAlert;
import com.ngtesting.platform.entity.TestRun;
import com.ngtesting.platform.vo.TestAlertVo;

import java.util.List;

public interface AlertService extends BaseService {

    List<TestAlertVo> list(Long userId, Boolean isRead);

    List<TestAlert> scanTestAlert(Long userId);

    TestAlert saveAlert(TestRun run);

    TestAlert getByRun(Long id);

    void markAllReadPers(String ids);

    List<TestAlertVo> genVos(List<TestAlert> pos);

    TestAlertVo genVo(TestAlert po);
}
