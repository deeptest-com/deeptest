package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.model.AiTestTask;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.inf.AiTestTaskService;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class AiTestTypeServiceImpl extends BaseServiceImpl implements AiTestTaskService {

    @Override
    public List<AiTestTask> query(Long projectId) {
        return null;
    }

    @Override
    public AiTestTask getById(Long caseId) {
        return null;
    }

    @Override
    public AiTestTask renamePers(JSONObject json, TstUser user) {
        return null;
    }

    @Override
    public AiTestTask delete(Long vo, TstUser user) {
        return null;
    }

    @Override
    public AiTestTask renamePers(Long id, String name, Long pId, Long projectId, TstUser user) {
        return null;
    }

    @Override
    public AiTestTask movePers(JSONObject json, TstUser user) {
        return null;
    }

    @Override
    public void loadNodeTree(AiTestTask vo, AiTestTask po) {

    }

    @Override
    public AiTestTask save(JSONObject json, TstUser user) {
        return null;
    }

    @Override
    public void updateParentIfNeededPers(Long pid) {

    }

    @Override
    public boolean cloneChildrenPers(AiTestTask testcase, AiTestTask src) {
        return false;
    }

    @Override
    public List<AiTestTask> getChildren(Long caseId) {
        return null;
    }

    @Override
    public void copyProperties(AiTestTask testCasePo, AiTestTask testCaseVo) {

    }
}

