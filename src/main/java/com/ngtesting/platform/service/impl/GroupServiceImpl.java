package com.ngtesting.platform.service.impl;

import java.util.Date;
import java.util.LinkedList;
import java.util.List;
import java.util.UUID;

import org.apache.commons.lang.StringUtils;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import com.ngtesting.platform.entity.EvtGuest;
import com.ngtesting.platform.entity.SysCompany;
import com.ngtesting.platform.entity.SysGroup;
import com.ngtesting.platform.entity.SysUser;
import com.ngtesting.platform.entity.SysUser.AgentType;
import com.ngtesting.platform.entity.SysVerifyCode;
import com.ngtesting.platform.service.GroupService;
import com.ngtesting.platform.service.UserService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.util.DateUtils;
import com.ngtesting.platform.util.StringUtil;
import com.ngtesting.platform.vo.CompanyVo;
import com.ngtesting.platform.vo.GroupVo;
import com.ngtesting.platform.vo.GuestVo;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.UserVo;

@Service
public class GroupServiceImpl extends BaseServiceImpl implements GroupService {

	@Override
	public Page listByPage(Long companyId, String keywords, String disabled, Integer currentPage, Integer itemsPerPage) {
        DetachedCriteria dc = DetachedCriteria.forClass(SysGroup.class);
        dc.add(Restrictions.eq("companyId", companyId));
        
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
	public SysGroup save(GroupVo vo, Long companyId) {
		if (vo == null) {
			return null;
		}
		
		SysGroup po = new SysGroup();
		if (vo.getId() != null) {
			po = (SysGroup) get(SysGroup.class, vo.getId());
		}
		
		po.setName(vo.getName());
		po.setDescr(vo.getDescr());
		po.setDisabled(vo.getDisabled());
		po.setCompanyId(companyId);
		
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
	public GroupVo genVo(SysGroup group) {
		GroupVo vo = new GroupVo();
		BeanUtilEx.copyProperties(vo, group);
		
		return vo;
	}
	@Override
	public List<GroupVo> genVos(List<SysGroup> pos) {
        List<GroupVo> vos = new LinkedList<GroupVo>();

        for (SysGroup po: pos) {
        	GroupVo vo = genVo(po);
        	vos.add(vo);
        }
		return vos;
	}
}
