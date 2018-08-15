package com.ngtesting.platform.service;

import com.ngtesting.platform.model.TstCustomFieldOption;

import java.util.List;

public interface CustomFieldOptionService extends BaseService {
    List<TstCustomFieldOption> listVos(Integer fieldId);
    TstCustomFieldOption save(TstCustomFieldOption option, Integer orgId);
    Boolean delete(Integer id, Integer orgId);
    Boolean changeOrder(Integer id, String act, Integer fieldId, Integer orgId);

}
