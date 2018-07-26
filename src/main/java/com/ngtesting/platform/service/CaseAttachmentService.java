package com.ngtesting.platform.service;

import com.ngtesting.platform.model.TstUser;

public interface CaseAttachmentService extends BaseService {

    void uploadAttachmentPers(Integer caseId, String name, String path, TstUser user);
    void removeAttachmentPers(Integer id, TstUser user);
}
