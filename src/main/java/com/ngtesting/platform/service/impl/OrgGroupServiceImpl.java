package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSONArray;
import com.ngtesting.platform.model.TstOrgGroup;
import com.ngtesting.platform.service.OrgGroupService;
import com.ngtesting.platform.vo.Page;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;

@Service
public class OrgGroupServiceImpl extends BaseServiceImpl implements OrgGroupService {

	@Override
	public Page listByPage(Long orgId, String keywords, String disabled, Integer currentPage, Integer itemsPerPage) {
//        DetachedCriteria dc = DetachedCriteria.forClass(TstOrgGroup.class);
//        dc.add(Restrictions.eq("orgId", orgId));
//
//        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//
//        if (StringUtil.isNotEmpty(keywords)) {
//        	dc.add(Restrictions.like("name", "%" + keywords + "%"));
//        }
//        if (StringUtil.isNotEmpty(disabled)) {
//        	dc.add(Restrictions.eq("disabled", Boolean.valueOf(disabled)));
//        }
//
//        dc.addOrder(Order.asc("id"));
//        Page page = findPage(dc, currentPage * itemsPerPage, itemsPerPage);
//
//		return page;

		return null;
	}

	@Override
	public List search(Long orgId, String keywords, JSONArray exceptIds) {
//		DetachedCriteria dc = DetachedCriteria.forClass(TstOrgGroup.class);
//
//		dc.add(Restrictions.eq("orgId", orgId));
//
//		List<Long> ids = new ArrayList();
//		for (Object json : exceptIds.toArray()) {
//			ids.add(Long.valueOf(json.toString()));
//		}
//
//		if (exceptIds.size() > 0) {
//			dc.add(Restrictions.not(Restrictions.in("id", ids)));
//		}
//
//		dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//		dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//
//		if (StringUtil.isNotEmpty(keywords)) {
//			dc.add(Restrictions.like("name", "%" + keywords + "%"));
//		}
//
//		dc.addOrder(Order.asc("id"));
//		Page page = findPage(dc, 0, 20);
//
//		return page.getItems();

		return null;
	}

	@Override
	public TstOrgGroup save(TstOrgGroup vo, Long orgId) {
//		if (vo == null) {
//			return null;
//		}
//
//		TstOrgGroup po = new TstOrgGroup();
//		if (vo.getId() != null) {
//			po = (TstOrgGroup) get(TstOrgGroup.class, vo.getId());
//		}
//
//		po.setName(vo.getName());
//		po.setDescr(vo.getDescr());
//		po.setDisabled(vo.getDisabled());
//		po.setOrgId(orgId);
//
//		saveOrUpdate(po);
//		return po;

		return null;
	}

	@Override
	public boolean delete(Long id) {
//		TstOrgGroup po = (TstOrgGroup) get(TstOrgGroup.class, id);
//		po.setDeleted(true);
//		saveOrUpdate(po);

		return true;
	}

//	@Override
//	public void initDefaultBasicDataPers(TestOrg org) {
//		String [] groups = new String[]{"测试主管","测试设计","测试执行"};
//		for(String name : groups) {
//			TstOrgGroup po = new TstOrgGroup();
//			po.setName(name);
//            po.setOrgId(org.getId());
//			saveOrUpdate(po);
//		}
//	}

	@Override
	public TstOrgGroup genVo(TstOrgGroup group) {
		TstOrgGroup vo = new TstOrgGroup();
//		BeanUtilEx.copyProperties(vo, group);

		return vo;
	}
	@Override
	public List<TstOrgGroup> genVos(List<TstOrgGroup> pos) {
        List<TstOrgGroup> vos = new LinkedList<TstOrgGroup>();

//        for (TstOrgGroup po: pos) {
//        	TstOrgGroup vo = genVo(po);
//        	vos.add(vo);
//        }
		return vos;
	}
}
