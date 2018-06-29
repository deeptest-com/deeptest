package com.ngtesting.platform.service.inf;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstCase;
import com.ngtesting.platform.model.TstCaseHistory;
import com.ngtesting.platform.model.TstUser;

import java.util.List;

public interface CaseService extends BaseService {

	List<TstCase> query(Integer projectId);

	List<TstCase> queryForSuiteSelection(Integer projectId, Integer caseProjectId, Integer suiteId);

	List<TstCase> queryForRunSelection(Integer projectId, Integer caseProjectId, Integer runId);

	TstCase getById(Integer caseId);

    TstCase renamePers(JSONObject json, TstUser user);
	TstCase delete(Integer vo, TstUser user);

	TstCase renamePers(Integer id, String name, Integer pId, Integer projectId, TstUser user);

	TstCase movePers(JSONObject json, TstUser user);

	void createRoot(Integer projectId, TstUser user);

	void loadNodeTree(TstCase vo, TstCase po);

	TstCase save(JSONObject json, TstUser user);

	void updateParentIfNeededPers(Integer pid);

	boolean cloneStepsAndChildrenPers(TstCase testcase, TstCase src);

	void saveHistory(TstUser user, Constant.CaseAct act, TstCase testCase, String field);

	TstCase saveField(JSONObject json, TstUser user);

	List<TstCase> getChildren(Integer caseId);

	List<TstCase> genVos(List<TstCase> pos);
    List<TstCase> genVos(List<TstCase> pos, boolean withSteps);

	List<TstCase> genVos(List<TstCase> pos, List<Integer> selectIds, boolean withSteps);

	TstCase genVo(TstCase po);

	TstCase genVo(TstCase po, List<Integer> selectIds, boolean withSteps);

	TstCase genVo(TstCase po, boolean withSteps);

    List<TstCaseHistory> findHistories(Integer testCaseId);

    void copyProperties(TstCase testCasePo, TstCase testCaseVo);

    TstCase changeContentTypePers(Integer id, String contentType);

    TstCase reviewPassPers(Integer id, Boolean pass);
}
