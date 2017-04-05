package com.ngtesting.platform.service.impl;

import java.util.HashMap;
import java.util.LinkedHashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;
import java.util.Set;

import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.alibaba.fastjson.JSON;
import com.ngtesting.platform.entity.SysProjectPrivilege;
import com.ngtesting.platform.entity.SysProjectRole;
import com.ngtesting.platform.entity.SysProjectPrivilege;
import com.ngtesting.platform.entity.SysRelationOrgGroupUser;
import com.ngtesting.platform.entity.SysUser;
import com.ngtesting.platform.service.OrgGroupService;
import com.ngtesting.platform.service.OrgPrivilegeService;
import com.ngtesting.platform.service.ProjectPrivilegeService;
import com.ngtesting.platform.service.RelationOrgGroupUserService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.util.StringUtil;
import com.ngtesting.platform.vo.OrgGroupVo;
import com.ngtesting.platform.vo.OrgPrivilegeVo;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.ProjectPrivilegeVo;
import com.ngtesting.platform.vo.RelationOrgGroupUserVo;

@Service
public class ProjectPrivilegeServiceImpl extends BaseServiceImpl implements ProjectPrivilegeService {

	@Override
	public Map<String, List<ProjectPrivilegeVo>> listPrivilegesByOrg(Long orgId, Long projectRoleId) {
		
        List<SysProjectPrivilege> allPrivileges = listAllProjectPrivileges();
        
        List<SysProjectPrivilege> projectRolePrivileges;
        if (projectRoleId == null) {
        	projectRolePrivileges = new LinkedList<SysProjectPrivilege>();
        } else {
        	projectRolePrivileges = listProjectRolePrivileges(orgId, projectRoleId);
        }
        
        Map<String, List<ProjectPrivilegeVo>> map = new LinkedHashMap<String, List<ProjectPrivilegeVo>>();
        for (SysProjectPrivilege po1 : allPrivileges) {
        	String key = po1.getName();
        	if (!map.containsKey(key)) {
        		List<ProjectPrivilegeVo> vos = new LinkedList<ProjectPrivilegeVo>();
        		map.put(key, vos);
        	}
        	
        	ProjectPrivilegeVo vo = genVo(orgId, po1);
        	
        	vo.setSelected(false);
        	vo.setSelecting(false);
        	for (SysProjectPrivilege po2 : projectRolePrivileges) {
        		if (po1.getId() == po2.getId()) {
            		vo.setSelected(true);
            		vo.setSelecting(true);
            	}
        	}
        	map.get(key).add(vo);
        }
        
		return map;
	}

	private ProjectPrivilegeVo genVo(Long orgId, SysProjectPrivilege po1) {
		ProjectPrivilegeVo vo = new ProjectPrivilegeVo(po1.getId(), po1.getCode().toString(), po1.getAction().toString(), 
				po1.getName(), po1.getDescr(), orgId);
		
		return vo;
	}

	private List<SysProjectPrivilege> listProjectRolePrivileges(Long orgId, Long projectRoleId) {
		
		DetachedCriteria dc = DetachedCriteria.forClass(SysProjectPrivilege.class);
		
        dc.createAlias("projectRoleSet", "roles");
        dc.add(Restrictions.eq("roles.id", projectRoleId));
        
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        
        dc.addOrder(Order.asc("id"));
        List ls = findAllByCriteria(dc);
		
		return ls;
	}

	private List<SysProjectPrivilege> listAllProjectPrivileges() {
		DetachedCriteria dc = DetachedCriteria.forClass(SysProjectPrivilege.class);

        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.addOrder(Order.asc("id"));
        List<SysProjectPrivilege> ls = findAllByCriteria(dc);
        
		return ls;
	}

	@Override
	public boolean saveProjectPrivileges(Long roleId, List<ProjectPrivilegeVo> projectPrivileges) {
		if (projectPrivileges == null) {
			return false;
		}
		
		SysProjectRole orgRole = (SysProjectRole) get(SysProjectRole.class, roleId);
		Set<SysProjectPrivilege> privilegeSet = orgRole.getProjectPrivilegeSet();
		
		for (Object obj: projectPrivileges) {
			ProjectPrivilegeVo vo = JSON.parseObject(JSON.toJSONString(obj), ProjectPrivilegeVo.class);
			if (vo.getSelecting() != vo.getSelected()) { // 变化了
				SysProjectPrivilege orgPrivilege = (SysProjectPrivilege) get(SysProjectPrivilege.class, vo.getId());
				
    			if (vo.getSelecting() && !privilegeSet.contains(orgPrivilege)) { // 勾选
    				privilegeSet.add(orgPrivilege);
    			} else if (orgPrivilege != null) { // 取消
    				privilegeSet.remove(orgPrivilege);
    			}
			}
		}
		saveOrUpdate(orgRole);
		
		return true;
	}
	
}
