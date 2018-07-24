package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstCaseType;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface CaseTypeDao {
    List<TstCaseType> list(@Param("orgId") Integer orgId);

    TstCaseType get(@Param("id") Integer id);
}
