package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.IsuQuery;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface IssueQueryDao {
    List<IsuQuery> list(@Param("prjId") Integer prjId,
                        @Param("userId") Integer userId,
                        @Param("keywords") String keywords);

    List<IsuQuery> listRecentQuery(@Param("prjId") Integer prjId,
                                   @Param("userId") Integer userId);

    IsuQuery get(@Param("id") Integer id,
                 @Param("userId") Integer userId);

    Integer save(IsuQuery query);

    Integer update(@Param("model") IsuQuery model,
                    @Param("userId") Integer userId);

    Integer delete(@Param("id") Integer id,
                   @Param("userId") Integer userId);

    Integer updateUseTime(IsuQuery query);
}
