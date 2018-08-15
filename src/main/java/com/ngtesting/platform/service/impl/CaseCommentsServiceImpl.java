package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.dao.CaseCommentsDao;
import com.ngtesting.platform.dao.CaseDao;
import com.ngtesting.platform.dao.UserDao;
import com.ngtesting.platform.model.TstCase;
import com.ngtesting.platform.model.TstCaseComments;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.CaseCommentsService;
import com.ngtesting.platform.service.CaseHistoryService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
public class CaseCommentsServiceImpl extends BaseServiceImpl implements CaseCommentsService {
    @Autowired
    CaseCommentsDao caseCommentsDao;
    @Autowired
    CaseHistoryService caseHistoryService;
    @Autowired
    UserDao userDao;
    @Autowired
    CaseDao caseDao;

    @Override
    @Transactional
    public TstCaseComments save(JSONObject json, TstUser user) {
        TstCaseComments vo = JSON.parseObject(JSON.toJSONString(json), TstCaseComments.class);

        vo.setUserId(user.getId());
        vo.setUserName(user.getNickname());
        vo.setUserAvatar(user.getAvatar());

        if (vo.getId() == null) {
            caseCommentsDao.save(vo);
//            caseHistoryService.saveHistory(user, Constant.CaseAct.comments_add, testCase, vo.getContent());
        } else {
            TstCase testCase = caseDao.get(vo.getCaseId(), user.getDefaultPrjId());
            if (testCase == null) {
                return null;
            }

            caseCommentsDao.update(vo);
//            caseHistoryService.saveHistory(user, Constant.CaseAct.comments_update, testCase, vo.getContent());
        }

        return vo;
    }

    @Override
    @Transactional
    public Boolean delete(Integer id, TstUser user) {
        TstCaseComments comments = caseCommentsDao.get(id);

        TstCase testCase = caseDao.get(comments.getCaseId(), user.getDefaultPrjId());
        if (testCase == null) {
            return false;
        }

        Boolean result = caseCommentsDao.delete(id, user.getId());

//        caseHistoryService.saveHistory(user, Constant.CaseAct.comments_delete, testCase, comments.getContent());
        return result;
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
