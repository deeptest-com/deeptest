package com.ngtesting.platform.service;

import com.ngtesting.platform.model.IsuIssue;

import java.util.List;

public interface IssueService extends BaseService {
	IsuIssue get(Integer id, Integer orgId);

	void genVos(List<IsuIssue> pos);

	void genVo(IsuIssue po);
}
