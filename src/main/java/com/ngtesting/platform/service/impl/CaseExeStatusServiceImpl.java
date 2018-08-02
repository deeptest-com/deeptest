package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.model.TstCaseExeStatus;
import com.ngtesting.platform.service.CaseExeStatusService;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class CaseExeStatusServiceImpl extends BaseServiceImpl implements CaseExeStatusService {
	@Override
	public List<TstCaseExeStatus> list(Integer orgId) {
//        DetachedCriteria dc = DetachedCriteria.forClass(TestCaseExeStatus.class);
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
	public TstCaseExeStatus save(TstCaseExeStatus vo, Integer orgId) {
//		if (vo == null) {
//			return null;
//		}
//
//		TestCaseExeStatus po;
//		if (vo.getId() != null) {
//			po = (TestCaseExeStatus) getDetail(TestCaseExeStatus.class, vo.getId());
//		} else {
//			po = new TestCaseExeStatus();
//		}
//
//		BeanUtilEx.copyProperties(po, vo);
//
//		po.setOrgId(orgId);
//
//		if (vo.getId() == null) {
//			po.setCode(UUID.randomUUID().toString());
//
//			String hql = "select max(displayOrder) from TestCaseExeStatus";
//			Integer maxOrder = (Integer) getByHQL(hql);
//	        po.setDisplayOrder(maxOrder + 10);
//		}
//
//		saveOrUpdate(po);
//		return po;

		return null;
	}

	@Override
	public boolean delete(Integer id) {
//		TestCaseExeStatus po = (TestCaseExeStatus) getDetail(TestCaseExeStatus.class, id);
//		po.setDeleted(true);
//		saveOrUpdate(po);

		return true;
	}

	@Override
	public boolean changeOrderPers(Integer id, String act) {
//		TestCaseExeStatus type = (TestCaseExeStatus) getDetail(TestCaseExeStatus.class, id);
//
//        String hql = "from TestCaseExeStatus tp where tp.deleted = false and tp.disabled = false ";
//        if ("up".equals(act)) {
//        	hql += "and tp.displayOrder < ? order by displayOrder desc";
//        } else if ("down".equals(act)) {
//        	hql += "and tp.displayOrder > ? order by displayOrder asc";
//        } else {
//        	return false;
//        }
//
//        TestCaseExeStatus neighbor = (TestCaseExeStatus) getDao().findFirstByHQL(hql, type.getDisplayOrder());
//
//        Integer order = type.getDisplayOrder();
//        type.setDisplayOrder(neighbor.getDisplayOrder());
//        neighbor.setDisplayOrder(order);
//
//        saveOrUpdate(type);
//        saveOrUpdate(neighbor);

		return true;
	}

//	@Override
//	public void createDefaultBasicDataPers(Long orgId) {
//		DetachedCriteria dc = DetachedCriteria.forClass(TestCaseExeStatus.class);
//		dc.add(Restrictions.eq("isBuildIn", true));
//		dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//		dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//
//		dc.addOrder(Order.asc("displayOrder"));
//		List<TestCaseExeStatus> ls = findAllByCriteria(dc);
//
//		for (TestCaseExeStatus p : ls) {
//			TestCaseExeStatus temp = new TestCaseExeStatus();
//			BeanUtilEx.copyProperties(temp, p);
//			temp.setId(null);
//			temp.setOrgId(orgId);
//			temp.setBuildIn(false);
//			saveOrUpdate(temp);
//		}
//	}

}
