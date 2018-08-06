package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstHistory;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface HistoryDao {
    List<TstHistory> listByProject(@Param("projectId") Integer projectId);

    List<TstHistory> listByProjectGroup(@Param("projectId") Integer projectId);

    List<TstHistory> listByOrg(@Param("orgId") Integer orgId);

    TstHistory get(@Param("id") Integer id);

    void create(TstHistory history);

}
