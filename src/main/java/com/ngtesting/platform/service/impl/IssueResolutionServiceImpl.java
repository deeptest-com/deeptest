package com.ngtesting.platform.service.impl;


import com.ngtesting.platform.dao.IssueResolutionDao;
import com.ngtesting.platform.model.IsuResolution;
import com.ngtesting.platform.service.IssueResolutionService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

@Service
public class IssueResolutionServiceImpl extends BaseServiceImpl implements IssueResolutionService {
    @Autowired
    IssueResolutionDao issueResolutionDao;

    @Override
    public List<IsuResolution> list(Integer orgId) {
        List<IsuResolution> ls = issueResolutionDao.list(orgId);

        return ls;
    }

    @Override
    public List<IsuResolution> list(Integer orgId, Integer prjId) {
        List<IsuResolution> ls = issueResolutionDao.list(orgId);

        return ls;
    }

    @Override
    public IsuResolution get(Integer id, Integer orgId) {
        return issueResolutionDao.get(id, orgId);
    }

    @Override
    public IsuResolution save(IsuResolution vo, Integer orgId) {

        if (vo.getId() == null) {
            Integer maxOrder = issueResolutionDao.getMaxOrdrNumb(orgId);
            if (maxOrder == null) {
                maxOrder = 0;
            }
            vo.setOrdr(maxOrder + 10);

            vo.setOrgId(orgId);
            issueResolutionDao.save(vo);
        } else {
            Integer count = issueResolutionDao.update(vo);
            if (count == 0) {
                return null;
            }
        }

        return vo;
    }

    @Override
    public Boolean delete(Integer id, Integer orgId) {
        Integer count = issueResolutionDao.delete(id, orgId);
        if (count == 0) {
            return false;
        }

        return true;
    }

    @Override
    @Transactional
    public Boolean setDefault(Integer id, Integer orgId) {
        issueResolutionDao.removeDefault(orgId);

        Integer count = issueResolutionDao.setDefault(id, orgId);
        if (count == 0) {
            return false;
        }
        return true;
    }

    @Override
    @Transactional
    public Boolean changeOrder(Integer id, String act, Integer orgId) {
        IsuResolution curr = issueResolutionDao.get(id, orgId);
        if (curr == null) {
            return false;
        }

        IsuResolution neighbor = null;
        if ("up".equals(act)) {
            neighbor = issueResolutionDao.getPrev(curr.getOrdr(), orgId);
        } else if ("down".equals(act)) {
            neighbor = issueResolutionDao.getNext(curr.getOrdr(), orgId);
        }
        if (neighbor == null) {
            return false;
        }

        Integer currOrder = curr.getOrdr();
        Integer neighborOrder = neighbor.getOrdr();
        issueResolutionDao.setOrder(id, neighborOrder, orgId);
        issueResolutionDao.setOrder(neighbor.getId(), currOrder, orgId);

        return true;
    }
}

