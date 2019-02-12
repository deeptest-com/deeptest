package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.model.IsuQuery;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.tql.query.builder.support.model.JsonRule;

import java.util.List;

public interface IssueQueryService extends BaseService {
	List<IsuQuery> list(Integer orgId, Integer userId, String keywords);
	List<IsuQuery> listRecentQuery(Integer orgId, Integer userId);

	IsuQuery get(Integer queryId, Integer id);

	IsuQuery save(String queryName, JsonRule rule, TstUser user);
	Integer update(IsuQuery vo, TstUser user);
    Integer delete(Integer id, TstUser user);

    Integer updateUseTime(IsuQuery query, TstUser user);
}
