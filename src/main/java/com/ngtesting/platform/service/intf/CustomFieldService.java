package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.model.CustomField;

import java.util.List;
import java.util.Map;

public interface CustomFieldService extends BaseService {

    List<CustomField> list(Integer orgId, String applyTo, String keywords);

    CustomField getDetail(Integer id, Integer orgId);

    CustomField save(CustomField vo, Integer orgId);

    Boolean delete(Integer id, Integer orgId);

    Boolean changeOrderPers(Integer id, String act, Integer orgId, String applyTo);

    String getLastUnusedColumn(Integer orgId);

    List<String> listApplyTo();

    Map<String, Map> fetchInputMap();

    Map inputMap();

    Map typeMap();

    List<String> listFormat();

    Map<String, Object> fetchProjectFieldForCase(Integer orgId, Integer projectId);
}
