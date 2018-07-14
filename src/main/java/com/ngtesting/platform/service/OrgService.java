package com.ngtesting.platform.service;

import com.ngtesting.platform.model.TstOrg;

import java.util.List;


public interface OrgService extends BaseService {

	List<TstOrg> list(String keywords, String disabled, Integer userId);
	List<TstOrg> listByUser(Integer userId);

	TstOrg getDetail(Integer id);

    Boolean disable(Integer id);
	Boolean delete(Integer id);

	List<TstOrg> genVos(List<TstOrg> pos, Integer userId);

    TstOrg save(TstOrg vo, Integer userId);

}
