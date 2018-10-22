package com.ngtesting.platform.service;

import com.ngtesting.platform.model.IsuPriority;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

public interface IssuePriorityService extends BaseService {
	List<IsuPriority> list(Integer orgId);

    List<IsuPriority> list(Integer orgId, Integer prjId);

    IsuPriority get(Integer id, Integer orgId);

	IsuPriority save(IsuPriority vo, Integer orgId);

	Boolean delete(Integer id, Integer orgId);

	@Transactional
	Boolean setDefault(Integer id, Integer orgId);

	@Transactional
	Boolean changeOrder(Integer id, String act, Integer orgId);
}
