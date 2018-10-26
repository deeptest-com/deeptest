package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.IsuPageSolution;

import java.util.List;

public interface IssuePageSolutionDao {

    List<IsuPageSolution> list(Integer orgId);
}
