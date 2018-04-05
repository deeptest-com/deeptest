package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.entity.TestVer;
import com.ngtesting.platform.service.VerService;
import com.ngtesting.platform.util.StringUtil;
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
    public List<TestVer> list(Long projectId, String keywords, String disabled) {
        DetachedCriteria dc = DetachedCriteria.forClass(TestVer.class);

        dc.add(Restrictions.eq("projectId", projectId));
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));

        if (StringUtil.isNotEmpty(keywords)) {
            dc.add(Restrictions.like("name", "%" + keywords + "%"));
        }
        if (StringUtil.isNotEmpty(disabled)) {
            dc.add(Restrictions.eq("disabled", Boolean.valueOf(disabled)));
        }

        dc.addOrder(Order.asc("displayOrder"));

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
            String hql = "select max(displayOrder) from TestVer tp where tp.projectId=?";
            Integer maxOrder = (Integer) getByHQL(hql, vo.getProjectId());
            po.setDisplayOrder(maxOrder + 10);

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
    public boolean changeOrderPers(Long id, String act, Long projectId) {
        TestVer ver = (TestVer) get(TestVer.class, id);

        String hql = "from TestVer tp where tp.projectId=? and tp.deleted = false and tp.disabled = false ";
        if ("up".equals(act)) {
            hql += "and tp.displayOrder < ? order by displayOrder desc";
        } else if ("down".equals(act)) {
            hql += "and tp.displayOrder > ? order by displayOrder asc";
        } else {
            return false;
        }

        TestVer neighbor = (TestVer) getFirstByHql(hql, projectId, ver.getDisplayOrder());

        Integer order = ver.getDisplayOrder();
        ver.setDisplayOrder(neighbor.getDisplayOrder());
        neighbor.setDisplayOrder(order);

        saveOrUpdate(ver);
        saveOrUpdate(neighbor);

        return true;
    }
    @Override
    public List<TestVerVo> listVos(Long projectId) {
        List ls = list(projectId, null, null);

        List<TestVerVo> vos = genVos(ls);
        return vos;
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
        vo.setDisabled(po.getDisabled());

        return vo;
    }

}

