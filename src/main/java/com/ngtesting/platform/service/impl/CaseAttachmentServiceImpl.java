package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.entity.TestCase;
import com.ngtesting.platform.entity.TestCaseAttachment;
import com.ngtesting.platform.service.CaseAttachmentService;
import com.ngtesting.platform.service.CaseService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.vo.TestCaseAttachmentVo;
import com.ngtesting.platform.vo.UserVo;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;

@Service
public class CaseAttachmentServiceImpl extends BaseServiceImpl implements CaseAttachmentService {
    @Autowired
    CaseService caseService;

    @Override
    public void uploadAttachmentPers(Long caseId, String name, String path, UserVo user) {
        TestCaseAttachment attach = new TestCaseAttachment();
        attach.setName(name);
        attach.setUri(path);
        attach.setTestCaseId(caseId);
        attach.setUserId(user.getId());

        saveOrUpdate(attach);

        TestCase testCase = (TestCase) get(TestCase.class, caseId);
        caseService.saveHistory(user, Constant.CaseAct.upload_attachment, testCase, name);
    }

    @Override
    public void removeAttachmentPers(Long id, UserVo user) {
        TestCaseAttachment attach = (TestCaseAttachment) get(TestCaseAttachment.class, id);
        attach.setDeleted(true);
        saveOrUpdate(attach);

        caseService.saveHistory(user, Constant.CaseAct.delete_attachment, attach.getTestCase(),attach.getName());
    }

    @Override
    public List<TestCaseAttachmentVo> listByCase(Long caseId) {
        DetachedCriteria dc = DetachedCriteria.forClass(TestCaseAttachment.class);

        dc.add(Restrictions.eq("testCaseId", caseId));

        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));

        dc.addOrder(Order.asc("createTime"));

        List<TestCaseAttachment> ls = findAllByCriteria(dc);

        return genVos(ls);
    }

    @Override
    public List<TestCaseAttachmentVo> genVos(List<TestCaseAttachment> pos) {
        List<TestCaseAttachmentVo> vos = new LinkedList<>();

        for (TestCaseAttachment po: pos) {
            TestCaseAttachmentVo vo = genVo(po);
            vos.add(vo);
        }
        return vos;
    }

    @Override
    public TestCaseAttachmentVo genVo(TestCaseAttachment po) {
        TestCaseAttachmentVo vo = new TestCaseAttachmentVo();

        BeanUtilEx.copyProperties(vo, po);

        return vo;
    }
}

