package com.ngtesting.platform.service;

import com.ngtesting.platform.model.TstCaseAttachment;
import com.ngtesting.platform.model.TstUser;

import java.util.List;

public interface CaseAttachmentService extends BaseService {

    void uploadAttachmentPers(Integer caseId, String name, String path, TstUser user);
    void removeAttachmentPers(Integer id, TstUser user);

    List<TstCaseAttachment> listByCase(Integer caseId);

    List<TstCaseAttachment> genVos(List<TstCaseAttachment> pos);
    TstCaseAttachment genVo(TstCaseAttachment attachment);
}
