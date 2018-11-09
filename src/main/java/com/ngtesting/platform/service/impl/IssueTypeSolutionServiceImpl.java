package com.ngtesting.platform.service.impl;


import com.ngtesting.platform.dao.IssueTypeSolutionDao;
import com.ngtesting.platform.model.IsuTypeSolution;
import com.ngtesting.platform.service.IssueTypeSolutionService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class IssueTypeSolutionServiceImpl extends BaseServiceImpl implements IssueTypeSolutionService {
    @Autowired
    IssueTypeSolutionDao solutionDao;

    @Override
    public List<IsuTypeSolution> list(Integer orgId) {
        List<IsuTypeSolution> ls = solutionDao.list(orgId);

        return ls;
    }

    @Override
    public List<IsuTypeSolution> list(Integer orgId, Integer prjId) {
        List<IsuTypeSolution> ls = solutionDao.list(orgId);

        return ls;
    }

    @Override
    public IsuTypeSolution get(Integer id, Integer orgId) {
        return solutionDao.get(id, orgId);
    }

    @Override
    public IsuTypeSolution save(IsuTypeSolution vo, Integer orgId) {

        if (vo.getId() == null) {
            vo.setOrgId(orgId);
            solutionDao.save(vo);
        } else {
            Integer count = solutionDao.update(vo);
            if (count == 0) {
                return null;
            }
        }

        return vo;
    }

    @Override
    public Boolean delete(Integer id, Integer orgId) {
        Integer count = solutionDao.delete(id, orgId);
        if (count == 0) {
            return false;
        }

        return true;
    }

}
