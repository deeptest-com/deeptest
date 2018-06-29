package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.model.TstCasePriority;
import com.ngtesting.platform.service.inf.CasePriorityService;
import com.ngtesting.platform.utils.BeanUtilEx;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;

@Service
public class CasePriorityServiceImpl extends BaseServiceImpl implements CasePriorityService {
	@Override
	public List<TstCasePriority> list(Integer orgId) {
//        DetachedCriteria dc = DetachedCriteria.forClass(TestCasePriority.class);
//
//        dc.add(Restrictions.eq("orgId", orgId));
//        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//
//        dc.addOrder(Order.asc("displayOrder"));
//        List ls = findAllByCriteria(dc);
//
//		return ls;
		return null;
	}
	@Override
	public List<TstCasePriority> listVos(Integer orgId) {
//        List ls = list(orgId);
//
//        List<TstCasePriority> vos = genVos(ls);
//		return vos;

		return null;
	}

	@Override
	public TstCasePriority save(TstCasePriority vo, Integer orgId) {
//		if (vo == null) {
//			return null;
//		}
//
//		TestCasePriority po;
//		if (vo.getId() != null) {
//			po = (TestCasePriority) get(TestCasePriority.class, vo.getId());
//		} else {
//			po = new TestCasePriority();
//		}
//
//		BeanUtilEx.copyProperties(po, vo);
//
//		po.setOrgId(orgId);
//
//		if (vo.getId() == null) {
//			po.setCode(UUID.randomUUID().toString());
//
//			String hql = "select max(displayOrder) from TestCasePriority pri where pri.orgId=?";
//			Integer maxOrder = (Integer) getByHQL(hql, orgId);
//	        po.setDisplayOrder(maxOrder + 10);
//		}
//
//		saveOrUpdate(po);
//		return po;

		return null;
	}

	@Override
	public boolean delete(Integer id) {
//		TestCasePriority po = (TestCasePriority) get(TestCasePriority.class, id);
//		po.setDeleted(true);
//		saveOrUpdate(po);

		return true;
	}

	@Override
	public boolean setDefaultPers(Integer id, Integer orgId) {
//		List<TestCasePriority> ls = list(orgId);
//		for (TestCasePriority priority : ls) {
//			if (priority.getId().longValue() == id.longValue()) {
//				priority.setIsDefault(true);
//				saveOrUpdate(priority);
//			} else if (priority.getIsDefault() != null && priority.getIsDefault()) {
//				priority.setIsDefault(false);
//				saveOrUpdate(priority);
//			}
//		}

		return true;
	}

	@Override
	public boolean changeOrderPers(Integer id, String act, Integer orgId) {
//		TestCasePriority type = (TestCasePriority) get(TestCasePriority.class, id);
//
//        String hql = "from TestCasePriority tp where where tp.orgId=? and tp.deleted = false and tp.disabled = false ";
//        if ("up".equals(act)) {
//        	hql += "and tp.displayOrder < ? order by displayOrder desc";
//        } else if ("down".equals(act)) {
//        	hql += "and tp.displayOrder > ? order by displayOrder asc";
//        } else {
//        	return false;
//        }
//
//        TestCasePriority neighbor = (TestCasePriority) getDao().findFirstByHQL(hql, orgId, type.getDisplayOrder());
//
//        Integer order = type.getDisplayOrder();
//        type.setDisplayOrder(neighbor.getDisplayOrder());
//        neighbor.setDisplayOrder(order);
//
//        saveOrUpdate(type);
//        saveOrUpdate(neighbor);

		return true;
	}

	@Override
	public TstCasePriority genVo(TstCasePriority po) {
		if (po == null) {
			return null;
		}
		TstCasePriority vo = new TstCasePriority();
		BeanUtilEx.copyProperties(vo, po);

		return vo;
	}
	@Override
	public List<TstCasePriority> genVos(List<TstCasePriority> pos) {
        List<TstCasePriority> vos = new LinkedList<TstCasePriority>();

        for (TstCasePriority po: pos) {
        	TstCasePriority vo = genVo(po);
        	vos.add(vo);
        }
		return vos;
	}

}
