package com.ngtesting.platform.service.impl;

import java.util.LinkedList;
import java.util.List;

import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import com.ngtesting.platform.entity.SysProjectRole;
import com.ngtesting.platform.service.ProjectRoleService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.util.StringUtil;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.ProjectRoleVo;

@Service
public class ProjectRoleServiceImpl extends BaseServiceImpl implements ProjectRoleService {

	@Override
	public Page listByPage(Long orgId, String keywords, String disabled, Integer currentPage, Integer itemsPerPage) {
        DetachedCriteria dc = DetachedCriteria.forClass(SysProjectRole.class);
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
	public SysProjectRole save(ProjectRoleVo vo, Long orgId) {
		if (vo == null) {
			return null;
		}
		
		SysProjectRole po = new SysProjectRole();
		if (vo.getId() != null) {
			po = (SysProjectRole) get(SysProjectRole.class, vo.getId());
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
		SysProjectRole po = (SysProjectRole) get(SysProjectRole.class, id);
		po.setDeleted(true);
		saveOrUpdate(po);
		
		return true;
	}
    
	@Override
	public ProjectRoleVo genVo(SysProjectRole role) {
		ProjectRoleVo vo = new ProjectRoleVo();
		BeanUtilEx.copyProperties(vo, role);
		
		return vo;
	}
	@Override
	public List<ProjectRoleVo> genVos(List<SysProjectRole> pos) {
        List<ProjectRoleVo> vos = new LinkedList<ProjectRoleVo>();

        for (SysProjectRole po: pos) {
        	ProjectRoleVo vo = genVo(po);
        	vos.add(vo);
        }
		return vos;
	}
}
