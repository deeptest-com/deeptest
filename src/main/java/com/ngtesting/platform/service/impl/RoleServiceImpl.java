package com.ngtesting.platform.service.impl;

import java.util.LinkedList;
import java.util.List;

import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import com.ngtesting.platform.entity.TestRole;
import com.ngtesting.platform.entity.TestUser;
import com.ngtesting.platform.service.RoleService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.util.StringUtil;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.RoleVo;

@Service
public class RoleServiceImpl extends BaseServiceImpl implements RoleService {

	@Override
	public Page listByPage(Long orgId, String keywords, String disabled, Integer currentPage, Integer itemsPerPage) {
        DetachedCriteria dc = DetachedCriteria.forClass(TestRole.class);
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
	public TestRole save(RoleVo vo, Long orgId) {
		if (vo == null) {
			return null;
		}
		
		TestRole po = new TestRole();
		if (vo.getId() != null) {
			po = (TestRole) get(TestRole.class, vo.getId());
		}
		
		po.setName(vo.getName());
		po.setDescr(vo.getDescr());
		po.setDisabled(vo.getDisabled());
		
		saveOrUpdate(po);
		return po;
	}

	@Override
	public boolean delete(Long id) {
		TestUser po = (TestUser) get(TestUser.class, id);
		po.setDeleted(true);
		saveOrUpdate(po);
		
		return true;
	}

	@Override
	public boolean disable(Long id) {
		TestRole po = (TestRole) get(TestRole.class, id);
		po.setDisabled(!po.getDisabled());
		saveOrUpdate(po);
		
		return true;
	}
    
	@Override
	public RoleVo genVo(TestRole role) {
		RoleVo vo = new RoleVo();
		BeanUtilEx.copyProperties(vo, role);
		
		return vo;
	}
	@Override
	public List<RoleVo> genVos(List<TestRole> pos) {
        List<RoleVo> vos = new LinkedList<RoleVo>();

        for (TestRole po: pos) {
        	RoleVo vo = genVo(po);
        	vos.add(vo);
        }
		return vos;
	}
}
