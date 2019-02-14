package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.CustomField;
import org.apache.ibatis.annotations.Param;

import java.util.List;
import java.util.Map;

public interface CustomFieldDao {
    List<CustomField> list(@Param("orgId") Integer orgId,
                           @Param("applyTo") String applyTo,
                           @Param("keywords") String keywords);

    List<CustomField> listForCase(@Param("orgId") Integer orgId);

    CustomField get(@Param("id") Integer id,
                    @Param("orgId") Integer orgId);
    CustomField getDetail(@Param("id") Integer id,
                          @Param("orgId") Integer orgId);

    Integer save(CustomField vo);
    Integer update(CustomField vo);
    Integer delete(@Param("id") Integer id,
                   @Param("orgId") Integer orgId);

    Integer setOrder(@Param("id") Integer id,
                     @Param("ordr") Integer ordr,
                     @Param("orgId") Integer orgId);

    List<String> getLastUnusedColumn(@Param("orgId") Integer orgId);

    Integer getMaxOrdrNumb(@Param("orgId") Integer orgId,
                           @Param("applyTo") String applyTo);

    CustomField getPrev(@Param("ordr") Integer ordr,
                        @Param("orgId") Integer orgId,
                        @Param("applyTo") String applyTo);
    CustomField getNext(@Param("ordr") Integer ordr,
                        @Param("orgId") Integer orgId,
                        @Param("applyTo") String applyTo);

    List<Map> fetchInputMap();

    List<Map> listInput();
    List<Map> listType();
}
