package com.ngtesting.platform.service.impl;

import java.util.LinkedList;
import java.util.List;

import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import com.alibaba.fastjson.JSON;
import com.ngtesting.platform.entity.SysOrgGroup;
import com.ngtesting.platform.entity.SysRelationOrgGroupUser;
import com.ngtesting.platform.entity.SysUser;
import com.ngtesting.platform.service.RelationOrgGroupUserService;
import com.ngtesting.platform.vo.RelationOrgGroupUserVo;

@Service
public class RelationOrgGroupUserServiceImpl extends BaseServiceImpl implements RelationOrgGroupUserService {

	@Override
	public List<RelationOrgGroupUserVo> listRelationsOrgGroupUsers(Long orgId, Long orgGroupId, Long userId) {

        List<SysOrgGroup> allOrgGroups = listAllOrgGroups(orgId);
        
        List<SysRelationOrgGroupUser> relations = listRelations(orgGroupId, userId);
        
        List<RelationOrgGroupUserVo> vos = new LinkedList<RelationOrgGroupUserVo>();
        for (SysOrgGroup orgGroup : allOrgGroups) {
        	RelationOrgGroupUserVo vo = genVo(orgGroupId, userId);
        	
        	vo.setSelected(false);
        	vo.setSelecting(false);
        	for (SysRelationOrgGroupUser po : relations) {
        		if (po.getOrgGroupId() == orgGroup.getId() && po.getUserId() == userId) {
            		vo.setSelected(true);
            		vo.setSelecting(true);
            	}
        	}
        	vos.add(vo);
        }
        
		return vos;
	}

	private List<SysRelationOrgGroupUser> listRelations(Long orgGroupId, Long userId) {
		DetachedCriteria dc2 = DetachedCriteria.forClass(SysRelationOrgGroupUser.class);
        if (orgGroupId != null) {
        	dc2.add(Restrictions.eq("orgGroupId", orgGroupId));
        }
        if (userId != null) {
        	dc2.add(Restrictions.eq("userId", userId));
        }
        dc2.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc2.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc2.addOrder(Order.asc("id"));
        List<SysRelationOrgGroupUser> relations = findAllByCriteria(dc2);
        
		return relations;
	}

	private List<SysOrgGroup> listAllOrgGroups(Long orgId) {
		DetachedCriteria dc = DetachedCriteria.forClass(SysOrgGroup.class);
        dc.add(Restrictions.eq("orgId", orgId));

        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.addOrder(Order.asc("id"));
        List<SysOrgGroup> ls = findAllByCriteria(dc);
        
		return ls;
	}

	@Override
	public boolean saveRelations(List<RelationOrgGroupUserVo> orgGroupUserVos) {
		for (Object obj: orgGroupUserVos) {
			RelationOrgGroupUserVo vo = JSON.parseObject(JSON.toJSONString(obj), RelationOrgGroupUserVo.class);
			if (vo.getSelecting() != vo.getSelected()) { // 变化了
				SysRelationOrgGroupUser orgGroupUserVo 
					= this.getRelationOrgGroupUser(vo.getOrgGroupId(), vo.getUserId());
				
    			if (vo.getSelecting() && orgGroupUserVo == null) { // 勾选
    				orgGroupUserVo = new SysRelationOrgGroupUser(vo.getOrgGroupId(), vo.getUserId());
    				saveOrUpdate(orgGroupUserVo);
    			} else if (orgGroupUserVo != null) { // 取消
    				getDao().delete(orgGroupUserVo);
    			}
			}
		}
		
		return true;
	}

	private SysRelationOrgGroupUser getRelationOrgGroupUser(Long orgGroupId, Long userId) {
		DetachedCriteria dc = DetachedCriteria.forClass(SysRelationOrgGroupUser.class);
        dc.add(Restrictions.eq("orgGroupId", orgGroupId));
        dc.add(Restrictions.eq("userId", userId));
        
        dc.addOrder(Order.asc("id"));
        List<SysRelationOrgGroupUser> ls = findAllByCriteria(dc);
        
        if (ls.size() == 0) {
        	return null;
        }
		return ls.get(0);
	}
	
	private RelationOrgGroupUserVo genVo(Long orgGroupId, Long userId) {
		SysOrgGroup orgGroup = (SysOrgGroup) get(SysOrgGroup.class, orgGroupId);
		SysUser user = (SysUser) get(SysUser.class, userId);
		
		RelationOrgGroupUserVo vo = new RelationOrgGroupUserVo();
		vo.setOrgGroupId(orgGroup.getId());
		vo.setOrgGroupName(orgGroup.getName());
		vo.setUserId(user.getId());
		vo.setUserName(user.getName());
		
		return vo;
	}
}
