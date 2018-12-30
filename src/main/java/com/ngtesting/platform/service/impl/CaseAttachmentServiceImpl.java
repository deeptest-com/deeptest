package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.dao.CaseAttachmentDao;
import com.ngtesting.platform.dao.CaseDao;
import com.ngtesting.platform.model.TstCase;
import com.ngtesting.platform.model.TstCaseAttachment;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.CaseAttachmentService;
import com.ngtesting.platform.service.intf.CaseHistoryService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
public class CaseAttachmentServiceImpl extends BaseServiceImpl implements CaseAttachmentService {
    @Autowired
    CaseAttachmentDao caseAttachmentDao;
    @Autowired
    CaseHistoryService caseHistoryService;
    @Autowired
    CaseDao caseDao;

    @Override
    @Transactional
    public Boolean save(Integer caseId, String name, String path, TstUser user) {
        TstCaseAttachment attach = new TstCaseAttachment(name, path, caseId, user.getId());
        caseAttachmentDao.save(attach);
        caseHistoryService.saveHistory(user, Constant.EntityAct.attachment_upload, caseId, name);
        return true;
    }

    @Override
    @Transactional
    public Boolean delete(Integer id, TstUser user) {
        TstCaseAttachment attach = caseAttachmentDao.get(id);

        caseAttachmentDao.delete(id, user.getId());
        caseHistoryService.saveHistory(user, Constant.EntityAct.attachment_delete,
                attach.getCaseId(), attach.getName());

        return true;
    }

}

