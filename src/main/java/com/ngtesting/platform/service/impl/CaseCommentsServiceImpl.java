package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.dao.UserDao;
import com.ngtesting.platform.model.TstCaseComments;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.CaseCommentsService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class CaseCommentsServiceImpl extends BaseServiceImpl implements CaseCommentsService {
    @Autowired
    UserDao userDao;

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
    public boolean delete(Integer id, Integer userId) {
//        TestCaseComments po = (TestCaseComments) get(TestCaseComments.class, id);
//        po.setDeleted(true);
//        saveOrUpdate(po);

        return true;
    }

    @Override
    public TstCaseComments genVo(TstCaseComments po) {
        if (po.getUpdateTime() == null) {
            po.setUpdateTime(po.getCreateTime());
        }

        TstUser user = userDao.get(po.getUserId());

        po.setUserName(user.getNickname());
        po.setUserAvatar(user.getAvatar());
        return po;
    }

}
