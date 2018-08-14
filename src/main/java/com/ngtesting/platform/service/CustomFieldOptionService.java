package com.ngtesting.platform.service;

import com.ngtesting.platform.model.TstCustomFieldOption;

import java.util.List;

public interface CustomFieldOptionService extends BaseService {
    List<TstCustomFieldOption> listVos(Integer fieldId);
    TstCustomFieldOption save(TstCustomFieldOption option, Integer orgId);
    boolean delete(Integer id, Integer orgId);
    boolean changeOrderPers(Integer id, String act, Integer fieldId, Integer orgId);

}
