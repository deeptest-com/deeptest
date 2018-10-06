package com.ngtesting.platform.service;

import com.ngtesting.platform.model.TstOrg;
import com.ngtesting.platform.model.TstUser;

import java.util.List;


public interface OrgService extends BaseService {

	List<TstOrg> list(Integer userId, String keywords, Boolean disabled);
	List<TstOrg> listByUser(Integer userId);

	TstOrg get(Integer id);

	Boolean delete(Integer id, TstUser user);

	void changeDefaultOrg(TstUser user, Integer orgId);
    void setUserDefaultOrgPrjToNullForDelete(Integer orgId);

    void genVos(List<TstOrg> pos, Integer userId);

    TstOrg save(TstOrg vo, TstUser user);
}
