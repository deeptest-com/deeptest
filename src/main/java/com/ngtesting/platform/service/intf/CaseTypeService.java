package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.model.TstCaseType;

import java.util.List;

public interface CaseTypeService extends BaseService {
	List<TstCaseType> list(Integer orgId);

	TstCaseType get(Integer id, Integer orgId);

    TstCaseType getDefault(Integer orgId);

    TstCaseType save(TstCaseType vo, Integer orgId);
	Boolean delete(Integer id, Integer orgId);

	Boolean setDefault(Integer orgId, Integer orgId2);
	Boolean changeOrder(Integer id, String act, Integer orgId);

}
