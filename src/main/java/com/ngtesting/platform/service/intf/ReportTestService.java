package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.model.TstProject;

import java.util.List;
import java.util.Map;

public interface ReportTestService extends BaseService {

    Map<String, List<Object>> chartExcutionProcess(Integer projectId, TstProject.ProjectType type, Integer numb);
    Map<String, List<Object>> chartDesignProgress(Integer projectId, TstProject.ProjectType type, Integer numb);

    List<Map<Object, Object>> chartExecutionResultByPlan(Integer planId);
    Map<String, List<Object>> chartExecutionProcessByPlan(Integer planId, Integer numb);
    Map<String, Object> chartExecutionProcessByPlanUser(Integer planId, Integer numb);
    Map<String, Object> chartExecutionProgressByPlan(Integer planId, Integer numb);
}
