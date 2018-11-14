package com.ngtesting.platform.service;

import com.ngtesting.platform.model.IsuStatus;
import com.ngtesting.platform.model.IsuWorkflow;

import java.util.List;

public interface IssueWorkflowService extends BaseService {
    List<IsuWorkflow> list(Integer orgId);
    IsuWorkflow get(Integer id, Integer orgId);

    IsuWorkflow save(IsuWorkflow vo, List<Integer> statusIds, Integer orgId);

    List<IsuStatus> listStatusForEdit(Integer id, Integer orgId);

    List<IsuStatus> listStatusForDesign(Integer id);
}
