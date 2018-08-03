package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstCustomField;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface CustomFieldDao {
    List<TstCustomField> list(@Param("orgId") Integer orgId);
    List<TstCustomField> listForCaseByProject(@Param("orgId")Integer orgId,
                                              @Param("projectId")Integer projectId,
                                              @Param("applyTo")String applyTo);


    TstCustomField get(@Param("id") Integer id);
    TstCustomField getDetail(@Param("id") Integer id);

    List<String> getLastUnusedColumn(@Param("orgId") Integer orgId);

    Integer getMaxOrdrNumb(@Param("orgId") Integer orgId);

    void save(TstCustomField vo);
    void update(TstCustomField vo);
    void delete(@Param("id") Integer id);

    TstCustomField getPrev(@Param("ordr")Integer ordr, @Param("orgId") Integer orgId);

    TstCustomField getNext(@Param("ordr")Integer ordr, @Param("orgId") Integer orgId);

    void setOrder(@Param("id")Integer id, @Param("ordr") Integer ordr);

}
