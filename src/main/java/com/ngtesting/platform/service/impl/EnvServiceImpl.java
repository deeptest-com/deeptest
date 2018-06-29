package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.model.TstEnv;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.inf.EnvService;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;

@Service
public class EnvServiceImpl extends BaseServiceImpl implements EnvService {

    @Override
    public List<TstEnv> list(Integer projectId, String keywords, String disabled) {
//        DetachedCriteria dc = DetachedCriteria.forClass(TstEnv.class);
//
//        dc.add(Restrictions.eq("projectId", projectId));
//        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//
//        if (StringUtil.isNotEmpty(keywords)) {
//            dc.add(Restrictions.like("name", "%" + keywords + "%"));
//        }
//        if (StringUtil.isNotEmpty(disabled)) {
//            dc.add(Restrictions.eq("disabled", Boolean.valueOf(disabled)));
//        }
//
//        dc.addOrder(Order.asc("displayOrder"));
//
//        List<TstEnv> ls = findAllByCriteria(dc);
//
//        return ls;

        return null;
    }

    @Override
    public TstEnv getById(Integer caseId) {
//        TstEnv po = (TstEnv) get(TstEnv.class, caseId);
//        TstEnv vo = genVo(po);
//
//        return vo;

        return null;
    }

    @Override
    public TstEnv save(JSONObject json, TstUser optUser) {
//        Long id = json.getLong("id");
//
//        TstEnv po;
//        TstEnv vo = JSON.parseObject(JSON.toJSONString(json), TstEnv.class);
//
//        Constant.MsgType action;
//        if (id != null) {
//            po = (TstEnv)get(TstEnv.class, id);
//            action = Constant.MsgType.update;
//        } else {
//            po = new TstEnv();
//            String hql = "select max(displayOrder) from TstEnv tp where tp.projectId=? and tp.deleted != true";
//            Integer maxOrder = (Integer) getByHQL(hql, vo.getProjectId());
//            if (maxOrder == null) {
//                maxOrder = 0;
//            }
//            po.setDisplayOrder(maxOrder + 10);
//
//            action = Constant.MsgType.create;
//        }
//        po.setName(vo.getName());
//        po.setDescr(vo.getDescr());
//        po.setProjectId(vo.getProjectId());
//
//        saveOrUpdate(po);
//
//        return po;

        return null;
    }

    @Override
    public TstEnv delete(Integer id, Integer clientId) {
//        TstEnv po = (TstEnv)get(TstEnv.class, id);
//        po.setDeleted(true);
//        saveOrUpdate(po);
//        return po;

        return null;
    }

    @Override
    public boolean changeOrderPers(Integer id, String act, Integer orgId) {
//        TstEnv ver = (TstEnv) get(TstEnv.class, id);
//
//        String hql = "from TstEnv tp where tp.projectId=? and tp.deleted = false and tp.disabled = false ";
//        if ("up".equals(act)) {
//            hql += "and tp.displayOrder < ? order by displayOrder desc";
//        } else if ("down".equals(act)) {
//            hql += "and tp.displayOrder > ? order by displayOrder asc";
//        } else {
//            return false;
//        }
//
//        TstEnv neighbor = (TstEnv) getFirstByHql(hql, orgId, ver.getDisplayOrder());
//
//        Integer order = ver.getDisplayOrder();
//        ver.setDisplayOrder(neighbor.getDisplayOrder());
//        neighbor.setDisplayOrder(order);
//
//        saveOrUpdate(ver);
//        saveOrUpdate(neighbor);

        return true;
    }
    @Override
    public List<TstEnv> listVos(Integer projectId) {
        List ls = list(projectId, null, null);

        List<TstEnv> vos = genVos(ls);
        return vos;
    }

    @Override
    public List<TstEnv> genVos(List<TstEnv> pos) {
        List<TstEnv> vos = new LinkedList<TstEnv>();

        for (TstEnv po : pos) {
            TstEnv vo = genVo(po);
            vos.add(vo);
        }
        return vos;
    }

    @Override
    public TstEnv genVo(TstEnv po) {
        TstEnv vo = new TstEnv();

        vo.setId(po.getId());
        vo.setName(po.getName());
        vo.setDescr(po.getDescr());
        vo.setProjectId(po.getProjectId());
        vo.setDisabled(po.getDisabled());

        return vo;
    }

}

