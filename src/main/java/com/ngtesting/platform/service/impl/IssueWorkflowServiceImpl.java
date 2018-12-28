package com.ngtesting.platform.service.impl;


import com.ngtesting.platform.dao.IssueWorkflowDao;
import com.ngtesting.platform.dao.IssueWorkflowTransitionDao;
import com.ngtesting.platform.model.IsuStatus;
import com.ngtesting.platform.model.IsuWorkflow;
import com.ngtesting.platform.model.IsuWorkflowTransition;
import com.ngtesting.platform.service.intf.IssueStatusService;
import com.ngtesting.platform.service.intf.IssueWorkflowService;
import org.apache.commons.lang3.StringUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

@Service
public class IssueWorkflowServiceImpl extends BaseServiceImpl implements IssueWorkflowService {
    @Autowired
    IssueWorkflowDao workflowDao;
    @Autowired
    IssueWorkflowTransitionDao transitionDao;

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
    @Transactional
    public Boolean setDefault(Integer id, Integer orgId) {
        workflowDao.removeDefault(orgId);

        Integer count = workflowDao.setDefault(id, orgId);
        return count > 0;
    }

    @Override
    public Boolean delete(Integer id, Integer orgId) {
        Integer count = workflowDao.delete(id, orgId);

        return count > 0;
    }

    @Override
    public List<IsuStatus> listStatusForEdit(Integer id, Integer orgId) {
        List<IsuStatus> all = statusService.list(orgId);
        List<IsuStatus> statuses = workflowDao.listStatus(id);

        for (IsuStatus status : all) {
            if (statuses.contains(status)) {
                status.setSelected(true);
            } else {
                status.setSelected(false);
            }
        }

        return all;
    }

    @Override
    public List<IsuStatus> listStatusForDesign(Integer id) {
        List<IsuStatus> statuses = workflowDao.listStatus(id);

        return statuses;
    }

    @Override
    public Map<String, IsuWorkflowTransition> getTransitionMap(Integer id) {
        List<IsuWorkflowTransition> trans = transitionDao.listTransition(id, null);

        Map<String, IsuWorkflowTransition> map = new HashMap<>();
        for (IsuWorkflowTransition tran : trans) {
            map.put(tran.getSrcStatusId() + "-" + tran.getDictStatusId(), tran);
        }

        return map;
    }
}
