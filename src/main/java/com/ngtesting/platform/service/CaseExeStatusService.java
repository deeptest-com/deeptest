package com.ngtesting.platform.service;

import com.ngtesting.platform.model.TstCaseExeStatus;

import java.util.List;

public interface CaseExeStatusService extends BaseService {
	List<TstCaseExeStatus> list(Integer orgId);

	TstCaseExeStatus save(TstCaseExeStatus vo, Long orgId);
	boolean delete(Integer id);

	boolean changeOrderPers(Integer id, String act);

//    void createDefaultBasicDataPers(Long id);
}
