package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.IssuePageSolutionDao;
import com.ngtesting.platform.model.IsuPageSolution;
import com.ngtesting.platform.model.IsuPageSolutionItem;
import com.ngtesting.platform.service.intf.IssuePageSolutionService;
import com.ngtesting.platform.service.intf.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;

@Service
public class IssuePageSolutionServiceImpl extends BaseServiceImpl implements IssuePageSolutionService {
    @Autowired
    UserService userService;

    @Autowired
    IssuePageSolutionDao pageSolutionDao;

    @Override
    public List<IsuPageSolution> list(Integer orgId) {
        return pageSolutionDao.list(orgId);
    }

    @Override
    public IsuPageSolution get(Integer solutionId, Integer orgId) {
        return pageSolutionDao.get(solutionId, orgId);
    }

    @Override
    public Map<String, Map<String, String>> getItemsMap(Integer solutionId, Integer orgId) {
        List<IsuPageSolutionItem> items = pageSolutionDao.getItems(solutionId, orgId);

        Map<String, Map<String, String>> map = new LinkedHashMap<>();
        for (IsuPageSolutionItem item : items) {
            String typeKey = item.getTypeId() + "-" + item.getTypeName();
            if (!map.containsKey(typeKey)) {
                map.put(typeKey, new LinkedHashMap<>());
            }

            String pageKey = item.getPageId() + "-" + item.getPageName();
            map.get(typeKey).put(item.getOpt().toString(), pageKey);
        }

        return map;
    }

    @Override
    public IsuPageSolution save(IsuPageSolution vo, Integer orgId) {
        vo.setOrgId(orgId);

        if (vo.getId() == null) {
            pageSolutionDao.save(vo);
        } else {
            Integer count = pageSolutionDao.update(vo);
            if (count == 0) {
                return null;
            }
        }

        return vo;
    }

    @Override
    public boolean delete(Integer id, Integer orgId) {
        Integer count = pageSolutionDao.delete(id, orgId);
        return count > 0;
    }

    @Override
    public boolean changeItem(Integer typeId, String opt, Integer pageId, Integer solutionId, Integer orgId) {
        Integer count = pageSolutionDao.changeItem(typeId, opt, pageId, solutionId, orgId);
        return count > 0;
    }

    @Override
    @Transactional
    public Boolean setDefault(Integer id, Integer orgId) {
        pageSolutionDao.removeDefault(orgId);

        Integer count = pageSolutionDao.setDefault(id, orgId);
        return count > 0;
    }

    // For Project
    @Override
    public IsuPageSolution getByProject(Integer projectId, Integer orgId) {
        return pageSolutionDao.getByProject(projectId, orgId);
    }

    @Override
    public void setByProject(Integer solutionId, Integer projectId, Integer orgId) {
        pageSolutionDao.setByProject(solutionId, projectId, orgId);
    }
}
