package com.ngtesting.platform.service.impl;

import java.util.LinkedList;
import java.util.List;

import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.ngtesting.platform.entity.SysCaseExeStatus;
import com.ngtesting.platform.entity.SysCustomField;
import com.ngtesting.platform.entity.SysOrg;
import com.ngtesting.platform.entity.SysUser;
import com.ngtesting.platform.service.AccountService;
import com.ngtesting.platform.service.CaseExeStatusService;
import com.ngtesting.platform.service.CustomFieldService;
import com.ngtesting.platform.service.RelationOrgGroupUserService;
import com.ngtesting.platform.service.RelationProjectRoleUserService;
import com.ngtesting.platform.service.UserService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.util.StringUtil;
import com.ngtesting.platform.vo.CaseExeStatusVo;
import com.ngtesting.platform.vo.CaseTypeVo;
import com.ngtesting.platform.vo.CustomFieldVo;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.UserVo;

@Service
public class CaseExeStatusServiceImpl extends BaseServiceImpl implements CaseExeStatusService {
	@Override
	public List<SysCaseExeStatus> list(Long orgId) {
        DetachedCriteria dc = DetachedCriteria.forClass(SysCaseExeStatus.class);
        
        dc.add(Restrictions.eq("orgId", orgId));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        
        dc.addOrder(Order.asc("displayOrder"));
        List ls = findAllByCriteria(dc);
		
		return ls;
	}
	@Override
	public List<CaseExeStatusVo> listVos(Long orgId) {
        List ls = list(orgId);
        
        List<CaseExeStatusVo> vos = genVos(ls);
		return vos;
	}

	@Override
	public SysCaseExeStatus save(CaseExeStatusVo vo, Long orgId) {
		if (vo == null) {
			return null;
		}
		
		SysCaseExeStatus po;
		if (vo.getId() != null) {
			po = (SysCaseExeStatus) get(SysCaseExeStatus.class, vo.getId());
		} else {
			po = new SysCaseExeStatus();
		}
		po.setOrgId(orgId);
		
		BeanUtilEx.copyProperties(po, vo);
		
		saveOrUpdate(po);
		return po;
	}

	@Override
	public boolean delete(Long id) {
		SysCaseExeStatus po = (SysCaseExeStatus) get(SysCaseExeStatus.class, id);
		po.setDeleted(true);
		saveOrUpdate(po);
		
		return true;
	}
    
	@Override
	public CaseExeStatusVo genVo(SysCaseExeStatus po) {
		if (po == null) {
			return null;
		}
		CaseExeStatusVo vo = new CaseExeStatusVo();
		BeanUtilEx.copyProperties(vo, po);
		
		return vo;
	}
	@Override
	public List<CaseExeStatusVo> genVos(List<SysCaseExeStatus> pos) {
        List<CaseExeStatusVo> vos = new LinkedList<CaseExeStatusVo>();

        for (SysCaseExeStatus po: pos) {
        	CaseExeStatusVo vo = genVo(po);
        	vos.add(vo);
        }
		return vos;
	}

}