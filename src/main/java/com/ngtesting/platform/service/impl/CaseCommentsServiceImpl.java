package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.dao.CaseCommentsDao;
import com.ngtesting.platform.dao.UserDao;
import com.ngtesting.platform.model.TstCaseComments;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.CaseCommentsService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class CaseCommentsServiceImpl extends BaseServiceImpl implements CaseCommentsService {
    @Autowired
    CaseCommentsDao caseCommentsDao;
    @Autowired
    UserDao userDao;

    @Override
    public TstCaseComments save(JSONObject json, TstUser user) {
        TstCaseComments vo = JSON.parseObject(JSON.toJSONString(json), TstCaseComments.class);
        vo.setUserId(user.getId());

        vo.setUserName(user.getNickname());
        vo.setUserAvatar(user.getAvatar());

        if (vo.getId() != null) {
            caseCommentsDao.update(vo);
        } else {
            caseCommentsDao.save(vo);
        }

        return vo;
    }

    @Override
    public boolean delete(Integer id, Integer userId) {
        return caseCommentsDao.delete(id, userId);
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
