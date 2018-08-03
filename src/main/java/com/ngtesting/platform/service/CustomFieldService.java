package com.ngtesting.platform.service;

import com.ngtesting.platform.model.TstCustomField;

import java.util.List;

public interface CustomFieldService extends BaseService {
	List<TstCustomField> list(Integer orgId);
    List<TstCustomField> listForCaseByProject(Integer orgId, Integer projectId);

	TstCustomField get(Integer customFieldId);

	TstCustomField save(TstCustomField vo, Integer orgId);
	boolean delete(Integer id);

	List<String> listApplyTo();

	List<String> listType();

	List<String> listFormat();

	boolean changeOrderPers(Integer id, String act, Integer orgId);

    String getLastUnusedColumn(Integer orgId);

}
