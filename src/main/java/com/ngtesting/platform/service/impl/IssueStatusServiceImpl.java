package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.IssueStatusDao;
import com.ngtesting.platform.model.IsuStatus;
import com.ngtesting.platform.model.IsuStatusCategoryDefine;
import com.ngtesting.platform.service.intf.IssueStatusService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

@Service
public class IssueStatusServiceImpl extends BaseServiceImpl implements IssueStatusService {
    @Autowired
    IssueStatusDao issueStatusDao;

    @Override
    public List<IsuStatus> list(Integer orgId) {
        List<IsuStatus> ls = issueStatusDao.list(orgId);

        return ls;
    }

    @Override
    public List<IsuStatusCategoryDefine> listCategory() {
        List<IsuStatusCategoryDefine> ls = issueStatusDao.listCategory();

        return ls;
    }

    @Override
    public IsuStatus get(Integer id, Integer orgId) {
        return issueStatusDao.get(id, orgId);
    }

    @Override
    public IsuStatus save(IsuStatus vo, Integer orgId) {
        IsuStatusCategoryDefine cate = issueStatusDao.getCategoryById(vo.getCategoryId());
        vo.setFinalVal(cate.getFinalVal());

        if (vo.getId() == null) {
            Integer maxOrder = issueStatusDao.getMaxOrdrNumb(orgId);
            if (maxOrder == null) {
                maxOrder = 0;
            }
            vo.setOrdr(maxOrder + 10);

            vo.setOrgId(orgId);
            vo.setBuildIn(Boolean.FALSE);
            issueStatusDao.save(vo);
        } else {
            Integer count = issueStatusDao.update(vo);
            if (count == 0) {
                return null;
            }
        }

        return vo;
    }

    @Override
    public Boolean delete(Integer id, Integer orgId) {
        Integer count = issueStatusDao.delete(id, orgId);
        return count > 0;
    }

    @Override
    @Transactional
    public Boolean setDefault(Integer id, Integer orgId) {
        issueStatusDao.removeDefault(orgId);

        Integer count = issueStatusDao.setDefault(id, orgId);
        return count > 0;
    }

    @Override
    @Transactional
    public Boolean changeOrder(Integer id, String act, Integer orgId) {
        IsuStatus curr = issueStatusDao.get(id, orgId);
        if (curr == null) {
            return false;
        }

        IsuStatus neighbor = null;
        if ("up".equals(act)) {
            neighbor = issueStatusDao.getPrev(curr.getOrdr(), orgId);
        } else if ("down".equals(act)) {
            neighbor = issueStatusDao.getNext(curr.getOrdr(), orgId);
        }
        if (neighbor == null) {
            return false;
        }

        Integer currOrder = curr.getOrdr();
        Integer neighborOrder = neighbor.getOrdr();
        issueStatusDao.setOrder(id, neighborOrder, orgId);
        issueStatusDao.setOrder(neighbor.getId(), currOrder, orgId);

        return true;
    }
}

