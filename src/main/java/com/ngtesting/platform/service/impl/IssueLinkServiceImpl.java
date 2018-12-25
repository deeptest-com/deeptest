package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.IssueDao;
import com.ngtesting.platform.dao.IssueLinkDao;
import com.ngtesting.platform.model.IsuLinkReason;
import com.ngtesting.platform.service.intf.IssueLinkService;
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

    @Override
    public Boolean link(Integer srcIssueId, Integer dictIssueId, Integer reasonId, String reasonName, Integer prjId) {
        List<Integer> ls = issueDao.getByIds(new LinkedList<>(Arrays.asList(srcIssueId, dictIssueId)));
        if (ls.size() < 2) {
            return false;
        }

        issueLinkDao.link(srcIssueId, dictIssueId, reasonId, reasonName);

        return true;
    }

    @Override
    public List<IsuLinkReason> listLinkReason() {
        return issueLinkDao.listLinkReason();
    }

}

