package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.model.IsuPageSolution;

import java.util.List;
import java.util.Map;

public interface IssuePageSolutionService extends BaseService {

    List<IsuPageSolution> list(Integer orgId);

    IsuPageSolution get(Integer solutionId, Integer orgId);
    Map<String, Map<String, String>> getItemsMap(Integer solutionId, Integer orgId);

    IsuPageSolution save(IsuPageSolution vo, Integer orgId);

    boolean delete(Integer id, Integer orgId);

    boolean changeItem(Integer typeId, String opt, Integer pageId, Integer solutionId, Integer orgId);

    Boolean setDefault(Integer id, Integer orgId);

    // For Project
    IsuPageSolution getByProject(Integer projectId, Integer orgId);

    void setByProject(Integer solutionId, Integer projectId, Integer orgId);
}
