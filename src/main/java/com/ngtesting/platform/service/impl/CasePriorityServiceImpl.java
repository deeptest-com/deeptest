package com.ngtesting.platform.service.impl;

import java.util.LinkedList;
import java.util.List;
import java.util.UUID;

import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.ngtesting.platform.entity.SysCasePriority;
import com.ngtesting.platform.entity.SysCasePriority;
import com.ngtesting.platform.entity.SysCustomField;
import com.ngtesting.platform.entity.SysOrg;
import com.ngtesting.platform.entity.SysUser;
import com.ngtesting.platform.service.AccountService;
import com.ngtesting.platform.service.CasePriorityService;
import com.ngtesting.platform.service.CustomFieldService;
import com.ngtesting.platform.service.RelationOrgGroupUserService;
import com.ngtesting.platform.service.RelationProjectRoleUserService;
import com.ngtesting.platform.service.UserService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.util.StringUtil;
import com.ngtesting.platform.vo.CaseExeStatusVo;
import com.ngtesting.platform.vo.CasePriorityVo;
import com.ngtesting.platform.vo.CustomFieldVo;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.UserVo;

@Service
public class CasePriorityServiceImpl extends BaseServiceImpl implements CasePriorityService {
	@Override
	public List<SysCasePriority> list(Long orgId) {
        DetachedCriteria dc = DetachedCriteria.forClass(SysCasePriority.class);
        
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
	public SysCasePriority save(CasePriorityVo vo, Long orgId) {
		if (vo == null) {
			return null;
		}
		
		SysCasePriority po;
		if (vo.getId() != null) {
			po = (SysCasePriority) get(SysCasePriority.class, vo.getId());
		} else {
			po = new SysCasePriority();
		}
		
		BeanUtilEx.copyProperties(po, vo);
		
		po.setOrgId(orgId);
		
		if (vo.getId() == null) {
			po.setCode(UUID.randomUUID().toString());
			
			String hql = "select max(displayOrder) from SysCasePriority";
			Integer maxOrder = (Integer) getByHQL(hql);
	        po.setDisplayOrder(maxOrder + 10);
		}
		
		saveOrUpdate(po);
		return po;
	}

	@Override
	public boolean delete(Long id) {
		SysCasePriority po = (SysCasePriority) get(SysCasePriority.class, id);
		po.setDeleted(true);
		saveOrUpdate(po);
		
		return true;
	}

	@Override
	public boolean setDefaultPers(Long id, Long orgId) {
		List<SysCasePriority> ls = list(orgId);
		for (SysCasePriority type : ls) {
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
		SysCasePriority type = (SysCasePriority) get(SysCasePriority.class, id);
		
        String hql = "from SysCasePriority tp where tp.deleted = false and tp.disabled = false ";
        if ("up".equals(act)) {
        	hql += "and tp.displayOrder < ? order by displayOrder desc";
        } else if ("down".equals(act)) {
        	hql += "and tp.displayOrder > ? order by displayOrder asc";
        } else {
        	return false;
        }
        
        SysCasePriority neighbor = (SysCasePriority) getDao().findFirstByHQL(hql, type.getDisplayOrder());
		
        Integer order = type.getDisplayOrder();
        type.setDisplayOrder(neighbor.getDisplayOrder());
        neighbor.setDisplayOrder(order);
        
        saveOrUpdate(type);
        saveOrUpdate(neighbor);
		
		return true;
	}
	
    
	@Override
	public CasePriorityVo genVo(SysCasePriority po) {
		if (po == null) {
			return null;
		}
		CasePriorityVo vo = new CasePriorityVo();
		BeanUtilEx.copyProperties(vo, po);
		
		return vo;
	}
	@Override
	public List<CasePriorityVo> genVos(List<SysCasePriority> pos) {
        List<CasePriorityVo> vos = new LinkedList<CasePriorityVo>();

        for (SysCasePriority po: pos) {
        	CasePriorityVo vo = genVo(po);
        	vos.add(vo);
        }
		return vos;
	}

}