package com.ngtesting.platform.service.impl;


import com.ngtesting.platform.dao.IssueWorkflowSolutionDao;
import com.ngtesting.platform.dao.IssueWorkflowTransitionDao;
import com.ngtesting.platform.dao.ProjectRoleDao;
import com.ngtesting.platform.dao.ProjectRoleEntityRelationDao;
import com.ngtesting.platform.model.IsuWorkflowSolutionItem;
import com.ngtesting.platform.model.IsuWorkflowTransition;
import com.ngtesting.platform.model.TstProjectRole;
import com.ngtesting.platform.service.intf.IssueStatusService;
import com.ngtesting.platform.service.intf.IssueWorkflowTransitionService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.LinkedHashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;

@Service
public class IssueWorkflowTransitionServiceImpl extends BaseServiceImpl implements IssueWorkflowTransitionService {
    @Autowired
    IssueWorkflowTransitionDao transitionDao;
    @Autowired
    IssueWorkflowSolutionDao issueWorkflowSolutionDao;
    @Autowired
    IssueStatusService statusService;
    @Autowired
    ProjectRoleDao projectRoleDao;
    @Autowired
    ProjectRoleEntityRelationDao projectRoleEntityRelationDao;

    @Override // TODO: cached，某个问题类型的状态对应的转换
    public Map<String, Map<String, List<IsuWorkflowTransition>>> getStatusTrainsMap(
            Integer projectId, Integer userId) {

        List<Integer> projectRoleIds = projectRoleEntityRelationDao
                .listIdsByUserAndProject(userId, projectId);

        List<IsuWorkflowSolutionItem> workflowItems =
                issueWorkflowSolutionDao.getIssueTypeWorkflow(projectId);

        Map<String, Map<String, List<IsuWorkflowTransition>>> typeMap = new LinkedHashMap();
        for (IsuWorkflowSolutionItem workflowItem : workflowItems) {
            Integer workflowId = workflowItem.getWorkflowId();

            List<IsuWorkflowTransition> trans = transitionDao.listTransition(workflowId, projectRoleIds);

            Map<String, List<IsuWorkflowTransition>> statusMap = new LinkedHashMap();
            for (IsuWorkflowTransition tran : trans) {
                String srcStatusId = tran.getSrcStatusId().toString();
                if (!statusMap.containsKey(srcStatusId)) {
                    statusMap.put(srcStatusId, new LinkedList<>());
                }

                statusMap.get(srcStatusId).add(tran);
            }

            typeMap.put(workflowItem.getTypeId().toString(), statusMap);
        }

        return typeMap;
    }

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
    public IsuWorkflowTransition emptyObject(Integer workflowId, Integer srcStatusId, Integer dictStatusId, Integer orgId) {
        IsuWorkflowTransition tran = transitionDao.emptyObject(srcStatusId, dictStatusId, orgId);
        tran.setWorkflowId(workflowId);
        return tran;
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

            transitionDao.removeAllRoles(tran.getId(), orgId);
        }

        transitionDao.addRoles(tran, projectRoleIds);

        return tran;
    }

    @Override
    public void delete(Integer id, Integer orgId) {
        transitionDao.removeAllRoles(id, orgId);
        transitionDao.delete(id, orgId);
    }

}
