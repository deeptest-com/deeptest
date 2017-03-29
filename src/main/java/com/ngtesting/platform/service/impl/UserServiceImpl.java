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
import com.ngtesting.platform.entity.SysUser;
import com.ngtesting.platform.entity.SysUser.AgentType;
import com.ngtesting.platform.entity.SysVerifyCode;
import com.ngtesting.platform.service.UserService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.util.DateUtils;
import com.ngtesting.platform.util.StringUtil;
import com.ngtesting.platform.vo.CompanyVo;
import com.ngtesting.platform.vo.GuestVo;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.UserVo;

@Service
public class UserServiceImpl extends BaseServiceImpl implements UserService {

	@Override
	public Page listByPage(Long companyId, String keywords, String disabled, Integer currentPage, Integer itemsPerPage) {
        DetachedCriteria dc = DetachedCriteria.forClass(SysUser.class);
        dc.add(Restrictions.eq("companyId", companyId));
        
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        
        if (StringUtil.isNotEmpty(keywords)) {
        	dc.add(Restrictions.or(Restrictions.like("name", "%" + keywords + "%"),
        		   Restrictions.like("email", "%" + keywords + "%"),
        		   Restrictions.like("phone", "%" + keywords + "%") ));
        }
        if (StringUtil.isNotEmpty(disabled)) {
        	dc.add(Restrictions.eq("disabled", Boolean.valueOf(disabled)));
        }
        
        dc.addOrder(Order.asc("id"));
        Page page = findPage(dc, currentPage * itemsPerPage, itemsPerPage);
		
		return page;
	}

	@Override
	public SysUser save(UserVo vo, Long companyId) {
		if (vo == null) {
			return null;
		}
		
		SysUser po = new SysUser();
		if (vo.getId() != null) {
			po = (SysUser) get(SysUser.class, vo.getId());
		}
		
		po.setName(vo.getName());
		po.setPhone(vo.getPhone());
		po.setEmail(vo.getEmail());
		po.setCompanyId(companyId);
		if (vo.getAvatar() == null) {
			po.setAvatar("upload/sample/user/avatar.png");
		}
		
		saveOrUpdate(po);
		return po;
	}

	@Override
	public boolean remove(Long id) {
		SysUser po = (SysUser) get(SysUser.class, id);
		po.setDeleted(true);
		saveOrUpdate(po);
		
		return true;
	}

	@Override
	public boolean disablePers(Long id) {
		SysUser po = (SysUser) get(SysUser.class, id);
		po.setDisabled(!po.getDisabled());
		saveOrUpdate(po);
		
		return true;
	}
    
	@Override
	public UserVo genVo(SysUser user) {
		UserVo vo = new UserVo();
		BeanUtilEx.copyProperties(vo, user);
		
		return vo;
	}
	@Override
	public List<UserVo> genVos(List<SysUser> pos) {
        List<UserVo> vos = new LinkedList<UserVo>();

        for (SysUser po: pos) {
        	UserVo vo = genVo(po);
        	vos.add(vo);
        }
		return vos;
	}
}
