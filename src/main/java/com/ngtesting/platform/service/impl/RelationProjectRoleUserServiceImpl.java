package com.ngtesting.platform.service.impl;

import java.util.List;

import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import com.ngtesting.platform.entity.TestRelationProjectRoleUser;
import com.ngtesting.platform.service.RelationProjectRoleUserService;

@Service
public class RelationProjectRoleUserServiceImpl extends BaseServiceImpl implements RelationProjectRoleUserService {
	@Override
	public TestRelationProjectRoleUser getRelationProjectRoleUser(Long projectRoleId) {
		DetachedCriteria dc = DetachedCriteria.forClass(TestRelationProjectRoleUser.class);
        dc.add(Restrictions.eq("projectRoleId", projectRoleId));
        
        dc.addOrder(Order.asc("id"));
        List<TestRelationProjectRoleUser> ls = findAllByCriteria(dc);
        
        if (ls.size() == 0) {
        	return null;
        }
		return ls.get(0);
	}

	@Override
	public List<TestRelationProjectRoleUser> listRelationProjectRoleUsers(Long projectRoleId) {
		DetachedCriteria dc = DetachedCriteria.forClass(TestRelationProjectRoleUser.class);
        dc.add(Restrictions.eq("projectRoleId", projectRoleId));
        
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        
        dc.addOrder(Order.asc("id"));
        List<TestRelationProjectRoleUser> ls = findAllByCriteria(dc);
        
		return ls;
	}
    
}
