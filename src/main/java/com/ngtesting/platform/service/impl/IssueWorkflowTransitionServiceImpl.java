package com.ngtesting.platform.service.impl;


import com.ngtesting.platform.dao.IssueWorkflowTransitionDao;
import com.ngtesting.platform.model.IsuWorkflowTransition;
import com.ngtesting.platform.service.IssueStatusService;
import com.ngtesting.platform.service.IssueWorkflowTransitionService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class IssueWorkflowTransitionServiceImpl extends BaseServiceImpl implements IssueWorkflowTransitionService {
    @Autowired
    IssueWorkflowTransitionDao transitionDao;

    @Autowired
    IssueStatusService statusService;

    @Override
    public IsuWorkflowTransition save(IsuWorkflowTransition tran, List<Integer> projectRoleIds, Integer orgId) {
        if (tran.getId() == null) {
            tran.setOrgId(orgId);
            transitionDao.save(tran);
        } else {
            Integer count = transitionDao.update(tran);
            if (count == 0) {
                return null;
            }
        }

        return tran;
    }

    @Override
    public void delete(Integer id, Integer orgId) {
        transitionDao.delete(id, orgId);
    }
}
