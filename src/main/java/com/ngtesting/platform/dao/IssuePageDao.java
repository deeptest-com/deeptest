package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.IsuPage;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface IssuePageDao {

    List<IsuPage> list(Integer orgId);

    IsuPage get(@Param("id") Integer id, @Param("orgId") Integer orgId);
}
