package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstCaseInTask;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface CaseInTaskDao {
    List<TstCaseInTask> query(@Param("taskId") Integer taskId,
                              @Param("projectId") Integer projectId);

    TstCaseInTask get(@Param("id") Integer id);
    TstCaseInTask getDetail(@Param("id") Integer id,
                            @Param("projectId") Integer projectId);

    void setResult(TstCaseInTask po);

    void setResult(@Param("id") Integer id,
                   @Param("result") String result,
                   @Param("status") String status,
                   @Param("exeBy") Integer exeBy);
}
