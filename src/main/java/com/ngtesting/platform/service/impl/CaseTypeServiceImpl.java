package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.entity.TestCaseType;
import com.ngtesting.platform.service.CaseTypeService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.vo.CaseTypeVo;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;
import java.util.UUID;

@Service
public class CaseTypeServiceImpl extends BaseServiceImpl implements CaseTypeService {

	@Override
	public List<TestCaseType> list(Long orgId) {
        DetachedCriteria dc = DetachedCriteria.forClass(TestCaseType.class);
        
        dc.add(Restrictions.eq("orgId", orgId));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        
        dc.addOrder(Order.asc("displayOrder"));
        List ls = findAllByCriteria(dc);
		
		return ls;
	}
	@Override
	public List<CaseTypeVo> listVos(Long orgId) {
        List ls = list(orgId);
        
        List<CaseTypeVo> vos = genVos(ls);
		return vos;
	}

	@Override
	public TestCaseType save(CaseTypeVo vo, Long orgId) {
		if (vo == null) {
			return null;
		}
		
		TestCaseType po;
		if (vo.getId() != null) {
			po = (TestCaseType) get(TestCaseType.class, vo.getId());
		} else {
			po = new TestCaseType();
		}
		
		BeanUtilEx.copyProperties(po, vo);
		po.setOrgId(orgId);
		
		if (vo.getId() == null) {
			po.setCode(UUID.randomUUID().toString());
			
			String hql = "select max(displayOrder) from TestCaseType";
			Integer maxOrder = (Integer) getByHQL(hql);
	        po.setDisplayOrder(maxOrder + 10);
		}
		
		saveOrUpdate(po);
		return po;
	}

	@Override
	public boolean delete(Long id) {
		TestCaseType po = (TestCaseType) get(TestCaseType.class, id);
		po.setDeleted(true);
		saveOrUpdate(po);
		
		return true;
	}

	@Override
	public boolean setDefaultPers(Long id, Long orgId) {
		List<TestCaseType> ls = list(orgId);
		for (TestCaseType type : ls) {
			if (type.getId().longValue() == id.longValue()) {
				type.setIsDefault(true);
				saveOrUpdate(type);
			} else if (type.getIsDefault()) {
				type.setIsDefault(false);
				saveOrUpdate(type);
			}
		}
		
		return true;
	}
	
	@Override
	public boolean changeOrderPers(Long id, String act) {
		TestCaseType type = (TestCaseType) get(TestCaseType.class, id);
		
        String hql = "from TestCaseType tp where tp.deleted = false and tp.disabled = false ";
        if ("up".equals(act)) {
        	hql += "and tp.displayOrder < ? order by displayOrder desc";
        } else if ("down".equals(act)) {
        	hql += "and tp.displayOrder > ? order by displayOrder asc";
        } else {
        	return false;
        }
        
        TestCaseType neighbor = (TestCaseType) getFirstByHql(hql, type.getDisplayOrder());
		
        Integer order = type.getDisplayOrder();
        type.setDisplayOrder(neighbor.getDisplayOrder());
        neighbor.setDisplayOrder(order);
        
        saveOrUpdate(type);
        saveOrUpdate(neighbor);
		
		return true;
	}

//	@Override
//	public void createDefaultBasicDataPers(Long orgId) {
//		DetachedCriteria dc = DetachedCriteria.forClass(TestCaseType.class);
//		dc.add(Restrictions.eq("isBuildIn", true));
//		dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//		dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//
//		dc.addOrder(Order.asc("displayOrder"));
//		List<TestCaseType> ls = findAllByCriteria(dc);
//
//		for (TestCaseType p : ls) {
//			TestCaseType temp = new TestCaseType();
//			BeanUtilEx.copyProperties(temp, p);
//			temp.setId(null);
//			temp.setOrgId(orgId);
//			saveOrUpdate(temp);
//		}
//	}

	@Override
	public CaseTypeVo genVo(TestCaseType po) {
		if (po == null) {
			return null;
		}
		CaseTypeVo vo = new CaseTypeVo();
		BeanUtilEx.copyProperties(vo, po);
		
		return vo;
	}
	@Override
	public List<CaseTypeVo> genVos(List<TestCaseType> pos) {
        List<CaseTypeVo> vos = new LinkedList<CaseTypeVo>();

        for (TestCaseType po: pos) {
        	CaseTypeVo vo = genVo(po);
        	vos.add(vo);
        }
		return vos;
	}

}
