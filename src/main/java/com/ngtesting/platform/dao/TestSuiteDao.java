package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstCaseInSuite;

import java.util.List;

public interface TestSuiteDao {
    List<TstCaseInSuite> listCases(Integer suiteId);

    List<Integer> listCaseIds(Integer suiteId);
}
