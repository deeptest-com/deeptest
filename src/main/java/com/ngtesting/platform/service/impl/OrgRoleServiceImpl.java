package com.ngtesting.platform.service.impl;

import java.util.LinkedList;
import java.util.List;

import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import com.ngtesting.platform.entity.SysOrgRole;
import com.ngtesting.platform.entity.SysUser;
import com.ngtesting.platform.service.OrgRoleService;
import com.ngtesting.platform.service.RoleService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.util.StringUtil;
import com.ngtesting.platform.vo.OrgRoleVo;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.RoleVo;

@Service
public class OrgRoleServiceImpl extends BaseServiceImpl implements OrgRoleService {

	@Override
	public Page listByPage(Long orgId, String keywords, String disabled, Integer currentPage, Integer itemsPerPage) {
        DetachedCriteria dc = DetachedCriteria.forClass(SysOrgRole.class);
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
	public SysOrgRole save(OrgRoleVo vo, Long orgId) {
		if (vo == null) {
			return null;
		}
		
		SysOrgRole po = new SysOrgRole();
		if (vo.getId() != null) {
			po = (SysOrgRole) get(SysOrgRole.class, vo.getId());
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
		SysOrgRole po = (SysOrgRole) get(SysOrgRole.class, id);
		po.setDeleted(true);
		saveOrUpdate(po);
		
		return true;
	}
    
	@Override
	public OrgRoleVo genVo(SysOrgRole role) {
		OrgRoleVo vo = new OrgRoleVo();
		BeanUtilEx.copyProperties(vo, role);
		
		return vo;
	}
	@Override
	public List<OrgRoleVo> genVos(List<SysOrgRole> pos) {
        List<OrgRoleVo> vos = new LinkedList<OrgRoleVo>();

        for (SysOrgRole po: pos) {
        	OrgRoleVo vo = genVo(po);
        	vos.add(vo);
        }
		return vos;
	}
}
