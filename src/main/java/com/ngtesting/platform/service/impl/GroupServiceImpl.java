package com.ngtesting.platform.service.impl;

import java.util.LinkedList;
import java.util.List;

import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.ngtesting.platform.entity.SysOrgGroup;
import com.ngtesting.platform.entity.SysUser;
import com.ngtesting.platform.service.GroupService;
import com.ngtesting.platform.service.RelationOrgGroupUserService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.util.StringUtil;
import com.ngtesting.platform.vo.GroupVo;
import com.ngtesting.platform.vo.Page;

@Service
public class GroupServiceImpl extends BaseServiceImpl implements GroupService {
	
	@Autowired
	RelationOrgGroupUserService userGroupService;

	@Override
	public Page listByPage(Long orgId, String keywords, String disabled, Integer currentPage, Integer itemsPerPage) {
        DetachedCriteria dc = DetachedCriteria.forClass(SysOrgGroup.class);
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
	public SysOrgGroup save(GroupVo vo, Long orgId) {
		if (vo == null) {
			return null;
		}
		
		SysOrgGroup po = new SysOrgGroup();
		if (vo.getId() != null) {
			po = (SysOrgGroup) get(SysOrgGroup.class, vo.getId());
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
		SysUser po = (SysUser) get(SysUser.class, id);
		po.setDeleted(true);
		saveOrUpdate(po);
		
		return true;
	}

	@Override
	public boolean disable(Long id) {
		SysUser po = (SysUser) get(SysUser.class, id);
		po.setDisabled(!po.getDisabled());
		saveOrUpdate(po);
		
		return true;
	}

	@Override
	public GroupVo genVo(SysOrgGroup group) {
		GroupVo vo = new GroupVo();
		BeanUtilEx.copyProperties(vo, group);
		
		return vo;
	}
	@Override
	public List<GroupVo> genVos(List<SysOrgGroup> pos) {
        List<GroupVo> vos = new LinkedList<GroupVo>();

        for (SysOrgGroup po: pos) {
        	GroupVo vo = genVo(po);
        	vos.add(vo);
        }
		return vos;
	}
}
