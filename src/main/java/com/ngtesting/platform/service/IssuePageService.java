package com.ngtesting.platform.service;

import com.ngtesting.platform.model.IsuPage;

import java.util.List;

public interface IssuePageService extends BaseService {

    List<IsuPage> list(Integer orgId);
}
