package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstCaseType;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface CaseTypeDao {
    List<TstCaseType> list(@Param("orgId") Integer orgId);

    TstCaseType get(@Param("id") Integer id,
                    @Param("orgId") Integer orgId);


    void save(TstCaseType vo);
    void update(TstCaseType vo);
    void delete(@Param("id") Integer id,
                @Param("orgId") Integer orgId);

    void removeDefault(@Param("orgId") Integer orgId);
    void setDefault(@Param("id") Integer id, @Param("orgId") Integer orgId);
    void setOrder(@Param("id") Integer id, @Param("ordr") Integer ordr,
                  @Param("orgId") Integer orgId);

    TstCaseType getPrev(@Param("ordr")Integer ordr, @Param("orgId") Integer orgId);

    TstCaseType getNext(@Param("ordr")Integer ordr, @Param("orgId") Integer orgId);

    Integer getMaxOrdrNumb(@Param("orgId") Integer orgId);
}
