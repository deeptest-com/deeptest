package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.entity.TestCustomFieldOption;
import com.ngtesting.platform.service.CustomFieldOptionService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.vo.CustomFieldOptionVo;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;

@Service
public class CustomFieldOptionServiceImpl extends BaseServiceImpl implements CustomFieldOptionService {

    @Override
    public List<CustomFieldOptionVo> listVos(Long fieldId) {
        DetachedCriteria dc = DetachedCriteria.forClass(TestCustomFieldOption.class);

        dc.add(Restrictions.eq("fieldId", fieldId));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));

        dc.addOrder(Order.asc("ordr"));
        List ls = findAllByCriteria(dc);

        List<CustomFieldOptionVo> vos = genVos(ls);
        return vos;
    }

    @Override
    public TestCustomFieldOption save(CustomFieldOptionVo vo) {
        if (vo == null) {
            return null;
        }

        TestCustomFieldOption po;
        if (vo.getId() != null) {
            po = (TestCustomFieldOption) get(TestCustomFieldOption.class, vo.getId());
        } else {
            po = new TestCustomFieldOption();
        }
        BeanUtilEx.copyProperties(po, vo);

        if (vo.getId() == null) {
            String hql = "select max(ordr) from TestCustomFieldOption opt where opt.fieldId = ?";
            Integer maxOrder = (Integer) getByHQL(hql, vo.getFieldId());
            po.setOrdr(maxOrder + 10);
        }

        saveOrUpdate(po);
        return po;
    }

    @Override
    public boolean delete(Long id) {
        getDao().delete(id);
        return true;
    }

    @Override
    public boolean changeOrderPers(Long id, String act, Long fieldId) {
        TestCustomFieldOption opt = (TestCustomFieldOption) get(TestCustomFieldOption.class, id);

        String hql = "from TestCustomFieldOption opt where opt.fieldId=? and opt.deleted = false and opt.disabled = false ";
        if ("up".equals(act)) {
            hql += "and opt.ordr < ? order by ordr desc";
        } else if ("down".equals(act)) {
            hql += "and opt.ordr > ? order by ordr asc";
        } else {
            return false;
        }

        TestCustomFieldOption neighbor = (TestCustomFieldOption) getDao().findFirstByHQL(hql, fieldId, opt.getOrdr());

        Integer order = opt.getOrdr();
        opt.setOrdr(neighbor.getOrdr());
        neighbor.setOrdr(order);

        saveOrUpdate(opt);
        saveOrUpdate(neighbor);

        return true;
    }

    @Override
    public List<CustomFieldOptionVo> genVos(List<TestCustomFieldOption> pos) {
        List<CustomFieldOptionVo> vos = new LinkedList<>();

        for (TestCustomFieldOption po : pos) {
            CustomFieldOptionVo vo = genVo(po);
            vos.add(vo);
        }
        return vos;
    }

    @Override
    public CustomFieldOptionVo genVo(TestCustomFieldOption po) {
        CustomFieldOptionVo vo = new CustomFieldOptionVo();
        BeanUtilEx.copyProperties(vo, po);

        return vo;
    }
}
