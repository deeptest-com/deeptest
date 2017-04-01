package com.ngtesting.platform.service.impl;

import java.util.Iterator;
import java.util.LinkedList;
import java.util.List;
import java.util.Set;

import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.SysGroup;
import com.ngtesting.platform.entity.SysGroupUser;
import com.ngtesting.platform.entity.SysUser;
import com.ngtesting.platform.service.GroupService;
import com.ngtesting.platform.service.UserGroupService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.util.StringUtil;
import com.ngtesting.platform.vo.GroupVo;
import com.ngtesting.platform.vo.Page;

@Service
public class UserGroupServiceImpl extends BaseServiceImpl implements UserGroupService {
	@Override
	public SysGroupUser getGroupUser(Long companyId, Long userId, Long groupId) {
		DetachedCriteria dc = DetachedCriteria.forClass(SysGroupUser.class);
        dc.add(Restrictions.eq("companyId", companyId));
        dc.add(Restrictions.eq("userId", userId));
        dc.add(Restrictions.eq("groupId", groupId));
//        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        
        dc.addOrder(Order.asc("id"));
        List<SysGroupUser> ls = findAllByCriteria(dc);
        
        if (ls.size() == 0) {
        	return null;
        }
		return ls.get(0);
	}

	@Override
	public List<SysGroupUser> listUserGroups(Long companyId, Long userId, Long groupId) {
		DetachedCriteria dc = DetachedCriteria.forClass(SysGroupUser.class);
        dc.add(Restrictions.eq("companyId", companyId));
        
        if (userId != null) {
        	dc.add(Restrictions.eq("userId", userId));
        }
        if (groupId != null) {
        	dc.add(Restrictions.eq("groupId", groupId));
        }
        
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        
        dc.addOrder(Order.asc("id"));
        List<SysGroupUser> ls = findAllByCriteria(dc);
        
		return ls;
	}
    
}
