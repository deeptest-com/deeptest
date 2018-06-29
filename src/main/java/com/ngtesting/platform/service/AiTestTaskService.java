package com.ngtesting.platform.service;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.model.AiTestTask;
import com.ngtesting.platform.model.TstUser;

import java.util.List;

public interface AiTestTaskService extends BaseService {

	List<AiTestTask> query(Long projectId);

	AiTestTask getById(Long caseId);

	AiTestTask renamePers(JSONObject json, TstUser user);
	AiTestTask delete(Long vo, TstUser user);

	AiTestTask renamePers(Long id, String name, Long pId, Long projectId, TstUser user);

	AiTestTask movePers(JSONObject json, TstUser user);

	void loadNodeTree(AiTestTask vo, AiTestTask po);

	AiTestTask save(JSONObject json, TstUser user);

	void updateParentIfNeededPers(Long pid);

	boolean cloneChildrenPers(AiTestTask testcase, AiTestTask src);

	List<AiTestTask> getChildren(Long caseId);


    void copyProperties(AiTestTask testCasePo, AiTestTask testCaseVo);

}
