package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.dao.CaseInTaskCommentsDao;
import com.ngtesting.platform.dao.CaseInTaskDao;
import com.ngtesting.platform.dao.UserDao;
import com.ngtesting.platform.model.TstCase;
import com.ngtesting.platform.model.TstCaseInTaskComments;
import com.ngtesting.platform.model.TstCaseInTask;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.CaseHistoryService;
import com.ngtesting.platform.service.intf.CaseInTaskCommentsService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
public class CaseInTaskCommentsServiceImpl extends BaseServiceImpl implements CaseInTaskCommentsService {
    @Autowired
    CaseInTaskCommentsDao caseInTaskCommentsDao;
    @Autowired
    CaseHistoryService caseHistoryService;
    @Autowired
    UserDao userDao;
    @Autowired
    CaseInTaskDao caseInTaskDao;

    @Override
    @Transactional
    public TstCaseInTaskComments save(JSONObject json, TstUser user) {
        TstCaseInTaskComments vo = JSON.parseObject(JSON.toJSONString(json), TstCaseInTaskComments.class);

        vo.setCaseInTaskId(json.getInteger("modelId"));

        save(vo, user);
        return vo;
    }

    @Override
    @Transactional
    public TstCaseInTaskComments save(TstCaseInTaskComments vo, TstUser user) {
        vo.setUserId(user.getId());
        vo.setUserName(user.getNickname());
        vo.setUserAvatar(user.getAvatar());

        if (vo.getId() == null) {
            caseInTaskCommentsDao.save(vo);
//            caseHistoryService.saveHistory(user, MsgUtil.MsgAction.comments_add, testCase, vo.getContent());
        } else {
            caseInTaskCommentsDao.update(vo, user.getId());
//            caseHistoryService.saveHistory(user, MsgUtil.MsgAction.comments_update, testCase, vo.getContent());
        }

        return vo;
    }

    @Override
    @Transactional
    public Boolean delete(Integer id, TstUser user) {
        TstCaseInTaskComments comments = caseInTaskCommentsDao.get(id);

        Boolean result = caseInTaskCommentsDao.delete(id, user.getId());

//        caseHistoryService.saveHistory(user, MsgUtil.MsgAction.comments_delete, testCase, comments.getContent());
        return result;
    }

}
