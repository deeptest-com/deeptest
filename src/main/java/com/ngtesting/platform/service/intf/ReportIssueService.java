package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.model.TstProject;

import java.util.List;
import java.util.Map;

public interface ReportIssueService extends ReportService {
    Map<String, List<Object>> chartIssueTrend(Integer id, TstProject.ProjectType type, Integer numb);

    Map<String, List<Object>> chartIssueAgeByProject(Integer projectId, Integer numb, Integer orgId);
    Map<String, List<Object>> chartIssueAgeByOrgOrGroup(Integer id, TstProject.ProjectType type, Integer numb);

    List<Map<Object, Object>> chartIssueDistribByPriority(Integer id, TstProject.ProjectType type);
    List<Map<Object, Object>> chartIssueDistribByStatus(Integer id, TstProject.ProjectType type);

}
