package com.ngtesting.platform.service.impl;

import java.util.LinkedList;
import java.util.List;

import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import com.ngtesting.platform.entity.SysOrg;
import com.ngtesting.platform.service.OrgService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.vo.OrgVo;

@Service
public class OrgServiceImpl extends BaseServiceImpl implements OrgService {

	@Override
	public List<SysOrg> list(String keywords, Boolean disabled, Long userId) {
        DetachedCriteria dc = DetachedCriteria.forClass(SysOrg.class);
        dc.createAlias("userSet", "users");
        dc.add(Restrictions.eq("users.id", userId));
        
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.addOrder(Order.asc("id"));
        List<SysOrg> ls = findAllByCriteria(dc);

		return ls;
	}

	@Override
	public SysOrg getDetail(Long id) {
		if (id == null) {
			return null;
		}
		SysOrg po = (SysOrg) get(SysOrg.class, id);

		return po;
	}

	@Override
	public SysOrg save(OrgVo vo, Long userId) {
		if (vo == null) {
			return null;
		}

		boolean isNew = vo.getId() == null;
		SysOrg po = new SysOrg();
		if (!isNew) {
			po = (SysOrg) get(SysOrg.class, vo.getId());
		} else {
			po.setAdminId(userId);
		}
		
		boolean disableChanged = vo.getDisabled() != po.getDisabled();
		
		po.setName(vo.getName());
		po.setWebsite(vo.getWebsite());
		po.setDisabled(vo.getDisabled());

		saveOrUpdate(po);
		
		return po;
	}

	@Override
	public Boolean disable(Long id) {
		if (id == null) {
			return false;
		}

		SysOrg po = (SysOrg) get(SysOrg.class, id);
		po.setDisabled(true);
		saveOrUpdate(po);

		return true;
	}
	
	@Override
	public Boolean delete(Long id) {
		if (id == null) {
			return false;
		}

		SysOrg po = (SysOrg) get(SysOrg.class, id);
		po.setDeleted(true);
		saveOrUpdate(po);

		return true;
	}

	@Override
	public List<OrgVo> genVos(List<SysOrg> pos) {
		List<OrgVo> voList = new LinkedList<OrgVo>();
		for (SysOrg po : pos) {
			OrgVo vo = genVo(po);
			voList.add(vo);
		}
		
		return voList;
	}

	@Override
	public OrgVo genVo(SysOrg po) {
		if (po == null) {
			return null;
		}
		OrgVo vo = new OrgVo();
		BeanUtilEx.copyProperties(vo, po);

		return vo;
	}

}
