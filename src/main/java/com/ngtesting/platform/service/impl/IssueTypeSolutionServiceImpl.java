package com.ngtesting.platform.service.impl;


import com.ngtesting.platform.dao.IssueTypeDao;
import com.ngtesting.platform.dao.IssueTypeSolutionDao;
import com.ngtesting.platform.model.IsuType;
import com.ngtesting.platform.model.IsuTypeSolution;
import com.ngtesting.platform.service.intf.IssueTypeSolutionService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class IssueTypeSolutionServiceImpl extends BaseServiceImpl implements IssueTypeSolutionService {
    @Autowired
    IssueTypeSolutionDao solutionDao;

    @Autowired
    IssueTypeDao typeDao;

    @Override
    public List<IsuTypeSolution> list(Integer orgId) {
        List<IsuTypeSolution> ls = solutionDao.list(orgId);

        return ls;
    }

//    @Override
//    public IsuTypeSolution get(Integer id, Integer orgId) {
//        return solutionDao.get(id, orgId);
//    }

    @Override
    public IsuTypeSolution getDetail(Integer id, Integer orgId) {
        return solutionDao.getDetail(id, orgId);
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
        return count > 0;
    }

    @Override
    public Boolean addType(Integer typeId, Integer solutionId, Integer orgId) {
        Integer count = solutionDao.addType(typeId, solutionId, orgId);
        return count > 0;
    }

    @Override
    public Boolean removeType(Integer typeId, Integer solutionId, Integer orgId) {
        Integer count = solutionDao.removeType(typeId, solutionId, orgId);
        return count > 0;
    }

    @Override
    public Boolean addAll(Integer solutionId, Integer orgId) {
        List<IsuType> types = typeDao.listNotInSolution(solutionId, orgId);

        Integer count = solutionDao.addAll(types, solutionId, orgId);
        return count > 0;
    }

    @Override
    public Boolean removeAll(Integer solutionId, Integer orgId) {
        Integer count = solutionDao.removeAll(solutionId, orgId);
        return count > 0;
    }

    // For Project
    @Override
    public IsuTypeSolution getByProject(Integer projectId, Integer orgId) {
        return solutionDao.getByProject(projectId, orgId);
    }

    @Override
    public void setByProject(Integer solutionId, Integer projectId, Integer orgId) {
        solutionDao.setByProject(solutionId, projectId, orgId);
    }

    @Override
    public Boolean setDefault(Integer id, Integer orgId) {
        solutionDao.removeDefault(orgId);

        Integer count = solutionDao.setDefault(id, orgId);
        return count > 0;
    }

}
