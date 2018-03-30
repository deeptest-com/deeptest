package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.entity.TestVer;
import com.ngtesting.platform.service.VerService;
import com.ngtesting.platform.vo.TestVerVo;
import com.ngtesting.platform.vo.UserVo;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;

@Service
public class VerServiceImpl extends BaseServiceImpl implements VerService {
    @Override
    public List<TestVer> list(Long projectId) {
        DetachedCriteria dc = DetachedCriteria.forClass(TestVer.class);

        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));

        dc.addOrder(Order.asc("createTime"));

        List<TestVer> ls = findAllByCriteria(dc);

        return ls;
    }

    @Override
    public TestVerVo getById(Long caseId) {
        TestVer po = (TestVer) get(TestVer.class, caseId);
        TestVerVo vo = genVo(po);

        return vo;
    }

    @Override
    public List<TestVerVo> genVos(List<TestVer> pos) {
        List<TestVerVo> vos = new LinkedList<TestVerVo>();

        for (TestVer po : pos) {
            TestVerVo vo = genVo(po);
            vos.add(vo);
        }
        return vos;
    }

    @Override
    public TestVer save(JSONObject json, UserVo optUser) {
        Long id = json.getLong("id");

        TestVer po;
        TestVerVo vo = JSON.parseObject(JSON.toJSONString(json), TestVerVo.class);

        Constant.MsgType action;
        if (id != null) {
            po = (TestVer)get(TestVer.class, id);
            action = Constant.MsgType.update;
        } else {
            po = new TestVer();
            action = Constant.MsgType.create;
        }
        po.setName(vo.getName());
        po.setStartTime(vo.getStartTime());
        po.setEndTime(vo.getEndTime());
        po.setDescr(vo.getDescr());
        po.setProjectId(vo.getProjectId());

        saveOrUpdate(po);

        return po;
    }

    @Override
    public TestVer delete(Long id, Long clientId) {
        TestVer po = (TestVer)get(TestVer.class, id);
        po.setDeleted(true);
        saveOrUpdate(po);
        return po;
    }

    @Override
    public TestVerVo genVo(TestVer po) {
        TestVerVo vo = new TestVerVo();

        vo.setId(po.getId());
        vo.setName(po.getName());
        vo.setStartTime(po.getStartTime());
        vo.setEndTime(po.getEndTime());
        vo.setDescr(po.getDescr());
        vo.setProjectId(po.getProjectId());
        vo.setStatus(po.getStatus().toString());

        return vo;
    }

}

