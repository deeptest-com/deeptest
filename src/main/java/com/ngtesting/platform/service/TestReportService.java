package com.ngtesting.platform.service;

import com.ngtesting.platform.model.TstProject;

import java.util.List;
import java.util.Map;

public interface TestReportService extends BaseService {

    Map<String, List<Object>> chartExcutionProcessByProject(Integer projectId, TstProject.ProjectType type, Integer numb);
    Map<String, List<Object>> chartDesignProgressByProject(Integer projectId, TstProject.ProjectType type, Integer numb);

    List<Map<Object, Object>> chartExecutionResultByPlan(Integer planId);
    Map<String, List<Object>> chartExecutionProcessByPlan(Integer planId, Integer numb);
    Map<String, Object> chartExecutionProcessByPlanUser(Integer planId, Integer numb);
    Map<String, Object> chartExecutionProgressByPlan(Integer planId, Integer numb);

    Map<String, List<Object>> countByStatus(List<Map> ls);

    Map<String, Object> countByUser(List<Map> ls);

    List<Map<Object, Object>> orderByStatus(Map map);
}
