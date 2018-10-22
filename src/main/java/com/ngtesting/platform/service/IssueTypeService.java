package com.ngtesting.platform.service;

import com.ngtesting.platform.model.IsuType;

import java.util.List;

public interface IssueTypeService extends BaseService {

    List<IsuType> list(Integer orgId);
}
