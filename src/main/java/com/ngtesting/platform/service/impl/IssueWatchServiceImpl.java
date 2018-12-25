package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.IssueDao;
import com.ngtesting.platform.dao.IssueWatchDao;
import com.ngtesting.platform.model.IsuIssue;
import com.ngtesting.platform.model.TstUser;
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
    @Autowired
    IssueDao issueDao;


    @Override
    public List<TstUser> list(Integer issueId) {
        return issueWatchDao.list(issueId);
    }

    @Override
    public List<TstUser> search(Integer issueId, Integer prjId, String keywords, List<Integer> exceptIds, TstUser user) {
        IsuIssue issue = issueDao.get(issueId, user.getId(), user.getDefaultPrjId());
        if (issue == null) {
            return null;
        }

        List<TstUser> watchedUsers = issueWatchDao.list(issueId);

        for (TstUser u: watchedUsers) {
            if (!exceptIds.contains(u.getId())) {
                exceptIds.add(u.getId());
            }
        }

        return issueWatchDao.search(issueId, prjId, keywords, exceptIds);
    }

    @Override
    public Boolean remove(Integer id, Integer issueId, TstUser user) {
        IsuIssue issue = issueDao.get(issueId, user.getId(), user.getDefaultPrjId());
        if (issue == null) {
            return false;
        }
        issueWatchDao.remove(id);

        return true;
    }

    @Override
    public Boolean batchSave(Integer issueId, List<Integer> userIds, TstUser user) {
        IsuIssue issue = issueDao.get(issueId, user.getId(), user.getDefaultPrjId());
        if (issue == null) {
            return false;
        }

        issueWatchDao.batchSave(issueId, userIds);

        return true;
    }

    @Override
    public Boolean watch(Integer issueId, TstUser user, Boolean status) {
        IsuIssue issue = issueDao.get(issueId, user.getId(), user.getDefaultPrjId());
        if (issue == null) {
            return false;
        }

        if (status) {
            issueWatchDao.watch(issueId, user.getId());
        } else {
            issueWatchDao.unwatch(issueId, user.getId());
        }

        return true;
    }

}

