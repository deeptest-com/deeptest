package com.ngtesting.platform.service;

import com.ngtesting.platform.model.IsuWorkflowSolution;

import java.util.List;
import java.util.Map;

public interface IssueWorkflowSolutionService extends BaseService {

    List<IsuWorkflowSolution> list(Integer orgId);

    IsuWorkflowSolution get(Integer solutionId, Integer orgId);
    Map<String,String> getItemsMap(Integer solutionId, Integer orgId);

    IsuWorkflowSolution save(IsuWorkflowSolution vo, Integer orgId);

    boolean delete(Integer id, Integer orgId);

    boolean changeItem(Integer typeId, Integer workflowId, Integer solutionId, Integer orgId);
}
