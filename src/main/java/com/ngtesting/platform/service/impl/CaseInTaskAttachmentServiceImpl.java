package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.dao.CaseDao;
import com.ngtesting.platform.dao.CaseInTaskAttachmentDao;
import com.ngtesting.platform.model.TstCase;
import com.ngtesting.platform.model.TstCaseAttachment;
import com.ngtesting.platform.model.TstCaseInTaskAttachment;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.CaseHistoryService;
import com.ngtesting.platform.service.intf.CaseInTaskAttachmentService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
public class CaseInTaskAttachmentServiceImpl extends BaseServiceImpl implements CaseInTaskAttachmentService {
    @Autowired
    CaseInTaskAttachmentDao caseInTaskAttachmentDao;
    @Autowired
    CaseHistoryService caseHistoryService;
    @Autowired
    CaseDao caseDao;

    @Override
    @Transactional
    public Boolean save(Integer caseId, String name, String path, TstUser user) {
        TstCaseInTaskAttachment attach = new TstCaseInTaskAttachment(name, path, caseId, user.getId());
        caseInTaskAttachmentDao.save(attach);
//        caseHistoryService.saveHistory(user, Constant.EntityAct.attachment_upload, testCase, name);
        return true;
    }

    @Override
    @Transactional
    public Boolean delete(Integer id, TstUser user) {
        caseInTaskAttachmentDao.delete(id, user.getId());
//        caseHistoryService.saveHistory(user, Constant.EntityAct.attachment_delete, testCase, attach.getName());

        return true;
    }

}

