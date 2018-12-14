package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.dao.IssueCommentsDao;
import com.ngtesting.platform.model.TstCase;
import com.ngtesting.platform.model.IsuComments;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.IssueCommentsService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
public class IssueCommentsServiceImpl extends BaseServiceImpl implements IssueCommentsService {
    @Autowired
    IssueCommentsDao issueCommentsDao;

    @Override
    @Transactional
    public IsuComments save(JSONObject json, TstUser user) {
        IsuComments vo = JSON.parseObject(JSON.toJSONString(json), IsuComments.class);

        vo.setIssueId(json.getInteger("modelId"));

        save(vo, user);
        return vo;
    }

    @Override
    @Transactional
    public IsuComments save(IsuComments vo, TstUser user) {
        vo.setUserId(user.getId());
        vo.setUserName(user.getNickname());
        vo.setUserAvatar(user.getAvatar());

        if (vo.getId() == null) {
            issueCommentsDao.save(vo);
        } else {
            issueCommentsDao.update(vo);
        }

        return vo;
    }

    @Override
    @Transactional
    public Boolean delete(Integer id, TstUser user) {
        IsuComments comments = issueCommentsDao.get(id);

        Boolean result = issueCommentsDao.delete(id, user.getId());

//        caseHistoryService.saveHistory(user, Constant.EntityAct.comments_delete, testCase, comments.getContent());
        return result;
    }

}
