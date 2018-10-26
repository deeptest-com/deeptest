package com.ngtesting.platform.service;

import com.ngtesting.platform.model.IsuPageSolution;

import java.util.List;

public interface IssuePageSolutionService extends BaseService {

    List<IsuPageSolution> list(Integer orgId);
}
