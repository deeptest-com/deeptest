package com.ngtesting.platform.service;

import com.ngtesting.platform.model.IsuTypeSolution;

import java.util.List;

public interface IssueTypeSolutionService extends BaseService {

    List<IsuTypeSolution> list(Integer orgId);

    IsuTypeSolution get(Integer id, Integer orgId);
    IsuTypeSolution getDetail(Integer id, Integer orgId);

    IsuTypeSolution save(IsuTypeSolution vo, Integer orgId);

    Boolean delete(Integer id, Integer orgId);

    Boolean addType(Integer typeId, Integer solutionId, Integer orgId);
    Boolean removeType(Integer typeId, Integer solutionId, Integer orgId);

    Boolean addAll(Integer solutionId, Integer orgId);
    Boolean removeAll(Integer solutionId, Integer orgId);
}
