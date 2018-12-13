package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.model.TstCaseComments;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.IssueCommentsService;
import org.springframework.stereotype.Service;

@Service
public class IssueCommentsServiceImpl extends BaseServiceImpl implements IssueCommentsService {

    @Override
    public TstCaseComments save(JSONObject json, TstUser TstUser) {
//        TstCaseComments vo = JSON.parseObject(JSON.toJSONString(json), TstCaseComments.class);
//
//        TestCaseComments po = new TestCaseComments();
//
//        if (vo.getCode() != null) {
//            po = (TestCaseComments)getDetail(TestCaseComments.class, vo.getCode());
//        } else {
//            po.setCode(null);
//        }
//        po.setSummary(vo.getSummary());
//        po.setContent(vo.getContent());
//        po.setTestCaseId(vo.getTestCaseId());
//        po.setUserId(TstUser.getCode());
//        po.setChangeTime(new Date());
//        saveOrUpdate(po);
//
//        return genVo(po);

        return null;
    }

    @Override
    public boolean delete(Long id, Long userId) {
//        TestCaseComments po = (TestCaseComments) getDetail(TestCaseComments.class, id);
//        po.setDeleted(true);
//        saveOrUpdate(po);

        return true;
    }

}
