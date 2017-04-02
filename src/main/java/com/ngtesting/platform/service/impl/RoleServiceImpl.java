package com.ngtesting.platform.service.impl;

import java.util.LinkedList;
import java.util.List;

import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import com.ngtesting.platform.entity.SysRole;
import com.ngtesting.platform.entity.SysUser;
import com.ngtesting.platform.service.RoleService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.util.StringUtil;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.RoleVo;

@Service
public class RoleServiceImpl extends BaseServiceImpl implements RoleService {

	@Override
	public Page listByPage(Long orgId, String keywords, String disabled, Integer currentPage, Integer itemsPerPage) {
        DetachedCriteria dc = DetachedCriteria.forClass(SysRole.class);
        dc.add(Restrictions.eq("orgId", orgId));
        
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
	public SysRole save(RoleVo vo, Long orgId) {
		if (vo == null) {
			return null;
		}
		
		SysRole po = new SysRole();
		if (vo.getId() != null) {
			po = (SysRole) get(SysRole.class, vo.getId());
		}
		
		po.setName(vo.getName());
		po.setDescr(vo.getDescr());
		po.setOrgId(orgId);
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
