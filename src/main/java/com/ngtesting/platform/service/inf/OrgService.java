package com.ngtesting.platform.service.inf;

import com.ngtesting.platform.model.TstOrg;
import com.ngtesting.platform.model.TstUser;

import java.util.List;


public interface OrgService extends BaseService {

	List<TstOrg> list(String keywords, String disabled, Integer userId);
	List<TstOrg> listVo(String keywords, String disabled, Integer id);

	TstOrg getDetail(Integer id);

    Boolean disable(Integer id);
	Boolean delete(Integer id);

	List<TstOrg> genVos(List<TstOrg> pos, Integer userId);

	TstOrg genVo(TstOrg po);

	void createDefaultBasicDataPers(TstUser user);

    TstOrg save(TstOrg vo, Integer userId);

	void setDefaultPers(Integer orgId, TstUser user);

}
