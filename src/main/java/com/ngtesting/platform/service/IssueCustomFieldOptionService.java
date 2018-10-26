package com.ngtesting.platform.service;

import com.ngtesting.platform.model.IsuCustomFieldOption;

import java.util.List;

public interface IssueCustomFieldOptionService extends BaseService {
    List<IsuCustomFieldOption> listVos(Integer fieldId);
    IsuCustomFieldOption save(IsuCustomFieldOption option, Integer orgId);
    Boolean delete(Integer id, Integer orgId);
    Boolean changeOrder(Integer id, String act, Integer fieldId, Integer orgId);
}
