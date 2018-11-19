package com.ngtesting.platform.service;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.model.TstCase;
import com.ngtesting.platform.model.TstUser;

import java.util.List;

public interface CaseService extends BaseService {
	List<TstCase> query(Integer projectId);

	List<TstCase> queryForSuiteSelection(Integer projectId, Integer suiteId);

	List<TstCase> queryForTaskSelection(Integer projectId, Integer runId);

	TstCase getDetail(Integer caseId, Integer prjId);

	TstCase rename(JSONObject json, TstUser user);
	TstCase rename(Integer id, String name, Integer pId, Integer projectId, TstUser user);

	TstCase move(JSONObject json, TstUser user);

	TstCase update(JSONObject json, TstUser user);
	Integer delete(Integer id, TstUser user);

    TstCase saveField(JSONObject json, TstUser user);

    TstCase changeContentType(Integer id, String contentType, TstUser user);
    TstCase reviewResult(Integer id, Boolean pass, TstUser user);

	void createSample(Integer projectId, TstUser user);
	boolean cloneStepsAndChildrenPers(TstCase testcase, TstCase src);
	void loadNodeTree(TstCase po);
	Integer getChildMaxOrderNumb(Integer parentId);

	void genVos(List<TstCase> pos, List<Integer> selectIds);

	void genVo(TstCase po, List<Integer> selectIds);

    List<String> genExtPropList();
}
