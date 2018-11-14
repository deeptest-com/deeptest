package com.ngtesting.platform.service.impl;


import com.ngtesting.platform.dao.IssueWorkflowDao;
import com.ngtesting.platform.model.IsuStatus;
import com.ngtesting.platform.model.IsuWorkflow;
import com.ngtesting.platform.service.IssueStatusService;
import com.ngtesting.platform.service.IssueWorkflowService;
import org.apache.commons.lang3.StringUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

@Service
public class IssueWorkflowServiceImpl extends BaseServiceImpl implements IssueWorkflowService {
    @Autowired
    IssueWorkflowDao workflowDao;

    @Autowired
    IssueStatusService statusService;

    @Override
    public List<IsuWorkflow> list(Integer orgId) {
        return workflowDao.list(orgId);
    }

    @Override
    public IsuWorkflow get(Integer id, Integer orgId) {
        return workflowDao.get(id, orgId);
    }

    @Override
    @Transactional
    public IsuWorkflow save(IsuWorkflow vo, List<Integer> statusIds, Integer orgId) {
        String str = statusIds.size() > 0?StringUtils.join(statusIds,","):"-1";

        if (vo.getId() == null) {
            vo.setOrgId(orgId);
            workflowDao.save(vo);
        } else {
            Integer count = workflowDao.update(vo);
            if (count == 0) {
                return null;
            }

            workflowDao.removeTransitions(vo.getId(), str, orgId);
        }

        workflowDao.removeStatuses(vo.getId(), str, orgId);

        if (statusIds.size() > 0) {
            workflowDao.saveStatuses(vo.getId(), statusIds, orgId);
        }

        return vo;
    }

    @Override
    public List<IsuStatus> listStatus(IsuWorkflow po, Integer orgId) {
        List<IsuStatus> all = statusService.list(orgId);
        List<IsuStatus> statuses = workflowDao.listStatus(po.getId());

        for (IsuStatus status : all) {
            if (statuses.contains(status)) {
                status.setSelected(true);
            } else {
                status.setSelected(false);
            }
        }

        return all;
    }
}
