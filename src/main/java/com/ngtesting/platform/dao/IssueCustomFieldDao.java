package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.IsuCustomField;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface IssueCustomFieldDao {
    List<IsuCustomField> list(@Param("orgId") Integer orgId);

    IsuCustomField get(@Param("id") Integer id,
                       @Param("orgId") Integer orgId);
    IsuCustomField getDetail(@Param("id") Integer id,
                             @Param("orgId") Integer orgId);

    Integer save(IsuCustomField vo);
    Integer update(IsuCustomField vo);
    Integer delete(@Param("id") Integer id,
                   @Param("orgId") Integer orgId);

    Integer setOrder(@Param("id") Integer id,
                     @Param("ordr") Integer ordr,
                     @Param("orgId") Integer orgId);

    List<String> getLastUnusedColumn(@Param("orgId") Integer orgId);
    Integer getMaxOrdrNumb(@Param("orgId") Integer orgId);

    IsuCustomField getPrev(@Param("ordr") Integer ordr, @Param("orgId") Integer orgId);
    IsuCustomField getNext(@Param("ordr") Integer ordr, @Param("orgId") Integer orgId);
}
