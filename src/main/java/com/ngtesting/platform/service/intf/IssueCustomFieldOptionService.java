package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.model.IsuCustomFieldOption;

import java.util.List;

public interface IssueCustomFieldOptionService extends BaseService {
    List<IsuCustomFieldOption> listVos(Integer fieldId);
    IsuCustomFieldOption save(IsuCustomFieldOption option, Integer orgId);
    Boolean delete(Integer id, Integer fieldId, Integer orgId);
    Boolean changeOrder(Integer id, String act, Integer fieldId, Integer orgId);

    Boolean setDefault(Integer id, Integer fieldId, Integer orgId);

    IsuCustomFieldOption get(Integer id, Integer fieldId, Integer orgId);
}
