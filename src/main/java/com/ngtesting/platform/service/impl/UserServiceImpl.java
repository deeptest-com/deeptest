package com.ngtesting.platform.service.impl;

import java.util.LinkedList;
import java.util.List;

import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.alibaba.fastjson.JSON;
import com.ngtesting.platform.entity.SysGroup;
import com.ngtesting.platform.entity.SysGroupUser;
import com.ngtesting.platform.entity.SysUser;
import com.ngtesting.platform.service.GroupService;
import com.ngtesting.platform.service.UserGroupService;
import com.ngtesting.platform.service.UserService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.util.StringUtil;
import com.ngtesting.platform.vo.GroupVo;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.UserVo;

@Service
public class UserServiceImpl extends BaseServiceImpl implements UserService {
	
	@Autowired
	UserGroupService userGroupService;

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
		po.setDisabled(vo.getDisabled());
		po.setCompanyId(companyId);
		if (vo.getAvatar() == null) {
			po.setAvatar("upload/sample/user/avatar.png");
		}
		
		saveOrUpdate(po);
		return po;
	}
	
	@Override
	public boolean disable(Long id) {
		SysUser po = (SysUser) get(SysUser.class, id);
		po.setDisabled(!po.getDisabled());
		saveOrUpdate(po);
		
		return true;
	}

	@Override
	public boolean delete(Long id) {
		SysUser po = (SysUser) get(SysUser.class, id);
		po.setDeleted(true);
		saveOrUpdate(po);
		
		return true;
	}
	
	@Override
	public List<UserVo> listByGroup(Long companyId, Long groupId){
        DetachedCriteria dc = DetachedCriteria.forClass(SysUser.class);
        dc.add(Restrictions.eq("companyId", companyId));
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.addOrder(Order.asc("id"));
        List<SysUser> allUsers = findAllByCriteria(dc);
        
        List<SysGroupUser> userGroups = userGroupService.listUserGroups(companyId, null, groupId);
        List<UserVo> vos = new LinkedList<UserVo>();
        
        for (SysUser user : allUsers) {
        	UserVo vo = genVo(user);
        	
        	vo.setSelected(false);
        	vo.setSelecting(false);
        	for (SysGroupUser po : userGroups) {
        		if (po.getUserId() == user.getId()) {
            		vo.setSelected(true);
            		vo.setSelecting(true);
            	}
        	}
        	vos.add(vo);
        }

		return vos;
	}

	@Override
	public boolean saveUsersByGroup(List<UserVo> users, Long companyId, Long groupId) {
		for (Object obj: users) {
			UserVo userVo = JSON.parseObject(JSON.toJSONString(obj), UserVo.class);
			if (userVo.getSelecting() != userVo.getSelected()) { // 变化了
				SysGroupUser userGroup;
				userGroup = userGroupService.getGroupUser(companyId, userVo.getId(), groupId);
				
    			if (userVo.getSelecting() && userGroup == null) { // 勾选
    				userGroup = new SysGroupUser(companyId, userVo.getId(), groupId);
    				saveOrUpdate(userGroup);
    			} else if (userGroup != null) { // 取消
    				getDao().delete(userGroup);
    			}
			}
		}
		
		return true;
	}
    
	@Override
	public UserVo genVo(SysUser user) {
		if (user == null) {
			return null;
		}
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
