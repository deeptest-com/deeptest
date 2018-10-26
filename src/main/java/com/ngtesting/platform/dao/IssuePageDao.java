package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.IsuPage;

import java.util.List;

public interface IssuePageDao {

    List<IsuPage> list(Integer orgId);
}
