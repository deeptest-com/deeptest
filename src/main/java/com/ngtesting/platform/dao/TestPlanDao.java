package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstMsg;
import com.ngtesting.platform.model.TstPlan;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface TestPlanDao {
    List<TstMsg> query(@Param("userId") Integer userId);

    List<TstPlan> listByProject(@Param("projectId") Integer projectId);

    List<TstPlan> listByProjectGroup(@Param("projectId") Integer projectId);
}
