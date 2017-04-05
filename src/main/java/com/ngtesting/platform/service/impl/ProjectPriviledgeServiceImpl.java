package com.ngtesting.platform.service.impl;

import java.util.LinkedList;
import java.util.List;
import java.util.Set;

import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.alibaba.fastjson.JSON;
import com.ngtesting.platform.entity.SysProjectPriviledge;
import com.ngtesting.platform.entity.SysProjectRole;
import com.ngtesting.platform.entity.SysProjectPriviledge;
import com.ngtesting.platform.entity.SysRelationOrgGroupUser;
import com.ngtesting.platform.entity.SysUser;
import com.ngtesting.platform.service.OrgGroupService;
import com.ngtesting.platform.service.OrgPriviledgeService;
import com.ngtesting.platform.service.ProjectPriviledgeService;
import com.ngtesting.platform.service.RelationOrgGroupUserService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.util.StringUtil;
import com.ngtesting.platform.vo.OrgGroupVo;
import com.ngtesting.platform.vo.OrgPriviledgeVo;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.ProjectPriviledgeVo;
import com.ngtesting.platform.vo.RelationOrgGroupUserVo;

@Service
public class ProjectPriviledgeServiceImpl extends BaseServiceImpl implements ProjectPriviledgeService {

	@Override
	public List<ProjectPriviledgeVo> listPriviledgesByOrg(Long orgId, Long projectRoleId) {
		
        List<SysProjectPriviledge> allPriviledges = listAllProjectPriviledges();
        
        List<SysProjectPriviledge> projectRolePriviledges;
        if (projectRoleId == null) {
        	projectRolePriviledges = new LinkedList<SysProjectPriviledge>();
        } else {
        	projectRolePriviledges = listProjectRolePriviledges(orgId, projectRoleId);
        }
        
        List<ProjectPriviledgeVo> vos = new LinkedList<ProjectPriviledgeVo>();
        for (SysProjectPriviledge po1 : allPriviledges) {
        	ProjectPriviledgeVo vo = genVo(orgId, po1);
        	
        	vo.setSelected(false);
        	vo.setSelecting(false);
        	for (SysProjectPriviledge po2 : projectRolePriviledges) {
        		if (po1.getId() == po2.getId()) {
            		vo.setSelected(true);
            		vo.setSelecting(true);
            	}
        	}
        	vos.add(vo);
        }
        
		return vos;
	}

	private ProjectPriviledgeVo genVo(Long orgId, SysProjectPriviledge po1) {
		ProjectPriviledgeVo vo = new ProjectPriviledgeVo(po1.getId(), po1.getName(), po1.getDescr(), orgId);
		
		return vo;
	}

	private List<SysProjectPriviledge> listProjectRolePriviledges(Long orgId, Long projectRoleId) {
		
		DetachedCriteria dc = DetachedCriteria.forClass(SysProjectPriviledge.class);
		
        dc.createAlias("projectRoleSet", "roles");
        dc.add(Restrictions.eq("roles.id", projectRoleId));
        
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        
        dc.addOrder(Order.asc("id"));
        List ls = findAllByCriteria(dc);
		
		return ls;
	}

	private List<SysProjectPriviledge> listAllProjectPriviledges() {
		DetachedCriteria dc = DetachedCriteria.forClass(SysProjectPriviledge.class);

        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.addOrder(Order.asc("id"));
        List<SysProjectPriviledge> ls = findAllByCriteria(dc);
        
		return ls;
	}

	@Override
	public boolean saveProjectPriviledges(Long roleId, List<ProjectPriviledgeVo> projectPriviledges) {
		if (projectPriviledges == null) {
			return false;
		}
		
		SysProjectRole orgRole = (SysProjectRole) get(SysProjectRole.class, roleId);
		Set<SysProjectPriviledge> priviledgeSet = orgRole.getProjectPriviledgeSet();
		
		for (Object obj: projectPriviledges) {
			ProjectPriviledgeVo vo = JSON.parseObject(JSON.toJSONString(obj), ProjectPriviledgeVo.class);
			if (vo.getSelecting() != vo.getSelected()) { // 变化了
				SysProjectPriviledge orgPriviledge = (SysProjectPriviledge) get(SysProjectPriviledge.class, vo.getId());
				
    			if (vo.getSelecting() && !priviledgeSet.contains(orgPriviledge)) { // 勾选
    				priviledgeSet.add(orgPriviledge);
    			} else if (orgPriviledge != null) { // 取消
    				priviledgeSet.remove(orgPriviledge);
    			}
			}
		}
		saveOrUpdate(orgRole);
		
		return true;
	}
	
}
