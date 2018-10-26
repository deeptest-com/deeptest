package com.ngtesting.platform.service;

import com.ngtesting.platform.model.TstCustomField;

import java.util.List;

public interface TestCustomFieldService extends BaseService {
	List<TstCustomField> list(Integer orgId);
    List<TstCustomField> listForCaseByProject(Integer orgId, Integer projectId);

	TstCustomField get(Integer customFieldId, Integer orgId);

	TstCustomField save(TstCustomField vo, Integer orgId);
	Boolean delete(Integer id, Integer orgId);
	Boolean changeOrderPers(Integer id, String act, Integer orgId);

	List<String> listApplyTo();

	List<String> listType();

	List<String> listFormat();

    String getLastUnusedColumn(Integer orgId);

}
