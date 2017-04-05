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
import com.ngtesting.platform.entity.SysOrgPriviledge;
import com.ngtesting.platform.entity.SysOrgRole;
import com.ngtesting.platform.entity.SysRelationOrgGroupUser;
import com.ngtesting.platform.entity.SysUser;
import com.ngtesting.platform.service.OrgGroupService;
import com.ngtesting.platform.service.OrgPriviledgeService;
import com.ngtesting.platform.service.RelationOrgGroupUserService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.util.StringUtil;
import com.ngtesting.platform.vo.OrgGroupVo;
import com.ngtesting.platform.vo.OrgPriviledgeVo;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.RelationOrgGroupUserVo;

@Service
public class OrgPriviledgeServiceImpl extends BaseServiceImpl implements OrgPriviledgeService {

	@Override
	public List<OrgPriviledgeVo> listPriviledgesByOrg(Long orgId, Long orgRoleId) {
		
        List<SysOrgPriviledge> allPriviledges = listAllOrgPriviledges();
        
        List<SysOrgPriviledge> orgRolePriviledges;
        if (orgRoleId == null) {
        	orgRolePriviledges = new LinkedList<SysOrgPriviledge>();
        } else {
        	orgRolePriviledges = listOrgRolePriviledges(orgId, orgRoleId);
        }
        
        List<OrgPriviledgeVo> vos = new LinkedList<OrgPriviledgeVo>();
        for (SysOrgPriviledge po1 : allPriviledges) {
        	OrgPriviledgeVo vo = genVo(orgId, po1);
        	
        	vo.setSelected(false);
        	vo.setSelecting(false);
        	for (SysOrgPriviledge po2 : orgRolePriviledges) {
        		if (po1.getId() == po2.getId()) {
            		vo.setSelected(true);
            		vo.setSelecting(true);
            	}
        	}
        	vos.add(vo);
        }
        
		return vos;
	}

	private OrgPriviledgeVo genVo(Long orgId, SysOrgPriviledge po1) {
		OrgPriviledgeVo vo = new OrgPriviledgeVo(po1.getId(), po1.getName(), po1.getDescr(), orgId);
		
		return vo;
	}

	private List<SysOrgPriviledge> listOrgRolePriviledges(Long orgId, Long orgRoleId) {
		
		DetachedCriteria dc = DetachedCriteria.forClass(SysOrgPriviledge.class);
		
        dc.createAlias("orgRoleSet", "roles");
        dc.add(Restrictions.eq("roles.id", orgRoleId));
        
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        
        dc.addOrder(Order.asc("id"));
        List ls = findAllByCriteria(dc);
		
		return ls;
	}

	private List<SysOrgPriviledge> listAllOrgPriviledges() {
		DetachedCriteria dc = DetachedCriteria.forClass(SysOrgPriviledge.class);

        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.addOrder(Order.asc("id"));
        List<SysOrgPriviledge> ls = findAllByCriteria(dc);
        
		return ls;
	}

	@Override
	public boolean saveOrgPriviledges(Long roleId, List<OrgPriviledgeVo> orgPriviledges) {
		if (orgPriviledges == null) {
			return false;
		}
		
		SysOrgRole orgRole = (SysOrgRole) get(SysOrgRole.class, roleId);
		Set<SysOrgPriviledge> priviledgeSet = orgRole.getOrgPriviledgeSet();
		
		for (Object obj: orgPriviledges) {
			OrgPriviledgeVo vo = JSON.parseObject(JSON.toJSONString(obj), OrgPriviledgeVo.class);
			if (vo.getSelecting() != vo.getSelected()) { // 变化了
				SysOrgPriviledge orgPriviledge = (SysOrgPriviledge) get(SysOrgPriviledge.class, vo.getId());
				
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
