package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstCustomField;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface TstCustomFieldDao {
    List<TstCustomField> list(@Param("orgId") Integer orgId);
    List<TstCustomField> listForCaseByProject(@Param("orgId")Integer orgId,
                                              @Param("projectId")Integer projectId,
                                              @Param("applyTo")String applyTo);

    TstCustomField get(@Param("id") Integer id,
                       @Param("orgId") Integer orgId);
    TstCustomField getDetail(@Param("id") Integer id,
                             @Param("orgId") Integer orgId);

    Integer save(TstCustomField vo);
    Integer update(TstCustomField vo);
    Integer delete(@Param("id") Integer id,
                @Param("orgId") Integer orgId);

    Integer setOrder(@Param("id")Integer id,
                  @Param("ordr") Integer ordr,
                  @Param("orgId") Integer orgId);

    List<String> getLastUnusedColumn(@Param("orgId") Integer orgId);
    Integer getMaxOrdrNumb(@Param("orgId") Integer orgId);

    TstCustomField getPrev(@Param("ordr")Integer ordr, @Param("orgId") Integer orgId);
    TstCustomField getNext(@Param("ordr")Integer ordr, @Param("orgId") Integer orgId);

}
