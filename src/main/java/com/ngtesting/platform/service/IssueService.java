package com.ngtesting.platform.service;

import com.ngtesting.platform.model.IsuIssue;

import java.util.List;

public interface IssueService extends BaseService {
	List<IsuIssue> queryByProject(Integer projectId);
	List<IsuIssue> queryByJql(String query);

	IsuIssue getById(Integer id);

	void genVos(List<IsuIssue> pos);

	void genVo(IsuIssue po);
}
