package com.ngtesting.platform.service;

import com.ngtesting.platform.model.IsuCustomField;

import java.util.List;

public interface IssueCustomFieldService extends BaseService {

    List<IsuCustomField> list(Integer orgId);

    IsuCustomField get(Integer id, Integer orgId);

    IsuCustomField save(IsuCustomField vo, Integer orgId);

    Boolean delete(Integer id, Integer orgId);

    Boolean changeOrderPers(Integer id, String act, Integer orgId);

    String getLastUnusedColumn(Integer orgId);

    List<String> listType();

    List<String> listFormat();
}
