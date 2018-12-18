package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.model.IsuAttachment;
import com.ngtesting.platform.model.TstUser;

import java.util.List;

public interface IssueAttachmentService extends BaseService {

    List<IsuAttachment> query(Integer issueId);

    Boolean save(Integer caseId, String name, String path, TstUser user);
    Boolean delete(Integer id, TstUser user);

}
