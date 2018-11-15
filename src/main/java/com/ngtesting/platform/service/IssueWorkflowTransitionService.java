package com.ngtesting.platform.service;

import com.ngtesting.platform.model.IsuWorkflowTransition;
import com.ngtesting.platform.model.TstProjectRole;

import java.util.List;

public interface IssueWorkflowTransitionService extends BaseService {
    List<TstProjectRole> listProjectRoles(Integer id, Integer orgId);

    IsuWorkflowTransition save(IsuWorkflowTransition tran, List<Integer> projectRoleIds, Integer orgId);

    void delete(Integer id, Integer orgId);
}
