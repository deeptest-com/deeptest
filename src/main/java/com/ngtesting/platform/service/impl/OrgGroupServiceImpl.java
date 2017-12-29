package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSONArray;
import com.ngtesting.platform.entity.TestOrg;
import com.ngtesting.platform.entity.TestOrgGroup;
import com.ngtesting.platform.service.OrgGroupService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.util.StringUtil;
import com.ngtesting.platform.vo.OrgGroupVo;
import com.ngtesting.platform.vo.Page;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;

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
	public List search(Long orgId, String keywords, JSONArray exceptIds) {
		DetachedCriteria dc = DetachedCriteria.forClass(TestOrgGroup.class);

		dc.add(Restrictions.eq("orgId", orgId));

		List<Long> ids = new ArrayList();
		for (Object json : exceptIds.toArray()) {
			ids.add(Long.valueOf(json.toString()));
		}

		if (exceptIds.size() > 0) {
			dc.add(Restrictions.not(Restrictions.in("id", ids)));
		}

		dc.add(Restrictions.eq("deleted", Boolean.FALSE));
		dc.add(Restrictions.eq("disabled", Boolean.FALSE));

		if (StringUtil.isNotEmpty(keywords)) {
			dc.add(Restrictions.like("name", "%" + keywords + "%"));
		}

		dc.addOrder(Order.asc("id"));
		Page page = findPage(dc, 0, 20);

		return page.getItems();
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
	public void initDefaultBasicDataPers(TestOrg org) {
		String [] groups = new String[]{"测试主管","测试设计","测试执行"};
		for(String name : groups) {
			TestOrgGroup po = new TestOrgGroup();
			po.setName(name);
            po.setOrgId(org.getId());
			saveOrUpdate(po);
		}
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
