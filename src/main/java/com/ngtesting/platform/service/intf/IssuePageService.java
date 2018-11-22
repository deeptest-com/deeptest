package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.model.IsuPage;

import java.util.List;

public interface IssuePageService extends BaseService {

    List<IsuPage> list(Integer orgId);

    IsuPage get(Integer pageId, Integer orgId);

    IsuPage save(IsuPage page, Integer orgId);

    boolean delete(Integer id, Integer orgId);

    Boolean setDefault(Integer id, Integer orgId);
}
