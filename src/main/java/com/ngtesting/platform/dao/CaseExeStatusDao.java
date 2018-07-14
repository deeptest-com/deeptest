package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstCaseExeStatus;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface CaseExeStatusDao {
    List<TstCaseExeStatus> listExeStatus(@Param("orgId") Integer orgId);
}
