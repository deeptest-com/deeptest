package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstCaseType;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface CaseTypeDao {
    List<TstCaseType> list(@Param("orgId") Integer orgId);

    TstCaseType get(@Param("id") Integer id,
                    @Param("orgId") Integer orgId);

    TstCaseType getDefault(Integer orgId);

    Integer save(TstCaseType vo);
    Integer update(TstCaseType vo);
    Integer delete(@Param("id") Integer id,
                @Param("orgId") Integer orgId);

    Integer removeDefault(@Param("orgId") Integer orgId);
    Integer setDefault(@Param("id") Integer id,
                    @Param("orgId") Integer orgId);
    Integer setOrder(@Param("id") Integer id,
                  @Param("ordr") Integer ordr,
                  @Param("orgId") Integer orgId);

    TstCaseType getPrev(@Param("ordr")Integer ordr, @Param("orgId") Integer orgId);

    TstCaseType getNext(@Param("ordr")Integer ordr, @Param("orgId") Integer orgId);

    Integer getMaxOrdrNumb(@Param("orgId") Integer orgId);
}
