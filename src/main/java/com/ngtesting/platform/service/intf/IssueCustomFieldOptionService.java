package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.model.CustomFieldOption;

import java.util.List;

public interface IssueCustomFieldOptionService extends BaseService {
    List<CustomFieldOption> list(Integer fieldId, Integer orgId);
    CustomFieldOption save(CustomFieldOption option, Integer orgId);
    Boolean delete(Integer id, Integer fieldId, Integer orgId);
    Boolean changeOrder(Integer id, String act, Integer fieldId, Integer orgId);

    Boolean setDefault(Integer id, Integer fieldId, Integer orgId);

    CustomFieldOption get(Integer id, Integer fieldId, Integer orgId);
}
