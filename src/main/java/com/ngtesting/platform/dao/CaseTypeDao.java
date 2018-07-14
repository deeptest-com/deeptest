package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstCaseType;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface CaseTypeDao {
    List<TstCaseType> listType(@Param("orgId") Integer orgId);
}
