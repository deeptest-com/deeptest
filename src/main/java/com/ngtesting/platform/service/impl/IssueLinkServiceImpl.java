package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.IssueLinkDao;
import com.ngtesting.platform.model.IsuLinkReason;
import com.ngtesting.platform.service.intf.IssueLinkService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class IssueLinkServiceImpl extends BaseServiceImpl implements IssueLinkService {
    @Autowired
    IssueLinkDao issueLinkDao;

    @Override
    public void link(Integer srcIssueId, Integer dictIssueId, Integer reasonId, String reasonName) {
        issueLinkDao.link(srcIssueId, dictIssueId, reasonId, reasonName);
    }

    @Override
    public List<IsuLinkReason> listLinkReason() {
        return issueLinkDao.listLinkReason();
    }

}

