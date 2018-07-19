package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstMsg;
import org.apache.ibatis.annotations.Param;

import java.util.List;
import java.util.Map;

public interface TestReportDao {
    List<TstMsg> query(@Param("userId") Integer userId);

    List<Map> chart_design_progress_by_project(@Param("projectId") Integer projectId,
                                               @Param("projectType") String projectType,
                                               @Param("numb") Integer numb);

    List<Map> chart_execution_process_by_project(@Param("projectId") Integer projectId,
                                                      @Param("projectType") String projectType,
                                                      @Param("numb") Integer numb);
}
