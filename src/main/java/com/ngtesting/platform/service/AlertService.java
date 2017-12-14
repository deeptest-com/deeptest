package com.ngtesting.platform.service;

import com.ngtesting.platform.entity.TestRun;
import com.ngtesting.platform.vo.TestAlertVo;

import java.util.List;

public interface AlertService extends BaseService {

    List<TestAlertVo> list(Long userId, Boolean isRead);

    List<TestRun> scanTestPlan(Long userId);

    List<TestAlertVo> genVos(List<TestRun> pos);

    List<TestAlertVo> genVosWithAction(List<TestRun> pos);

    TestAlertVo genVo(TestRun po);

    TestAlertVo genVoWithAction(TestRun po, Long startTimeOfToday, Long endTimeOfToday);
}
