package com.ngtesting.platform.service;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.AiTestTask;
import com.ngtesting.platform.vo.*;

import java.util.List;

public interface AiTestTaskService extends BaseService {

	List<AiTestTask> query(Long projectId);

	AiTestTaskVo getById(Long caseId);

    AiTestTask renamePers(JSONObject json, UserVo user);
	AiTestTask delete(Long vo, UserVo user);

	AiTestTask renamePers(Long id, String name, Long pId, Long projectId, UserVo user);

	AiTestTaskVo movePers(JSONObject json, UserVo user);

	void loadNodeTree(AiTestTaskVo vo, AiTestTask po);

	AiTestTask save(JSONObject json, UserVo user);

	void updateParentIfNeededPers(Long pid);

	boolean cloneChildrenPers(AiTestTask testcase, AiTestTask src);

	List<AiTestTask> getChildren(Long caseId);

	List<AiTestTaskVo> genVos(List<AiTestTask> pos);
	AiTestTaskVo genVo(AiTestTask po);

    void copyProperties(AiTestTask testCasePo, AiTestTaskVo testCaseVo);

}
