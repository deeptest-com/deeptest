package com.ngtesting.platform.service.impl;

import java.util.Date;
import java.util.LinkedList;
import java.util.List;
import java.util.UUID;

import org.apache.commons.lang.StringUtils;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import com.ngtesting.platform.entity.EvtGuest;
import com.ngtesting.platform.entity.SysCompany;
import com.ngtesting.platform.entity.SysRole;
import com.ngtesting.platform.entity.SysUser;
import com.ngtesting.platform.entity.SysUser.AgentType;
import com.ngtesting.platform.entity.SysVerifyCode;
import com.ngtesting.platform.service.RoleService;
import com.ngtesting.platform.service.UserService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.util.DateUtils;
import com.ngtesting.platform.util.StringUtil;
import com.ngtesting.platform.vo.CompanyVo;
import com.ngtesting.platform.vo.GuestVo;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.RoleVo;
import com.ngtesting.platform.vo.UserVo;

@Service
public class RoleServiceImpl extends BaseServiceImpl implements RoleService {

	@Override
	public Page listByPage(Long companyId, String keywords, String disabled, Integer currentPage, Integer itemsPerPage) {
        DetachedCriteria dc = DetachedCriteria.forClass(SysRole.class);
        dc.add(Restrictions.eq("companyId", companyId));
        
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        
        if (StringUtil.isNotEmpty(keywords)) {
        	dc.add(Restrictions.like("name", "%" + keywords + "%"));
        }
        if (StringUtil.isNotEmpty(disabled)) {
        	dc.add(Restrictions.eq("disabled", Boolean.valueOf(disabled)));
        }
        
        dc.addOrder(Order.asc("id"));
        Page page = findPage(dc, currentPage * itemsPerPage, itemsPerPage);
		
		return page;
	}

	@Override
	public SysRole save(RoleVo vo, Long companyId) {
		if (vo == null) {
			return null;
		}
		
		SysRole po = new SysRole();
		if (vo.getId() != null) {
			po = (SysRole) get(SysRole.class, vo.getId());
		}
		
		po.setName(vo.getName());
		po.setDescr(vo.getDescr());
		po.setCompanyId(companyId);
		po.setDisabled(vo.getDisabled());
		
		saveOrUpdate(po);
		return po;
	}

	@Override
	public boolean delete(Long id) {
		SysUser po = (SysUser) get(SysUser.class, id);
		po.setDeleted(true);
		saveOrUpdate(po);
		
		return true;
	}

	@Override
	public boolean disable(Long id) {
		SysRole po = (SysRole) get(SysRole.class, id);
		po.setDisabled(!po.getDisabled());
		saveOrUpdate(po);
		
		return true;
	}
    
	@Override
	public RoleVo genVo(SysRole role) {
		RoleVo vo = new RoleVo();
		BeanUtilEx.copyProperties(vo, role);
		
		return vo;
	}
	@Override
	public List<RoleVo> genVos(List<SysRole> pos) {
        List<RoleVo> vos = new LinkedList<RoleVo>();

        for (SysRole po: pos) {
        	RoleVo vo = genVo(po);
        	vos.add(vo);
        }
		return vos;
	}
}
