package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.entity.TestOrg;
import com.ngtesting.platform.entity.TestUser;
import com.ngtesting.platform.service.AccountService;
import com.ngtesting.platform.service.RelationOrgGroupUserService;
import com.ngtesting.platform.service.RelationProjectRoleUserService;
import com.ngtesting.platform.service.UserService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.util.StringUtil;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.UserVo;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;

@Service
public class UserServiceImpl extends BaseServiceImpl implements UserService {
	
	@Autowired
	AccountService accountService;
	@Autowired
	RelationProjectRoleUserService relationProjectRoleUserService;
	@Autowired
	RelationOrgGroupUserService relationOrgGroupUserService;

	@Override
	public Page listByPage(Long orgId, String keywords, String disabled, Integer currentPage, Integer itemsPerPage) {
        DetachedCriteria dc = DetachedCriteria.forClass(TestUser.class);
        
        dc.createAlias("orgSet", "companies");
        dc.add(Restrictions.eq("companies.id", orgId));
        
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
	public TestUser save(UserVo userVo, Long orgId) {
		if (userVo == null) {
			return null;
		}
		
		TestUser temp = accountService.getByEmail(userVo.getEmail());
		if (temp != null && temp.getId() != userVo.getId()) {
			return null;
		}
		
		TestUser po;
		if (userVo.getId() != null) {
			po = (TestUser) get(TestUser.class, userVo.getId());
		} else {
			po = new TestUser();
			po.setDefaultOrgId(orgId);
		}
		
		po.setName(userVo.getName());
		po.setPhone(userVo.getPhone());
		po.setEmail(userVo.getEmail());
		po.setDisabled(userVo.getDisabled());
		if (userVo.getAvatar() == null) {
			po.setAvatar("upload/sample/user/avatar.png");
		}
		
		TestOrg org = (TestOrg)get(TestOrg.class, orgId);
		if (!contains(org.getUserSet(), userVo.getId())) {
			org.getUserSet().add(po);
			saveOrUpdate(org);
		}
		
		saveOrUpdate(po);
		return po;
	}
	
	@Override
	public boolean disable(Long userId, Long orgId) {
		TestUser po = (TestUser) get(TestUser.class, userId);
		po.setDisabled(!po.getDisabled());
		saveOrUpdate(po);
		
		return true;
	}

	@Override
	public boolean remove(Long userId, Long orgId) {
		TestUser po = (TestUser) get(TestUser.class, userId);
		po.setDeleted(true);
		saveOrUpdate(po);
		
		return true;
	}

	public boolean setSizePers(Long userId, Integer left, Integer right) {
        TestUser po = (TestUser) get(TestUser.class, userId);
        po.setCaseBoardLeftSize(left);
        po.setCaseBoardRightSize(right);
        saveOrUpdate(po);
		return true;
	}
    
	@Override
	public UserVo genVo(TestUser user) {
		if (user == null) {
			return null;
		}
		UserVo vo = new UserVo();
		BeanUtilEx.copyProperties(vo, user);
		
		return vo;
	}
	@Override
	public List<UserVo> genVos(List<TestUser> pos) {
        List<UserVo> vos = new LinkedList<UserVo>();

        for (TestUser po: pos) {
        	UserVo vo = genVo(po);
        	vos.add(vo);
        }
		return vos;
	}

}
