package com.ngtesting.platform.service;

import com.ngtesting.platform.entity.TestAlert;

import java.util.List;

public interface AlertService extends BaseService {

    List<TestAlert> list(Long userId);

    void scanTestPlan();

}
