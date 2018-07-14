package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstCasePriority;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface CasePriorityDao {
    List<TstCasePriority> listPriority(@Param("orgId") Integer orgId);
}
