package com.ngtesting.platform.service.impl;

import java.util.LinkedList;
import java.util.List;

import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import com.ngtesting.platform.entity.TestProjectRole;
import com.ngtesting.platform.service.ProjectRoleService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.util.StringUtil;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.ProjectRoleVo;

@Service
public class ProjectRoleServiceImpl extends BaseServiceImpl implements ProjectRoleService {

	@Override
	public List list(Long orgId, String keywords, String disabled) {
        DetachedCriteria dc = DetachedCriteria.forClass(TestProjectRole.class);
        dc.add(Restrictions.eq("orgId", orgId));
        
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        
        if (StringUtil.isNotEmpty(keywords)) {
        	dc.add(Restrictions.like("name", "%" + keywords + "%"));
        }
        if (StringUtil.isNotEmpty(disabled)) {
        	dc.add(Restrictions.eq("disabled", Boolean.valueOf(disabled)));
        }
        
        dc.addOrder(Order.asc("id"));
        List ls = findAllByCriteria(dc);
		
		return ls;
	}

	@Override
	public TestProjectRole save(ProjectRoleVo vo, Long orgId) {
		if (vo == null) {
			return null;
		}
		
		TestProjectRole po = new TestProjectRole();
		if (vo.getId() != null) {
			po = (TestProjectRole) get(TestProjectRole.class, vo.getId());
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
		TestProjectRole po = (TestProjectRole) get(TestProjectRole.class, id);
		po.setDeleted(true);
		saveOrUpdate(po);
		
		return true;
	}
    
	@Override
	public ProjectRoleVo genVo(TestProjectRole role) {
		ProjectRoleVo vo = new ProjectRoleVo();
		BeanUtilEx.copyProperties(vo, role);
		
		return vo;
	}
	@Override
	public List<ProjectRoleVo> genVos(List<TestProjectRole> pos) {
        List<ProjectRoleVo> vos = new LinkedList<ProjectRoleVo>();

        for (TestProjectRole po: pos) {
        	ProjectRoleVo vo = genVo(po);
        	vos.add(vo);
        }
		return vos;
	}
}
