package com.ngtesting.platform.service;

import com.ngtesting.platform.model.TstCaseAttachment;
import com.ngtesting.platform.model.TstUser;

import java.util.List;

public interface CaseAttachmentService extends BaseService {

    void uploadAttachmentPers(Long caseId, String name, String path, TstUser user);
    void removeAttachmentPers(Long id, TstUser user);

    List<TstCaseAttachment> listByCase(Long caseId);

    List<TstCaseAttachment> genVos(List<TstCaseAttachment> pos);
    TstCaseAttachment genVo(TstCaseAttachment attachment);
}
