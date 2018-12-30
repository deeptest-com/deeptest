package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.model.TstUser;

public interface CaseInTaskAttachmentService extends BaseService {

    Boolean save(Integer caseId, String name, String path, TstUser user);
    Boolean delete(Integer id, TstUser user);

}
