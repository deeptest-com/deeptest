package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.entity.TestEnv;
import com.ngtesting.platform.service.EnvService;
import com.ngtesting.platform.vo.TestEnvVo;
import com.ngtesting.platform.vo.UserVo;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;

@Service
public class EnvServiceImpl extends BaseServiceImpl implements EnvService {

    @Override
    public List<TestEnv> list(Long projectId) {
        DetachedCriteria dc = DetachedCriteria.forClass(TestEnv.class);

        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));

        dc.addOrder(Order.asc("createTime"));

        List<TestEnv> ls = findAllByCriteria(dc);

        return ls;
    }

    @Override
    public TestEnvVo getById(Long caseId) {
        TestEnv po = (TestEnv) get(TestEnv.class, caseId);
        TestEnvVo vo = genVo(po);

        return vo;
    }

    @Override
    public TestEnv save(JSONObject json, UserVo optUser) {
        Long id = json.getLong("id");

        TestEnv po;
        TestEnvVo vo = JSON.parseObject(JSON.toJSONString(json), TestEnvVo.class);

        Constant.MsgType action;
        if (id != null) {
            po = (TestEnv)get(TestEnv.class, id);
            action = Constant.MsgType.update;
        } else {
            po = new TestEnv();
            action = Constant.MsgType.create;
        }
        po.setName(vo.getName());
        po.setDescr(vo.getDescr());
        po.setProjectId(vo.getProjectId());

        saveOrUpdate(po);

        return po;
    }

    @Override
    public TestEnv delete(Long id, Long clientId) {
        TestEnv po = (TestEnv)get(TestEnv.class, id);
        po.setDeleted(true);
        saveOrUpdate(po);
        return po;
    }

    @Override
    public List<TestEnvVo> genVos(List<TestEnv> pos) {
        List<TestEnvVo> vos = new LinkedList<TestEnvVo>();

        for (TestEnv po : pos) {
            TestEnvVo vo = genVo(po);
            vos.add(vo);
        }
        return vos;
    }

    @Override
    public TestEnvVo genVo(TestEnv po) {
        TestEnvVo vo = new TestEnvVo();

        vo.setId(po.getId());
        vo.setName(po.getName());
        vo.setDescr(po.getDescr());
        vo.setProjectId(po.getProjectId());

        return vo;
    }

}

