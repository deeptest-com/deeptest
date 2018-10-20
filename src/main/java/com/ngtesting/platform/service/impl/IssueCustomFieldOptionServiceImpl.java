package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.model.TstCustomFieldOption;
import com.ngtesting.platform.service.IssueCustomFieldOptionService;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;

@Service
public class IssueCustomFieldOptionServiceImpl extends BaseServiceImpl implements IssueCustomFieldOptionService {

    @Override
    public List<TstCustomFieldOption> listVos(Integer fieldId) {
//        DetachedCriteria dc = DetachedCriteria.forClass(TestCustomFieldOption.class);
//
//        dc.add(Restrictions.eq("fieldId", fieldId));
//        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//
//        dc.addOrder(Order.asc("ordr"));
//        List ls = findAllByCriteria(dc);
//
//        List<TstCustomFieldOption> vos = genVos(ls);
//        return vos;

        return null;
    }

    @Override
    public TstCustomFieldOption save(TstCustomFieldOption vo) {
//        if (vo == null) {
//            return null;
//        }
//
//        TestCustomFieldOption po;
//        if (vo.getCode() != null) {
//            po = (TestCustomFieldOption) getDetail(TestCustomFieldOption.class, vo.getCode());
//        } else {
//            po = new TestCustomFieldOption();
//        }
//        BeanUtilEx.copyProperties(po, vo);
//
//        if (vo.getCode() == null) {
//            String hql = "select max(ordr) from TestCustomFieldOption opt where opt.fieldId = ?";
//            Object obj = getByHQL(hql, vo.getFieldId());
//            Integer maxOrder = obj!=null?(Integer) getByHQL(hql, vo.getFieldId()): 10;
//            po.setOrdr(maxOrder + 10);
//        }
//
//        saveOrUpdate(po);
//        return po;

        return null;
    }

    @Override
    public boolean delete(Integer id) {
//        getDao().delete(id);
        return true;
    }

    @Override
    public boolean changeOrderPers(Integer id, String act, Integer fieldId) {
//        TestCustomFieldOption opt = (TestCustomFieldOption) getDetail(TestCustomFieldOption.class, id);
//
//        String hql = "from TestCustomFieldOption opt where opt.fieldId=? and opt.deleted = false and opt.disabled = false ";
//        if ("up".equals(act)) {
//            hql += "and opt.ordr < ? order by ordr desc";
//        } else if ("down".equals(act)) {
//            hql += "and opt.ordr > ? order by ordr asc";
//        } else {
//            return false;
//        }
//
//        TestCustomFieldOption neighbor = (TestCustomFieldOption) getDao().findFirstByHQL(hql, fieldId, opt.getOrdr());
//
//        Integer order = opt.getOrdr();
//        opt.setOrdr(neighbor.getOrdr());
//        neighbor.setOrdr(order);
//
//        saveOrUpdate(opt);
//        saveOrUpdate(neighbor);

        return true;
    }

    @Override
    public List<TstCustomFieldOption> genVos(List<TstCustomFieldOption> pos) {
        List<TstCustomFieldOption> vos = new LinkedList<>();

//        for (TestCustomFieldOption po : pos) {
//            TstCustomFieldOption vo = genVo(po);
//            vos.add(vo);
//        }
        return vos;
    }

    @Override
    public TstCustomFieldOption genVo(TstCustomFieldOption po) {
//        TstCustomFieldOption vo = new TstCustomFieldOption();
//        BeanUtilEx.copyProperties(vo, po);
//
//        return vo;

        return null;
    }
}
