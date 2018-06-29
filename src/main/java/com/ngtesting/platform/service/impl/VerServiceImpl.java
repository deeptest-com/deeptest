package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.model.TstVer;
import com.ngtesting.platform.service.VerService;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class VerServiceImpl extends BaseServiceImpl implements VerService {
    @Override
    public List<TstVer> list(Long projectId, String keywords, String disabled) {
//        DetachedCriteria dc = DetachedCriteria.forClass(TstVer.class);
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
//        List<TstVer> ls = findAllByCriteria(dc);
//
//        return ls;

        return null;
    }

    @Override
    public TstVer getById(Long caseId) {
//        TstVer po = (TstVer) get(TstVer.class, caseId);
//        TstVer vo = genVo(po);
//
//        return vo;

        return null;
    }

    @Override
    public TstVer save(JSONObject json, TstUser optUser) {
        Long id = json.getLong("id");

        TstVer po = null;
//        TstVer vo = JSON.parseObject(JSON.toJSONString(json), TstVer.class);
//
//        Constant.MsgType action;
//        if (id != null) {
//            po = (TstVer)get(TstVer.class, id);
//            action = Constant.MsgType.update;
//        } else {
//            po = new TstVer();
//            String hql = "select max(displayOrder) from TstVer tp where tp.projectId=? and tp.deleted != true";
//            Integer maxOrder = (Integer) getByHQL(hql, vo.getProjectId());
//            if (maxOrder == null) {
//                maxOrder = 0;
//            }
//            po.setDisplayOrder(maxOrder + 10);
//
//            action = Constant.MsgType.create;
//        }
//        po.setName(vo.getName());
//        po.setStartTime(vo.getStartTime());
//        po.setEndTime(vo.getEndTime());
//        po.setDescr(vo.getDescr());
//        po.setProjectId(vo.getProjectId());
//
//        saveOrUpdate(po);

        return po;
    }

    @Override
    public TstVer delete(Long id, Long clientId) {
//        TstVer po = (TstVer)get(TstVer.class, id);
//        po.setDeleted(true);
//        saveOrUpdate(po);
//        return po;

        return null;
    }

    @Override
    public boolean changeOrderPers(Long id, String act, Long projectId) {
//        TstVer ver = (TstVer) get(TstVer.class, id);
//
//        String hql = "from TstVer tp where tp.projectId=? and tp.deleted = false and tp.disabled = false ";
//        if ("up".equals(act)) {
//            hql += "and tp.displayOrder < ? order by displayOrder desc";
//        } else if ("down".equals(act)) {
//            hql += "and tp.displayOrder > ? order by displayOrder asc";
//        } else {
//            return false;
//        }
//
//        TstVer neighbor = (TstVer) getFirstByHql(hql, projectId, ver.getDisplayOrder());
//
//        Integer order = ver.getDisplayOrder();
//        ver.setDisplayOrder(neighbor.getDisplayOrder());
//        neighbor.setDisplayOrder(order);
//
//        saveOrUpdate(ver);
//        saveOrUpdate(neighbor);

        return true;
    }

}

