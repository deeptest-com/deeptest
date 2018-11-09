package com.ngtesting.platform.service;

import com.ngtesting.platform.model.IsuTypeSolution;

import java.util.List;

public interface IssueTypeSolutionService extends BaseService {

    List<IsuTypeSolution> list(Integer orgId);

    List<IsuTypeSolution> list(Integer orgId, Integer prjId);

    IsuTypeSolution get(Integer id, Integer orgId);

    IsuTypeSolution save(IsuTypeSolution vo, Integer orgId);

    Boolean delete(Integer id, Integer orgId);

}
