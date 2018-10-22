package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.IsuStatus;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface IssueStatusDao {
    List<IsuStatus> listExeStatus(@Param("orgId") Integer orgId);
}
