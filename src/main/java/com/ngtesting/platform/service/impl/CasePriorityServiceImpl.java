package com.ngtesting.platform.service.impl;

import java.util.LinkedList;
import java.util.List;
import java.util.UUID;

import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import com.ngtesting.platform.entity.TestCasePriority;
import com.ngtesting.platform.service.CasePriorityService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.vo.CasePriorityVo;

@Service
public class CasePriorityServiceImpl extends BaseServiceImpl implements CasePriorityService {
	@Override
	public List<TestCasePriority> list(Long orgId) {
        DetachedCriteria dc = DetachedCriteria.forClass(TestCasePriority.class);
        
        dc.add(Restrictions.eq("orgId", orgId));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        
        dc.addOrder(Order.asc("displayOrder"));
        List ls = findAllByCriteria(dc);
		
		return ls;
	}
	@Override
	public List<CasePriorityVo> listVos(Long orgId) {
        List ls = list(orgId);
        
        List<CasePriorityVo> vos = genVos(ls);
		return vos;
	}

	@Override
	public TestCasePriority save(CasePriorityVo vo, Long orgId) {
		if (vo == null) {
			return null;
		}
		
		TestCasePriority po;
		if (vo.getId() != null) {
			po = (TestCasePriority) get(TestCasePriority.class, vo.getId());
		} else {
			po = new TestCasePriority();
		}
		
		BeanUtilEx.copyProperties(po, vo);
		
		po.setOrgId(orgId);
		
		if (vo.getId() == null) {
			po.setCode(UUID.randomUUID().toString());
			
			String hql = "select max(displayOrder) from TestCasePriority";
			Integer maxOrder = (Integer) getByHQL(hql);
	        po.setDisplayOrder(maxOrder + 10);
		}
		
		saveOrUpdate(po);
		return po;
	}

	@Override
	public boolean delete(Long id) {
		TestCasePriority po = (TestCasePriority) get(TestCasePriority.class, id);
		po.setDeleted(true);
		saveOrUpdate(po);
		
		return true;
	}

	@Override
	public boolean setDefaultPers(Long id, Long orgId) {
		List<TestCasePriority> ls = list(orgId);
		for (TestCasePriority type : ls) {
			if (type.getId() == id) {
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
		TestCasePriority type = (TestCasePriority) get(TestCasePriority.class, id);
		
        String hql = "from TestCasePriority tp where tp.deleted = false and tp.disabled = false ";
        if ("up".equals(act)) {
        	hql += "and tp.displayOrder < ? order by displayOrder desc";
        } else if ("down".equals(act)) {
        	hql += "and tp.displayOrder > ? order by displayOrder asc";
        } else {
        	return false;
        }
        
        TestCasePriority neighbor = (TestCasePriority) getDao().findFirstByHQL(hql, type.getDisplayOrder());
		
        Integer order = type.getDisplayOrder();
        type.setDisplayOrder(neighbor.getDisplayOrder());
        neighbor.setDisplayOrder(order);
        
        saveOrUpdate(type);
        saveOrUpdate(neighbor);
		
		return true;
	}
	
    
	@Override
	public CasePriorityVo genVo(TestCasePriority po) {
		if (po == null) {
			return null;
		}
		CasePriorityVo vo = new CasePriorityVo();
		BeanUtilEx.copyProperties(vo, po);
		
		return vo;
	}
	@Override
	public List<CasePriorityVo> genVos(List<TestCasePriority> pos) {
        List<CasePriorityVo> vos = new LinkedList<CasePriorityVo>();

        for (TestCasePriority po: pos) {
        	CasePriorityVo vo = genVo(po);
        	vos.add(vo);
        }
		return vos;
	}

}