package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.IssuePageDao;
import com.ngtesting.platform.model.IsuPage;
import com.ngtesting.platform.service.intf.IssuePageService;
import com.ngtesting.platform.service.intf.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

@Service
public class IssuePageServiceImpl extends BaseServiceImpl implements IssuePageService {
    @Autowired
    UserService userService;

	@Autowired
    IssuePageDao pageDao;

    @Override
    public List<IsuPage> list(Integer orgId) {
        return pageDao.list(orgId);
    }

    @Override
    public IsuPage get(Integer pageId, Integer orgId) {
        return pageDao.get(pageId, orgId);
    }

    @Override
    @Transactional
    public IsuPage save(IsuPage vo, Integer orgId) {
        vo.setOrgId(orgId);

        if (vo.getId() == null) {
            pageDao.save(vo);
        } else {
            Integer count = pageDao.update(vo);
            if (count == 0) {
                return null;
            }
        }

        return vo;
    }

    @Override
    public boolean delete(Integer id, Integer orgId) {
        Integer count = pageDao.delete(id, orgId);

        return count > 0;
    }

    @Override
    @Transactional
    public Boolean setDefault(Integer id, Integer orgId) {
        pageDao.removeDefault(orgId);

        Integer count = pageDao.setDefault(id, orgId);
        return count > 0;
    }

}
