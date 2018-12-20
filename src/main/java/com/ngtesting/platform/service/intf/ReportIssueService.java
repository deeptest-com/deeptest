package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.model.TstProject;

import java.util.List;
import java.util.Map;

public interface ReportIssueService extends ReportService {
    Map<String, List<Object>> chartIssueTrend(Integer orgId, TstProject.ProjectType type, Integer numb);
    Map<String, List<Object>> chartIssueAge(Integer projectId, TstProject.ProjectType type,
                                            Integer numb, Integer orgId, Integer prjId);
    List<Map<Object, Object>> chartIssueDistribByPriority(Integer orgId, TstProject.ProjectType type);
    List<Map<Object, Object>> chartIssueDistribByStatus(Integer orgId, TstProject.ProjectType type);

}
