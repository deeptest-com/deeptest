package com.ngtesting.platform.service.inf;

import com.ngtesting.platform.model.TstCaseType;

import java.util.List;

public interface CaseTypeService extends BaseService {
	List<TstCaseType> list(Integer orgId);

	TstCaseType save(TstCaseType vo, Integer orgId);
	boolean delete(Integer id);

	List<TstCaseType> genVos(List<TstCaseType> pos);
	TstCaseType genVo(TstCaseType user);

	boolean setDefaultPers(Integer orgId, Integer orgId2);

	List<TstCaseType> listVos(Integer orgId);

	boolean changeOrderPers(Integer id, String act, Integer orgId);

//    void createDefaultBasicDataPers(Integer id);
}
