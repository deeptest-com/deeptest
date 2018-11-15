package com.ngtesting.platform.service;

import com.ngtesting.platform.model.IsuWorkflowTransition;

import java.util.List;

public interface IssueWorkflowTransitionService extends BaseService {

    IsuWorkflowTransition save(IsuWorkflowTransition tran, List<Integer> projectRoleIds, Integer orgId);

    void delete(Integer id, Integer orgId);
}
