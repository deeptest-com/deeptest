package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.IsuPage;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface IssuePageDao {

    List<IsuPage> list(Integer orgId);

    IsuPage get(@Param("id") Integer id,
                @Param("orgId") Integer orgId);

    void save(IsuPage vo);

    Integer update(IsuPage vo);

    Integer delete(@Param("id") Integer id,
                   @Param("orgId") Integer orgId);

//    void addField(IsuPageElement element);
//
//    Integer getMaxFieldOrdr(Integer tabId);

    Integer removeDefault(@Param("orgId") Integer orgId);
    Integer setDefault(@Param("id") Integer id,
                       @Param("orgId") Integer orgId);
}