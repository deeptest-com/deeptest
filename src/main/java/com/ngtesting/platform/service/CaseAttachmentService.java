package com.ngtesting.platform.service;

import com.ngtesting.platform.entity.TestCaseAttachment;
import com.ngtesting.platform.vo.TestCaseAttachmentVo;
import com.ngtesting.platform.vo.UserVo;

import java.util.List;

public interface CaseAttachmentService extends BaseService {

    void uploadAttachmentPers(Long caseId, String name, String path, UserVo user);
    void removeAttachmentPers(Long id, UserVo user);

    List<TestCaseAttachmentVo> listByCase(Long caseId);

    List<TestCaseAttachmentVo> genVos(List<TestCaseAttachment> pos);
    TestCaseAttachmentVo genVo(TestCaseAttachment attachment);
}
