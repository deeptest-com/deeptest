package com.ngtesting.platform.service;

import com.ngtesting.platform.model.TstOrg;

import java.util.List;


public interface OrgService extends BaseService {

	List<TstOrg> list(Integer userId, String keywords, Boolean disabled);
	List<TstOrg> listByUser(Integer userId);

	TstOrg get(Integer id);

	Boolean delete(Integer id);

	void genVos(List<TstOrg> pos, Integer userId);

    TstOrg save(TstOrg vo, Integer userId);

}
