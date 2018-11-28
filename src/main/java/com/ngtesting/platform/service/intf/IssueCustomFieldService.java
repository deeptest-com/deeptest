package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.model.IsuCustomField;

import java.util.List;
import java.util.Map;

public interface IssueCustomFieldService extends BaseService {

    List<IsuCustomField> list(Integer orgId);

    IsuCustomField get(Integer id, Integer orgId);

    IsuCustomField save(IsuCustomField vo, Integer orgId);

    Boolean delete(Integer id, Integer orgId);

    Boolean changeOrderPers(Integer id, String act, Integer orgId);

    String getLastUnusedColumn(Integer orgId);

    Map<String, Map> fetchInputMap();

    Map inputMap();

    Map typeMap();

    List<String> listFormat();
}
