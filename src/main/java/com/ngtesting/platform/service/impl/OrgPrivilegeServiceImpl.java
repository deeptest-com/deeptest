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
import com.ngtesting.platform.entity.SysOrgGroup;
import com.ngtesting.platform.entity.SysOrgPrivilege;
import com.ngtesting.platform.entity.SysOrgRole;
import com.ngtesting.platform.entity.SysRelationOrgGroupUser;
import com.ngtesting.platform.entity.SysUser;
import com.ngtesting.platform.service.OrgGroupService;
import com.ngtesting.platform.service.OrgPrivilegeService;
import com.ngtesting.platform.service.RelationOrgGroupUserService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.util.StringUtil;
import com.ngtesting.platform.vo.OrgGroupVo;
import com.ngtesting.platform.vo.OrgPrivilegeVo;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.RelationOrgGroupUserVo;

@Service
public class OrgPrivilegeServiceImpl extends BaseServiceImpl implements OrgPrivilegeService {

	@Override
	public List<OrgPrivilegeVo> listPrivilegesByOrg(Long orgId, Long orgRoleId) {
		
        List<SysOrgPrivilege> allPrivileges = listAllOrgPrivileges();
        
        List<SysOrgPrivilege> orgRolePrivileges;
        if (orgRoleId == null) {
        	orgRolePrivileges = new LinkedList<SysOrgPrivilege>();
        } else {
        	orgRolePrivileges = listOrgRolePrivileges(orgId, orgRoleId);
        }
        
        List<OrgPrivilegeVo> vos = new LinkedList<OrgPrivilegeVo>();
        for (SysOrgPrivilege po1 : allPrivileges) {
        	OrgPrivilegeVo vo = genVo(orgId, po1);
        	
        	vo.setSelected(false);
        	vo.setSelecting(false);
        	for (SysOrgPrivilege po2 : orgRolePrivileges) {
        		if (po1.getId() == po2.getId()) {
            		vo.setSelected(true);
            		vo.setSelecting(true);
            	}
        	}
        	vos.add(vo);
        }
        
		return vos;
	}

	private OrgPrivilegeVo genVo(Long orgId, SysOrgPrivilege po1) {
		OrgPrivilegeVo vo = new OrgPrivilegeVo(po1.getId(), po1.getName(), po1.getDescr(), orgId);
		
		return vo;
	}

	private List<SysOrgPrivilege> listOrgRolePrivileges(Long orgId, Long orgRoleId) {
		
		DetachedCriteria dc = DetachedCriteria.forClass(SysOrgPrivilege.class);
		
        dc.createAlias("orgRoleSet", "roles");
        dc.add(Restrictions.eq("roles.id", orgRoleId));
        
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        
        dc.addOrder(Order.asc("id"));
        List ls = findAllByCriteria(dc);
		
		return ls;
	}

	private List<SysOrgPrivilege> listAllOrgPrivileges() {
		DetachedCriteria dc = DetachedCriteria.forClass(SysOrgPrivilege.class);

        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.addOrder(Order.asc("id"));
        List<SysOrgPrivilege> ls = findAllByCriteria(dc);
        
		return ls;
	}

	@Override
	public boolean saveOrgPrivileges(Long roleId, List<OrgPrivilegeVo> orgPrivileges) {
		if (orgPrivileges == null) {
			return false;
		}
		
		SysOrgRole orgRole = (SysOrgRole) get(SysOrgRole.class, roleId);
		Set<SysOrgPrivilege> privilegeSet = orgRole.getOrgPrivilegeSet();
		
		for (Object obj: orgPrivileges) {
			OrgPrivilegeVo vo = JSON.parseObject(JSON.toJSONString(obj), OrgPrivilegeVo.class);
			if (vo.getSelecting() != vo.getSelected()) { // 变化了
				SysOrgPrivilege orgPrivilege = (SysOrgPrivilege) get(SysOrgPrivilege.class, vo.getId());
				
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
