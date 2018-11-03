package com.ngtesting.platform.service;

import com.ngtesting.platform.model.IsuPage;
import com.ngtesting.platform.model.IsuPageElement;
import com.ngtesting.platform.model.IsuPageTab;

import java.util.List;

public interface IssuePageService extends BaseService {

    List<IsuPage> list(Integer orgId);

    IsuPage get(Integer pageId, Integer orgId);

    IsuPage save(IsuPage page, Integer orgId);

    boolean delete(Integer id, Integer orgId);

    void addTab(IsuPageTab tab);

    void addField(IsuPageElement element);
}
