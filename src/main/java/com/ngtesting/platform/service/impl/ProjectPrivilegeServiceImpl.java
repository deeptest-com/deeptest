package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.ngtesting.platform.entity.TestProjectPrivilege;
import com.ngtesting.platform.entity.TestProjectRole;
import com.ngtesting.platform.service.ProjectPrivilegeService;
import com.ngtesting.platform.vo.ProjectPrivilegeVo;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import java.util.*;

@Service
public class ProjectPrivilegeServiceImpl extends BaseServiceImpl implements ProjectPrivilegeService {

	@Override
	public Map<String, List<ProjectPrivilegeVo>> listPrivilegesByOrg(Long orgId, Long projectRoleId) {
		
        List<TestProjectPrivilege> allPrivileges = listAllProjectPrivileges();
        
        List<TestProjectPrivilege> projectRolePrivileges;
        if (projectRoleId == null) {
        	projectRolePrivileges = new LinkedList<TestProjectPrivilege>();
        } else {
        	projectRolePrivileges = listProjectRolePrivileges(orgId, projectRoleId);
        }
        
        Map<String, List<ProjectPrivilegeVo>> map = new LinkedHashMap<String, List<ProjectPrivilegeVo>>();
        for (TestProjectPrivilege po1 : allPrivileges) {
        	String key = po1.getName();
        	if (!map.containsKey(key)) {
        		List<ProjectPrivilegeVo> vos = new LinkedList<ProjectPrivilegeVo>();
        		map.put(key, vos);
        	}
        	
        	ProjectPrivilegeVo vo = genVo(orgId, po1);
        	
        	vo.setSelected(false);
        	vo.setSelecting(false);
        	for (TestProjectPrivilege po2 : projectRolePrivileges) {
        		if (po1.getId() == po2.getId()) {
            		vo.setSelected(true);
            		vo.setSelecting(true);
            	}
        	}
        	map.get(key).add(vo);
        }
        
		return map;
	}

	private ProjectPrivilegeVo genVo(Long orgId, TestProjectPrivilege po1) {
		ProjectPrivilegeVo vo = new ProjectPrivilegeVo(po1.getId(), po1.getCode().toString(), po1.getAction().toString(), 
				po1.getName(), po1.getDescr(), orgId);
		
		return vo;
	}

	private List<TestProjectPrivilege> listProjectRolePrivileges(Long orgId, Long projectRoleId) {
		
		DetachedCriteria dc = DetachedCriteria.forClass(TestProjectPrivilege.class);
		
        dc.createAlias("projectRoleSet", "roles");
        dc.add(Restrictions.eq("roles.id", projectRoleId));
        
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        
        dc.addOrder(Order.asc("id"));
        List ls = findAllByCriteria(dc);
		
		return ls;
	}

	private List<TestProjectPrivilege> listAllProjectPrivileges() {
		DetachedCriteria dc = DetachedCriteria.forClass(TestProjectPrivilege.class);

        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.addOrder(Order.asc("id"));
        List<TestProjectPrivilege> ls = findAllByCriteria(dc);
        
		return ls;
	}

	@Override
	public boolean saveProjectPrivileges(Long roleId, Map<String, List<ProjectPrivilegeVo>> map) {
		if (map == null) {
			return false;
		}
		
		TestProjectRole orgRole = (TestProjectRole) get(TestProjectRole.class, roleId);
		Set<TestProjectPrivilege> privilegeSet = orgRole.getProjectPrivilegeSet();
		
		for (String key: map.keySet()) {
			List<ProjectPrivilegeVo> ls = JSON.parseObject(JSON.toJSONString(map.get(key)), List.class);
			
			for (Object obj: ls) {
				ProjectPrivilegeVo vo = JSON.parseObject(JSON.toJSONString(obj), ProjectPrivilegeVo.class);
				if (vo.getSelecting() != vo.getSelected()) { // 变化了
					TestProjectPrivilege orgPrivilege = (TestProjectPrivilege) get(TestProjectPrivilege.class, vo.getId());
					
	    			if (vo.getSelecting() && !privilegeSet.contains(orgPrivilege)) { // 勾选
	    				privilegeSet.add(orgPrivilege);
	    			} else if (orgPrivilege != null) { // 取消
	    				privilegeSet.remove(orgPrivilege);
	    			}
				}
			}
		}
		
		saveOrUpdate(orgRole);
		
		return true;
	}

	@Override
	public Map<String, Map<String, Boolean>> listByUser(Long userId) {
		return null;
	}

}
