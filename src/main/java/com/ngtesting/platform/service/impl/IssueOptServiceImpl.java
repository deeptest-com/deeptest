package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.IssueOptDao;
import com.ngtesting.platform.model.IsuComments;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.IssueCommentsService;
import com.ngtesting.platform.service.intf.IssueOptService;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class IssueOptServiceImpl extends BaseServiceImpl implements IssueOptService {
    Log logger = LogFactory.getLog(IssueOptServiceImpl.class);

    @Autowired
    IssueOptDao issueOptDao;

    @Autowired
    IssueCommentsService issueCommentsService;

    @Override
    public void statusTran(Integer id, Integer dictStatusId, Integer projectId) {
        issueOptDao.statusTran(id, dictStatusId, projectId);
    }

    @Override
    public void assign(Integer id, TstUser user, String content) {
        IsuComments po = new IsuComments(id, "修改经办人", content);
        issueCommentsService.save(po, user);

        issueOptDao.assign(id, user.getId(), user.getDefaultPrjId());
    }

}

