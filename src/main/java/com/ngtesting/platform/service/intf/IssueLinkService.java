package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.model.IsuLinkReason;

import java.util.List;

public interface IssueLinkService extends BaseService {
    Boolean link(Integer srcIssueId, Integer dictIssueId, Integer reasonId, String reason, Integer prjId);
    List<IsuLinkReason> listLinkReason();
}
