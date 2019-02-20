package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.IssueDao;
import com.ngtesting.platform.dao.IssueWatchDao;
import com.ngtesting.platform.model.IsuIssue;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.IssueHistoryService;
import com.ngtesting.platform.service.intf.IssueWatchService;
import com.ngtesting.platform.utils.MsgUtil;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Map;

@Service
public class IssueWatchServiceImpl extends BaseServiceImpl implements IssueWatchService {
    Log logger = LogFactory.getLog(IssueWatchServiceImpl.class);

    @Autowired
    IssueWatchDao issueWatchDao;
    @Autowired
    IssueDao issueDao;

    @Autowired
    IssueHistoryService issueHistoryService;

    @Override
    public List<Map> list(Integer issueId) {
        return issueWatchDao.listByIssueId(issueId);
    }

    @Override
    public List<TstUser> search(Integer issueId, Integer prjId, String keywords, List<Integer> exceptIds, TstUser user) {
        IsuIssue issue = issueDao.get(issueId, user.getId(), user.getDefaultPrjId());
        if (issue == null) {
            return null;
        }

        List<Map> watchedUsers = issueWatchDao.listByIssueId(issueId);

        for (Map w: watchedUsers) {
            Integer id = Integer.valueOf(w.get("userId").toString());
            if (!exceptIds.contains(id)) {
                exceptIds.add(id);
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

        issueHistoryService.saveHistory(user, MsgUtil.MsgAction.removeWatch, issueId, null);

        return true;
    }

    @Override
    public Boolean batchWatch(Integer issueId, List<Integer> userIds, TstUser user) {
        IsuIssue issue = issueDao.get(issueId, user.getId(), user.getDefaultPrjId());
        if (issue == null) {
            return false;
        }

        issueWatchDao.batchWatch(issueId, userIds);

        issueHistoryService.saveHistory(user, MsgUtil.MsgAction.changeWatch, issueId, null);

        return true;
    }

    @Override
    public Boolean watch(Integer issueId, TstUser user, Boolean status) {
        IsuIssue issue = issueDao.get(issueId, user.getId(), user.getDefaultPrjId());
        if (issue == null) {
            return false;
        }

        MsgUtil.MsgAction act;
        if (status) {
            issueWatchDao.watch(issueId, user.getId());

            act = MsgUtil.MsgAction.watch;
        } else {
            issueWatchDao.unwatch(issueId, user.getId());

            act = MsgUtil.MsgAction.unwatch;
        }

        issueHistoryService.saveHistory(user, act, issueId, null);

        return true;
    }

}

