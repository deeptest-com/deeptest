package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.model.TstCaseComments;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.CaseCommentsService;
import org.springframework.stereotype.Service;

@Service
public class CaseCommentsServiceImpl extends BaseServiceImpl implements CaseCommentsService {

    @Override
    public TstCaseComments save(JSONObject json, TstUser TstUser) {
        TstCaseComments vo = JSON.parseObject(JSON.toJSONString(json), TstCaseComments.class);

//        TestCaseComments po = new TestCaseComments();
//
//        if (vo.getId() != null) {
//            po = (TestCaseComments)get(TestCaseComments.class, vo.getId());
//        } else {
//            po.setId(null);
//        }
//        po.setSummary(vo.getSummary());
//        po.setContent(vo.getContent());
//        po.setTestCaseId(vo.getTestCaseId());
//        po.setUserId(TstUser.getId());
//        po.setChangeTime(new Date());
//        saveOrUpdate(po);

//        return genVo(po);

        return null;
    }

    @Override
    public boolean delete(Long id, Long userId) {
//        TestCaseComments po = (TestCaseComments) get(TestCaseComments.class, id);
//        po.setDeleted(true);
//        saveOrUpdate(po);

        return true;
    }

    @Override
    public TstCaseComments genVo(TstCaseComments po) {
        TstCaseComments vo = new TstCaseComments();
//        BeanUtilEx.copyProperties(vo, po);
//        if (vo.getUpdateTime() == null) {
//            vo.setUpdateTime(vo.getCreateTime());
//        }
//
//        TestUser user = (TestUser)get(TestUser.class, po.getUserId());
//
//        vo.setUserName(user.getName());
//        vo.setUserAvatar(user.getAvatar());
        return vo;
    }

}
