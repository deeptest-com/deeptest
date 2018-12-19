package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.model.IsuWorkflowTransition;
import com.ngtesting.platform.model.TstProjectRole;

import java.util.List;
import java.util.Map;

public interface IssueWorkflowTransitionService extends BaseService {
    Map<Integer, Map<Integer, List<IsuWorkflowTransition>>> getStatusTrainsMap(Integer projectId);

    List<TstProjectRole> listProjectRoles(Integer id, Integer orgId);
    IsuWorkflowTransition get(Integer id, Integer orgId);

    IsuWorkflowTransition save(IsuWorkflowTransition tran, List<Integer> projectRoleIds, Integer orgId);

    void delete(Integer id, Integer orgId);

    IsuWorkflowTransition emptyObject(Integer srcStatusId, Integer dictStatusId, Integer orgId);
}
