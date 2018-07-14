package com.ngtesting.platform.service;

import com.ngtesting.platform.model.TstProject;

import java.util.List;
import java.util.Map;

public interface TestReportService extends BaseService {

    Map<String, List<Object>> chart_excution_process_by_project(Integer projectId, TstProject.ProjectType type, Integer numb);
    Map<String, List<Object>> chart_design_progress_by_project(Integer projectId, TstProject.ProjectType type, Integer numb);

    List<Map<Object, Object>> chart_execution_result_by_plan(Integer planId);
    Map<String, List<Object>> chart_execution_process_by_plan(Integer planId, Integer numb);
    Map<String, Object> chart_execution_process_by_plan_user(Integer planId, Integer numb);
    Map<String, Object> chart_execution_progress_by_plan(Integer planId, Integer numb);

    Map<String, List<Object>> countByStatus(List<Object[]> ls);

    Map<String, Object> countByUser(List<Object[]> ls);

    List<Map<Object, Object>> orderByStatus(Map map);

    String getUserName(String id);
}
