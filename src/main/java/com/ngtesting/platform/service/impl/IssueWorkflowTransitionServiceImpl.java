package com.ngtesting.platform.service.impl;


import com.ngtesting.platform.dao.IssueWorkflowTransitionDao;
import com.ngtesting.platform.dao.ProjectRoleDao;
import com.ngtesting.platform.model.IsuWorkflowTransition;
import com.ngtesting.platform.model.TstProjectRole;
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
    @Autowired
    ProjectRoleDao projectRoleDao;

    @Override
    public List<TstProjectRole> listProjectRoles(Integer id, Integer orgId) {
        List<TstProjectRole> allRoles = projectRoleDao.query(orgId, null, null);
        List<Integer> roleIds = transitionDao.listProjectRoleId(id, orgId);

        for (TstProjectRole role : allRoles) {
            if (roleIds.contains(role.getId())) {
                role.setSelected(true);
            }
        }

        return allRoles;
    }

    @Override
    public IsuWorkflowTransition get(Integer id, Integer orgId) {
        return transitionDao.get(id, orgId);
    }

    @Override
    public IsuWorkflowTransition emptyObject(Integer srcStatusId, Integer dictStatusId, Integer orgId) {
        return transitionDao.emptyObject(srcStatusId, dictStatusId, orgId);
    }

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
