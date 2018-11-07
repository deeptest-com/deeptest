package com.ngtesting.platform.service;

import com.ngtesting.platform.model.IsuPageSolution;

import java.util.List;

public interface IssuePageSolutionService extends BaseService {

    List<IsuPageSolution> list(Integer orgId);

    IsuPageSolution get(Integer solutionId, Integer orgId);

    IsuPageSolution save(IsuPageSolution vo, Integer orgId);

    boolean delete(Integer id, Integer orgId);
}
