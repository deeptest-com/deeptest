package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.model.TstCaseType;
import com.ngtesting.platform.service.CaseTypeService;
import com.ngtesting.platform.utils.BeanUtilEx;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;

@Service
public class CaseTypeServiceImpl extends BaseServiceImpl implements CaseTypeService {

	@Override
	public List<TstCaseType> list(Integer orgId) {
//        DetachedCriteria dc = DetachedCriteria.forClass(TstCaseType.class);
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
	public List<TstCaseType> listVos(Integer orgId) {
//        List ls = list(orgId);
//
//        List<TstCaseType> vos = genVos(ls);
//		return vos;

		return null;
	}

	@Override
	public TstCaseType save(TstCaseType vo, Integer orgId) {
//		if (vo == null) {
//			return null;
//		}
//
//		TstCaseType po;
//		if (vo.getId() != null) {
//			po = (TstCaseType) get(TstCaseType.class, vo.getId());
//		} else {
//			po = new TstCaseType();
//		}
//
//		BeanUtilEx.copyProperties(po, vo);
//		po.setOrgId(orgId);
//
//		if (vo.getId() == null) {
//			po.setCode(UUID.randomUUID().toString());
//
//			String hql = "select max(displayOrder) from TstCaseType tp where tp.orgId=?";
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
//		TstCaseType po = (TstCaseType) get(TstCaseType.class, id);
//		po.setDeleted(true);
//		saveOrUpdate(po);

		return true;
	}

	@Override
	public boolean setDefaultPers(Integer id, Integer orgId) {
//		List<TstCaseType> ls = list(orgId);
//		for (TstCaseType type : ls) {
//			if (type.getId().longValue() == id.longValue()) {
//				type.setIsDefault(true);
//				saveOrUpdate(type);
//			} else if (type.getIsDefault() != null && type.getIsDefault()) {
//				type.setIsDefault(false);
//				saveOrUpdate(type);
//			}
//		}

		return true;
	}

	@Override
	public boolean changeOrderPers(Integer id, String act, Integer orgId) {
//		TstCaseType type = (TstCaseType) get(TstCaseType.class, id);
//
//        String hql = "from TstCaseType tp where tp.orgId=? and tp.deleted = false and tp.disabled = false ";
//        if ("up".equals(act)) {
//        	hql += "and tp.displayOrder < ? order by displayOrder desc";
//        } else if ("down".equals(act)) {
//        	hql += "and tp.displayOrder > ? order by displayOrder asc";
//        } else {
//        	return false;
//        }
//
//        TstCaseType neighbor = (TstCaseType) getFirstByHql(hql, orgId, type.getDisplayOrder());
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
	public TstCaseType genVo(TstCaseType po) {
		if (po == null) {
			return null;
		}
		TstCaseType vo = new TstCaseType();
		BeanUtilEx.copyProperties(vo, po);

		return vo;
	}
	@Override
	public List<TstCaseType> genVos(List<TstCaseType> pos) {
        List<TstCaseType> vos = new LinkedList<TstCaseType>();

        for (TstCaseType po: pos) {
        	TstCaseType vo = genVo(po);
        	vos.add(vo);
        }
		return vos;
	}

}
