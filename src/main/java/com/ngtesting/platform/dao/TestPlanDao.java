package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstPlan;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface TestPlanDao {
    List<TstPlan> query(@Param("projectId") Integer projectId,
                             @Param("keywords") String keywords,
                             @Param("status") String status);

    List<TstPlan> listByProject(@Param("projectId") Integer projectId);

    List<TstPlan> listByProjectGroup(@Param("projectId") Integer projectId);

    List<TstPlan> listByOrg(@Param("orgId") Integer orgId);


    TstPlan get(@Param("id") Integer id,
                @Param("projectId") Integer projectId);

    Integer save(TstPlan vo);

    Integer update(TstPlan vo);

    Integer delete(@Param("id") Integer id,
                @Param("projectId") Integer projectId);

    void closePlanIfAllTaskClosed(@Param("planId") Integer planId);

    void start(@Param("id") Integer id);
}
