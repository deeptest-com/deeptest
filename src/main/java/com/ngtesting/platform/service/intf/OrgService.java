package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.model.TstOrg;
import com.ngtesting.platform.model.TstUser;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;


public interface OrgService extends BaseService {

	List<TstOrg> list(Integer userId, String keywords, Boolean disabled);
	List<TstOrg> listByUser(Integer userId);

	TstOrg get(Integer id);

    @Transactional
    TstOrg update(TstOrg vo, TstUser user);

    Boolean delete(Integer id, TstUser user);

	void changeDefaultOrg(TstUser user, Integer orgId);
    void setUserDefaultOrgPrjToNullForDelete(Integer orgId);

    void genVos(List<TstOrg> pos, Integer userId);

    TstOrg save(TstOrg vo, TstUser user);
}
