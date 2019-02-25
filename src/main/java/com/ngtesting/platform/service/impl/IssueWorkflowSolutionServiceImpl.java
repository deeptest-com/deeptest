package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.IssueTypeDao;
import com.ngtesting.platform.dao.IssueWorkflowDao;
import com.ngtesting.platform.dao.IssueWorkflowSolutionDao;
import com.ngtesting.platform.model.IsuType;
import com.ngtesting.platform.model.IsuWorkflow;
import com.ngtesting.platform.model.IsuWorkflowSolution;
import com.ngtesting.platform.model.IsuWorkflowSolutionItem;
import com.ngtesting.platform.service.intf.IssueWorkflowSolutionService;
import com.ngtesting.platform.service.intf.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;

@Service
public class IssueWorkflowSolutionServiceImpl extends BaseServiceImpl implements IssueWorkflowSolutionService {
    @Autowired
    UserService userService;

    @Autowired
    IssueWorkflowSolutionDao workflowSolutionDao;
    @Autowired
    IssueWorkflowDao workflowDao;

    @Autowired
    IssueTypeDao issueTypeDao;

    @Override
    public List<IsuWorkflowSolution> list(Integer orgId) {
        return workflowSolutionDao.list(orgId);
    }

    @Override
    public IsuWorkflowSolution get(Integer solutionId, Integer orgId) {
        return workflowSolutionDao.get(solutionId, orgId);
    }

    @Override
    public Map<String,String> getItemsMap(Integer solutionId, Integer orgId) {
        List<IsuWorkflowSolutionItem> items = workflowSolutionDao.getItems(solutionId, orgId);

        Map<String, String> map = new LinkedHashMap<>();
        for (IsuWorkflowSolutionItem item : items) {
            String typeKey = item.getTypeId() + "-" + item.getTypeName();

            String workflowKey = item.getWorkflowId() + "-" + item.getWorkflowName();
            map.put(typeKey, workflowKey);
        }

        List<IsuType> types = issueTypeDao.list(orgId);
        for (IsuType type : types) {
            String typeKey = type.getId() + "-" + type.getLabel();
            if (!map.containsKey(typeKey)) {
                map.put(typeKey, "");
            }
        }

        return map;
    }

    @Override
    public IsuWorkflowSolution save(IsuWorkflowSolution vo, Integer orgId) {
        vo.setOrgId(orgId);

        if (vo.getId() == null) {
            workflowSolutionDao.save(vo);

            IsuWorkflow workflow = workflowDao.getDefault(orgId);
            workflowSolutionDao.setDefaultPage(workflow.getId(), vo.getId(), orgId);
        } else {
            Integer count = workflowSolutionDao.update(vo);
            if (count == 0) {
                return null;
            }
        }

        return vo;
    }

    @Override
    public boolean delete(Integer id, Integer orgId) {
        Integer count = workflowSolutionDao.delete(id, orgId);
        return count > 0;
    }

    @Override
    @Transactional
    public Boolean setDefault(Integer id, Integer orgId) {
        workflowSolutionDao.removeDefault(orgId);

        Integer count = workflowSolutionDao.setDefault(id, orgId);
        return count > 0;
    }

    @Override
    public boolean changeItem(Integer typeId, Integer workflowId, Integer solutionId, Integer orgId) {
        Integer count = workflowSolutionDao.changeItem(typeId, workflowId, solutionId, orgId);
        return count > 0;
    }

    // For Project
    @Override
    public IsuWorkflowSolution getByProject(Integer projectId, Integer orgId) {
        return workflowSolutionDao.getByProject(projectId, orgId);
    }

    @Override
    public void setByProject(Integer solutionId, Integer projectId, Integer orgId) {
        workflowSolutionDao.setByProject(solutionId, projectId, orgId);
    }
}
