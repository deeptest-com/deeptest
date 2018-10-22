package com.ngtesting.platform.service.impl;


import com.ngtesting.platform.dao.IssueTypeDao;
import com.ngtesting.platform.model.IsuType;
import com.ngtesting.platform.service.IssueTypeService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

@Service
public class IssueTypeServiceImpl extends BaseServiceImpl implements IssueTypeService {
    @Autowired
    IssueTypeDao issueTypeDao;

    @Override
    public List<IsuType> list(Integer orgId) {
        List<IsuType> ls = issueTypeDao.list(orgId);

        return ls;
    }

    @Override
    public List<IsuType> list(Integer orgId, Integer prjId) {
        List<IsuType> ls = issueTypeDao.list(orgId);

        return ls;
    }

    @Override
    public IsuType get(Integer id, Integer orgId) {
        return issueTypeDao.get(id, orgId);
    }

    @Override
    public IsuType save(IsuType vo, Integer orgId) {

        if (vo.getId() == null) {
            Integer maxOrder = issueTypeDao.getMaxOrdrNumb(orgId);
            if (maxOrder == null) {
                maxOrder = 0;
            }
            vo.setOrdr(maxOrder + 10);

            vo.setOrgId(orgId);
            issueTypeDao.save(vo);
        } else {
            Integer count = issueTypeDao.update(vo);
            if (count == 0) {
                return null;
            }
        }

        return vo;
    }

    @Override
    public Boolean delete(Integer id, Integer orgId) {
        Integer count = issueTypeDao.delete(id, orgId);
        if (count == 0) {
            return false;
        }

        return true;
    }

    @Override
    @Transactional
    public Boolean setDefault(Integer id, Integer orgId) {
        issueTypeDao.removeDefault(orgId);

        Integer count = issueTypeDao.setDefault(id, orgId);
        if (count == 0) {
            return false;
        }
        return true;
    }

    @Override
    @Transactional
    public Boolean changeOrder(Integer id, String act, Integer orgId) {
        IsuType curr = issueTypeDao.get(id, orgId);
        if (curr == null) {
            return false;
        }

        IsuType neighbor = null;
        if ("up".equals(act)) {
            neighbor = issueTypeDao.getPrev(curr.getOrdr(), orgId);
        } else if ("down".equals(act)) {
            neighbor = issueTypeDao.getNext(curr.getOrdr(), orgId);
        }
        if (neighbor == null) {
            return false;
        }

        Integer currOrder = curr.getOrdr();
        Integer neighborOrder = neighbor.getOrdr();
        issueTypeDao.setOrder(id, neighborOrder, orgId);
        issueTypeDao.setOrder(neighbor.getId(), currOrder, orgId);

        return true;
    }
}
