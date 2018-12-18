package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.IssuePageDao;
import com.ngtesting.platform.dao.IssuePageElementDao;
import com.ngtesting.platform.dao.IssueWatchDao;
import com.ngtesting.platform.model.IsuTag;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.IssueCommentsService;
import com.ngtesting.platform.service.intf.IssueWatchService;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class IssueWatchServiceImpl extends BaseServiceImpl implements IssueWatchService {
    Log logger = LogFactory.getLog(IssueWatchServiceImpl.class);

    @Autowired
    IssueWatchDao issueWatchDao;

    @Override
    public List<TstUser> list(Integer issueId) {
        return issueWatchDao.list(issueId);
    }

    @Override
    public List<TstUser> search(Integer issueId, Integer prjId, String keywords, List<Integer> exceptIds) {
        List<TstUser> watchedIds = issueWatchDao.list(issueId);

        for (TstUser user: watchedIds) {
            if (!exceptIds.contains(user.getId())) {
                exceptIds.add(user.getId());
            }
        }

        return issueWatchDao.search(issueId, prjId, keywords, exceptIds);
    }

    @Override
    public void remove(Integer id) {
        issueWatchDao.remove(id);
    }

    @Override
    public void batchSave(Integer issueId, List<Integer> userIds) {
        issueWatchDao.batchSave(issueId, userIds);
    }

    @Override
    public void watch(Integer id, TstUser user, Boolean status) {
        if (status) {
            issueWatchDao.watch(id, user.getId());
        } else {
            issueWatchDao.unwatch(id, user.getId());
        }
    }

}

