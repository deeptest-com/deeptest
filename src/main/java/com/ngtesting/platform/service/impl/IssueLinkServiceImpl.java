package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.IssueDao;
import com.ngtesting.platform.dao.IssueLinkDao;
import com.ngtesting.platform.model.IsuLinkReason;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.IssueHistoryService;
import com.ngtesting.platform.service.intf.IssueLinkService;
import com.ngtesting.platform.utils.MsgUtil;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.Arrays;
import java.util.LinkedList;
import java.util.List;

@Service
public class IssueLinkServiceImpl extends BaseServiceImpl implements IssueLinkService {
    @Autowired
    IssueLinkDao issueLinkDao;

    @Autowired
    IssueDao issueDao;
    @Autowired
    IssueHistoryService issueHistoryService;

    @Override
    public Boolean link(Integer srcIssueId, Integer dictIssueId, Integer reasonId,
                        String reasonName, TstUser user) {
        List<Integer> ls = issueDao.getByIds(new LinkedList<>(Arrays.asList(srcIssueId, dictIssueId)));
        if (ls.size() < 2) {
            return false;
        }

        issueLinkDao.link(srcIssueId, dictIssueId, reasonId, reasonName);

        issueHistoryService.saveHistory(user, MsgUtil.MsgAction.link, srcIssueId, "IS-" + dictIssueId);

        return true;
    }

    @Override
    public List<IsuLinkReason> listLinkReason() {
        return issueLinkDao.listLinkReason();
    }

}

