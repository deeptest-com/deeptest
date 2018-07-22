package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstHistory;

import java.util.List;

public interface HistoryDao {
    List<TstHistory> listByProject(Integer projectId);

    List<TstHistory> listByProjectGroup(Integer projectId);

    List<TstHistory> listByOrg(Integer orgId);
}
