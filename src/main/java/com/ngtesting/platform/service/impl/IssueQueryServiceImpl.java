package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.model.TstCasePriority;
import com.ngtesting.platform.service.IssueQueryService;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class IssueQueryServiceImpl extends BaseServiceImpl implements IssueQueryService {
	@Override
	public List<TstCasePriority> list(Integer orgId) {
//        DetachedCriteria dc = DetachedCriteria.forClass(TstCasePriority.class);
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
//		TstCasePriority po;
//		if (vo.getCode() != null) {
//			po = (TstCasePriority) getDetail(TstCasePriority.class, vo.getCode());
//		} else {
//			po = new TstCasePriority();
//		}
//
//		BeanUtilEx.copyProperties(po, vo);
//
//		po.setOrgId(orgId);
//
//		if (vo.getCode() == null) {
//			po.setCode(UUID.randomUUID().toString());
//
//			String hql = "select max(displayOrder) from TstCasePriority pri where pri.orgId=?";
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
//		TstCasePriority po = (TstCasePriority) getDetail(TstCasePriority.class, id);
//		po.setDeleted(true);
//		saveOrUpdate(po);

		return true;
	}

	@Override
	public boolean setDefaultPers(Integer id, Integer orgId) {
//		List<TstCasePriority> ls = list(orgId);
//		for (TstCasePriority priority : ls) {
//			if (priority.getCode().longValue() == id.longValue()) {
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
//		TstCasePriority type = (TstCasePriority) getDetail(TstCasePriority.class, id);
//
//        String hql = "from TstCasePriority tp where where tp.orgId=? and tp.deleted = false and tp.disabled = false ";
//        if ("up".equals(act)) {
//        	hql += "and tp.displayOrder < ? order by displayOrder desc";
//        } else if ("down".equals(act)) {
//        	hql += "and tp.displayOrder > ? order by displayOrder asc";
//        } else {
//        	return false;
//        }
//
//        TstCasePriority neighbor = (TstCasePriority) getDao().findFirstByHQL(hql, orgId, type.getDisplayOrder());
//
//        Integer order = type.getDisplayOrder();
//        type.setDisplayOrder(neighbor.getDisplayOrder());
//        neighbor.setDisplayOrder(order);
//
//        saveOrUpdate(type);
//        saveOrUpdate(neighbor);

		return true;
	}

}
