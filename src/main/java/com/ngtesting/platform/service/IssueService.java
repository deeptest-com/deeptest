package com.ngtesting.platform.service;

import com.ngtesting.platform.model.IsuIssue;

import java.util.List;

public interface IssueService extends BaseService {
	List<IsuIssue> queryByProject(Integer projectId, String columns);
	List<IsuIssue> queryByJql(String query, String columns);

	IsuIssue getById(Integer id);

	void genVos(List<IsuIssue> pos);

	void genVo(IsuIssue po);
}
