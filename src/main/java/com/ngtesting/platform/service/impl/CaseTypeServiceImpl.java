package com.ngtesting.platform.service.impl;

import java.util.LinkedList;
import java.util.List;

import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.ngtesting.platform.entity.SysCaseType;
import com.ngtesting.platform.entity.SysOrg;
import com.ngtesting.platform.entity.SysUser;
import com.ngtesting.platform.service.AccountService;
import com.ngtesting.platform.service.CaseTypeService;
import com.ngtesting.platform.service.CustomFieldService;
import com.ngtesting.platform.service.RelationOrgGroupUserService;
import com.ngtesting.platform.service.RelationProjectRoleUserService;
import com.ngtesting.platform.service.UserService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.util.StringUtil;
import com.ngtesting.platform.vo.CaseTypeVo;
import com.ngtesting.platform.vo.CaseTypeVo;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.TestProjectAccessHistoryVo;
import com.ngtesting.platform.vo.UserVo;

@Service
public class CaseTypeServiceImpl extends BaseServiceImpl implements CaseTypeService {

	@Override
	public List<SysCaseType> list(Long orgId) {
        DetachedCriteria dc = DetachedCriteria.forClass(SysCaseType.class);
        
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
	public SysCaseType save(CaseTypeVo vo, Long orgId) {
		if (vo == null) {
			return null;
		}
		
		SysCaseType po;
		if (vo.getId() != null) {
			po = (SysCaseType) get(SysCaseType.class, vo.getId());
		} else {
			po = new SysCaseType();
		}
		po.setOrgId(orgId);
		
		BeanUtilEx.copyProperties(po, vo);
		
		saveOrUpdate(po);
		return po;
	}

	@Override
	public boolean delete(Long id) {
		SysCaseType po = (SysCaseType) get(SysCaseType.class, id);
		po.setDeleted(true);
		saveOrUpdate(po);
		
		return true;
	}
    
	@Override
	public CaseTypeVo genVo(SysCaseType po) {
		if (po == null) {
			return null;
		}
		CaseTypeVo vo = new CaseTypeVo();
		BeanUtilEx.copyProperties(vo, po);
		
		return vo;
	}
	@Override
	public List<CaseTypeVo> genVos(List<SysCaseType> pos) {
        List<CaseTypeVo> vos = new LinkedList<CaseTypeVo>();

        for (SysCaseType po: pos) {
        	CaseTypeVo vo = genVo(po);
        	vos.add(vo);
        }
		return vos;
	}

	@Override
	public boolean setDefaultPers(Long id, Long orgId) {
		List<SysCaseType> ls = list(orgId);
		for (SysCaseType type : ls) {
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

}
