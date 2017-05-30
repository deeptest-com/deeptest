package com.ngtesting.platform.service.impl;

import java.util.LinkedList;
import java.util.List;

import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.ngtesting.platform.entity.TestOrgGroup;
import com.ngtesting.platform.entity.TestUser;
import com.ngtesting.platform.service.OrgGroupService;
import com.ngtesting.platform.service.RelationOrgGroupUserService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.util.StringUtil;
import com.ngtesting.platform.vo.OrgGroupVo;
import com.ngtesting.platform.vo.Page;

@Service
public class OrgGroupServiceImpl extends BaseServiceImpl implements OrgGroupService {

	@Override
	public Page listByPage(Long orgId, String keywords, String disabled, Integer currentPage, Integer itemsPerPage) {
        DetachedCriteria dc = DetachedCriteria.forClass(TestOrgGroup.class);
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
	public TestOrgGroup save(OrgGroupVo vo, Long orgId) {
		if (vo == null) {
			return null;
		}
		
		TestOrgGroup po = new TestOrgGroup();
		if (vo.getId() != null) {
			po = (TestOrgGroup) get(TestOrgGroup.class, vo.getId());
		}
		
		po.setName(vo.getName());
		po.setDescr(vo.getDescr());
		po.setDisabled(vo.getDisabled());
		po.setOrgId(orgId);
		
		saveOrUpdate(po);
		return po;
	}

	@Override
	public boolean delete(Long id) {
		TestOrgGroup po = (TestOrgGroup) get(TestOrgGroup.class, id);
		po.setDeleted(true);
		saveOrUpdate(po);
		
		return true;
	}

	@Override
	public OrgGroupVo genVo(TestOrgGroup group) {
		OrgGroupVo vo = new OrgGroupVo();
		BeanUtilEx.copyProperties(vo, group);
		
		return vo;
	}
	@Override
	public List<OrgGroupVo> genVos(List<TestOrgGroup> pos) {
        List<OrgGroupVo> vos = new LinkedList<OrgGroupVo>();

        for (TestOrgGroup po: pos) {
        	OrgGroupVo vo = genVo(po);
        	vos.add(vo);
        }
		return vos;
	}
}
