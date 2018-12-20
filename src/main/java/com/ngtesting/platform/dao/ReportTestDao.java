package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstMsg;
import org.apache.ibatis.annotations.Param;

import java.util.List;
import java.util.Map;

public interface ReportTestDao {
    List<TstMsg> query(@Param("userId") Integer userId);

    List<Map> chartDesignProgressByProject(@Param("projectId") Integer projectId,
                                           @Param("projectType") String projectType,
                                           @Param("numb") Integer numb);

    List<Map> chartExecutionProcessByProject(@Param("projectId") Integer projectId,
                                             @Param("projectType") String projectType,
                                             @Param("numb") Integer numb);

    List<Map> chartExecutionResultByPlan(@Param("planId") Integer planId);
    List<Map> chartExecutionProcessByPlan(@Param("planId") Integer planId,
                                          @Param("numb") Integer numb);

    List<Map> chartExecutionProcessByPlanUser(@Param("planId") Integer planId,
                                                   @Param("numb") Integer numb);

    List<Map> chartExecutionProgressByPlan(@Param("planId") Integer planId,
                                                @Param("numb") Integer numb);
}
