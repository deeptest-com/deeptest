package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.TestCaseComments;
import com.ngtesting.platform.entity.TestUser;
import com.ngtesting.platform.service.CaseCommentsService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.vo.TestCaseCommentsVo;
import com.ngtesting.platform.vo.UserVo;
import org.springframework.stereotype.Service;

import java.util.Date;

@Service
public class CaseCommentsServiceImpl extends BaseServiceImpl implements CaseCommentsService {

    @Override
    public TestCaseCommentsVo save(JSONObject json, UserVo userVo) {
        TestCaseCommentsVo vo = JSON.parseObject(JSON.toJSONString(json), TestCaseCommentsVo.class);

        TestCaseComments po = new TestCaseComments();

        if (vo.getId() != null) {
            po = (TestCaseComments)get(TestCaseComments.class, vo.getId());
        } else {
            po.setId(null);
        }
        po.setSummary(vo.getSummary());
        po.setContent(vo.getContent());
        po.setTestCaseId(vo.getTestCaseId());
        po.setUpdateById(userVo.getId());
        po.setUpdateTime(new Date());
        saveOrUpdate(po);

        return genVo(po);
    }

    @Override
    public boolean delete(Long id, Long userId) {
        TestCaseComments po = (TestCaseComments) get(TestCaseComments.class, id);
        po.setDeleted(true);
        saveOrUpdate(po);

        return true;
    }

    @Override
    public TestCaseCommentsVo genVo(TestCaseComments po) {
        TestCaseCommentsVo vo = new TestCaseCommentsVo();
        BeanUtilEx.copyProperties(vo, po);

        TestUser user = po.getUpdateBy() != null?po.getUpdateBy(): (TestUser)get(TestUser.class, po.getId());
        vo.setUserName(user.getName());
        vo.setUserAvatar(user.getAvatar());
        return vo;
    }

}
