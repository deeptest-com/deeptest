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
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.util.StringUtil;
import com.ngtesting.platform.vo.GroupVo;
import com.ngtesting.platform.vo.Page;

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
	public List<GroupVo> listByUser(Long companyId, Long userId){
        DetachedCriteria dc = DetachedCriteria.forClass(SysGroup.class);
        dc.add(Restrictions.eq("companyId", companyId));
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.addOrder(Order.asc("id"));
        List<SysGroup> allGroups = findAllByCriteria(dc);
        
        List<SysGroupUser> userGroups = this.listUserGroups(companyId, userId);
        List<GroupVo> vos = new LinkedList<GroupVo>();
        
        for (SysGroup group : allGroups) {
        	GroupVo vo = genVo(group);
        	
        	vo.setSelected(false);
        	vo.setSelecting(false);
        	for (SysGroupUser po : userGroups) {
        		if (po.getGroupId() == group.getId()) {
            		vo.setSelected(true);
            		vo.setSelecting(true);
            	}
        	}
        	vos.add(vo);
        }

		return vos;
	}

	@Override
	public boolean saveGroupsByUser(List<GroupVo> groups, Long companyId, Long userId) {
		SysUser user = (SysUser)get(SysUser.class, userId);
		
		for (Object obj: groups) {
			GroupVo groupVo = JSON.parseObject(JSON.toJSONString(obj), GroupVo.class);
			if (groupVo.getSelecting() != groupVo.getSelected()) { // 变化了
				SysGroupUser userGroup;
    			if (groupVo.getSelecting()) { // 勾选
    				userGroup = new SysGroupUser(companyId, userId, groupVo.getId());
    				saveOrUpdate(userGroup);
    			} else { // 取消
    				userGroup = getGroupUser(companyId, userId, groupVo.getId());
    				getDao().delete(userGroup);
    			}
			}
		}
		saveOrUpdate(user);
		
		return true;
	}
	
	@Override
	public SysGroupUser getGroupUser(Long companyId, Long userId, Long groupId) {
		DetachedCriteria dc = DetachedCriteria.forClass(SysGroupUser.class);
        dc.add(Restrictions.eq("companyId", companyId));
        dc.add(Restrictions.eq("userId", userId));
        dc.add(Restrictions.eq("groupId", groupId));
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        
        dc.addOrder(Order.asc("id"));
        List<SysGroupUser> ls = findAllByCriteria(dc);
        
        if (ls.size() == 0) {
        	return null;
        }
		return ls.get(0);
	}

	@Override
	public List<SysGroupUser> listUserGroups(Long companyId, Long userId) {
		DetachedCriteria dc = DetachedCriteria.forClass(SysGroupUser.class);
        dc.add(Restrictions.eq("companyId", companyId));
        dc.add(Restrictions.eq("userId", userId));
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        
        dc.addOrder(Order.asc("id"));
        List<SysGroupUser> ls = findAllByCriteria(dc);
        
		return ls;
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
