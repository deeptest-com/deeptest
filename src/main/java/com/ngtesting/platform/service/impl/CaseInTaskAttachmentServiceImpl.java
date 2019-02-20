package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.CaseDao;
import com.ngtesting.platform.dao.CaseInTaskAttachmentDao;
import com.ngtesting.platform.model.TstCaseInTaskAttachment;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.CaseInTaskAttachmentService;
import com.ngtesting.platform.service.intf.CaseInTaskHistoryService;
import com.ngtesting.platform.utils.MsgUtil;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
public class CaseInTaskAttachmentServiceImpl extends BaseServiceImpl implements CaseInTaskAttachmentService {
    @Autowired
    CaseInTaskAttachmentDao caseInTaskAttachmentDao;
    @Autowired
    CaseInTaskHistoryService caseInTaskHistoryService;
    @Autowired
    CaseDao caseDao;

    @Override
    @Transactional
    public Boolean save(Integer caseId, String name, String path, TstUser user) {
        TstCaseInTaskAttachment attach = new TstCaseInTaskAttachment(name, path, caseId, user.getId());
        caseInTaskAttachmentDao.save(attach);
        caseInTaskHistoryService.saveHistory(user, MsgUtil.MsgAction.attachment_upload, caseId, name);
        return true;
    }

    @Override
    @Transactional
    public Boolean delete(Integer id, TstUser user) {
        TstCaseInTaskAttachment attach = caseInTaskAttachmentDao.get(id);

        caseInTaskAttachmentDao.delete(id, user.getId());
        caseInTaskHistoryService.saveHistory(user, MsgUtil.MsgAction.attachment_delete,
                attach.getCaseInTaskId(), attach.getName());

        return true;
    }

}

