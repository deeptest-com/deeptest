package com.ngtesting.platform.service;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.model.TstCase;
import com.ngtesting.platform.model.TstUser;

import java.util.List;

public interface CaseService extends BaseService {

	List<TstCase> query(Integer projectId);

	List<TstCase> queryForSuiteSelection(Integer projectId, Integer caseProjectId, Integer suiteId);

	List<TstCase> queryForTaskSelection(Integer projectId, Integer caseProjectId, Integer runId);

	TstCase getById(Integer caseId);

    TstCase renamePers(JSONObject json, TstUser user);
	void delete(Integer id, TstUser user);

	TstCase renamePers(Integer id, String name, Integer pId, Integer projectId, TstUser user);

	TstCase movePers(JSONObject json, TstUser user);

	void loadNodeTree(TstCase po);

	void createSample(Integer projectId, TstUser user);

	void create(TstCase testCase);

	TstCase update(JSONObject json, TstUser user);

	boolean cloneStepsAndChildrenPers(TstCase testcase, TstCase src);

    TstCase saveField(JSONObject json, TstUser user);

    Integer getChildMaxOrderNumb(Integer parentId);

    TstCase changeContentTypePers(Integer id, String contentType);

    TstCase reviewResult(Integer id, Boolean pass);

	void genVos(List<TstCase> pos, List<Integer> selectIds);

	void genVo(TstCase po, List<Integer> selectIds);

    List<String> genExtPropList();
}
