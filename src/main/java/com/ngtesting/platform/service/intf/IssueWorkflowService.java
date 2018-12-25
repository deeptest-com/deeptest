package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.model.IsuStatus;
import com.ngtesting.platform.model.IsuWorkflow;
import com.ngtesting.platform.model.IsuWorkflowTransition;

import java.util.List;
import java.util.Map;

public interface IssueWorkflowService extends BaseService {
    List<IsuWorkflow> list(Integer orgId);
    IsuWorkflow get(Integer id, Integer orgId);

    IsuWorkflow save(IsuWorkflow vo, List<Integer> statusIds, Integer orgId);

    List<IsuStatus> listStatusForEdit(Integer id, Integer orgId);

    List<IsuStatus> listStatusForDesign(Integer id);

    Map<String, IsuWorkflowTransition> getTransitionMap(Integer id);

    Boolean setDefault(Integer id, Integer orgId);

    Boolean delete(Integer id, Integer orgId);
}
