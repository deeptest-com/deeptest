package com.ngtesting.platform.dao;

import org.apache.ibatis.annotations.Param;

public interface IssueOptDao {

    void statusTran(@Param("id") Integer id,
                    @Param("dictStatusId") Integer dictStatusId,
                    @Param("finalVal") Boolean finalVal,
                    @Param("projectId") Integer projectId);

    void assign(@Param("id") Integer id,
                @Param("userId") Integer userId,
                @Param("projectId") Integer projectId);
}
