package com.ngtesting.platform.service.impl;

import java.util.LinkedList;
import java.util.List;

import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.ngtesting.platform.entity.SysOrg;
import com.ngtesting.platform.entity.SysUser;
import com.ngtesting.platform.entity.TestProject;
import com.ngtesting.platform.service.OrgService;
import com.ngtesting.platform.service.TestProjectService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.util.StringUtil;
import com.ngtesting.platform.vo.OrgVo;
import com.ngtesting.platform.vo.TestProjectAccessHistoryVo;
import com.ngtesting.platform.vo.TestProjectVo;
import com.ngtesting.platform.vo.UserVo;

@Service
public class OrgServiceImpl extends BaseServiceImpl implements OrgService {
	
	@Autowired
	TestProjectService projectService;

	@Override
	public List<SysOrg> list(String keywords, String disabled, Long userId) {
        DetachedCriteria dc = DetachedCriteria.forClass(SysOrg.class);
        dc.createAlias("userSet", "users");
        dc.add(Restrictions.eq("users.id", userId));
        
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        
        if (StringUtil.isNotEmpty(keywords)) {
			dc.add(Restrictions.like("name", "%" + keywords + "%"));
		}
        if (StringUtil.isNotEmpty(disabled)) {
			dc.add(Restrictions.eq("disabled", Boolean.valueOf(disabled)));
		}
        
        dc.addOrder(Order.asc("id"));
        List<SysOrg> ls = findAllByCriteria(dc);

		return ls;
	}
	
	@Override
	public List<OrgVo> listVo(String keywords, String disabled, Long id) {
		List ls = list(keywords, disabled, id);
		List<OrgVo> vos = genVos(ls, id);
		return vos;
	}

	@Override
	public SysOrg getDetail(Long id) {
		if (id == null) {
			return null;
		}
		SysOrg po = (SysOrg) get(SysOrg.class, id);

		return po;
	}

	@Override
	public SysOrg save(OrgVo vo, Long userId) {
		if (vo == null) {
			return null;
		}

		SysUser user = (SysUser)get(SysUser.class, userId);
		
		boolean isNew = vo.getId() == null;
		SysOrg po = new SysOrg();
		if (!isNew) {
			po = (SysOrg) get(SysOrg.class, vo.getId());
		} else {
			po.setAdminId(userId);
			po.getUserSet().add(user);
		}
		
		po.setName(vo.getName());
		po.setWebsite(vo.getWebsite());
		po.setDisabled(vo.getDisabled());

		saveOrUpdate(po);
		
		if (user.getDefaultOrgId() == null) {
			user.setDefaultOrgId(po.getId());
			saveOrUpdate(user);
		}
		
		return po;
	}

	@Override
	public Boolean disable(Long id) {
		if (id == null) {
			return false;
		}

		SysOrg po = (SysOrg) get(SysOrg.class, id);
		po.setDisabled(true);
		saveOrUpdate(po);

		return true;
	}
	
	@Override
	public Boolean delete(Long id) {
		if (id == null) {
			return false;
		}

		SysOrg po = (SysOrg) get(SysOrg.class, id);
		po.setDeleted(true);
		saveOrUpdate(po);

		return true;
	}
	
	@Override
	public List<TestProjectAccessHistoryVo> setDefaultPers(Long orgId, UserVo userVo) {
		SysUser user = (SysUser) get(SysUser.class, userVo.getId());
		
		user.setDefaultOrgId(orgId);
		
		List<TestProjectAccessHistoryVo> recentProjects = projectService.listRecentProjectVo(orgId, userVo.getId());
		if (recentProjects.size() > 0) {
			user.setDefaultProjectId(recentProjects.get(0).getId());
		}
		
		saveOrUpdate(user);
		
		userVo.setDefaultOrgId(user.getDefaultOrgId());
		userVo.setDefaultProjectId(user.getDefaultProjectId());
		
		return recentProjects;
	}

	@Override
	public List<OrgVo> genVos(List<SysOrg> pos, Long userId) {
		SysUser user = (SysUser)get(SysUser.class, userId);
		
		List<OrgVo> voList = new LinkedList<OrgVo>();
		for (SysOrg po : pos) {
			OrgVo vo = genVo(po);
			if (po.getId() == user.getDefaultOrgId()) {
				vo.setDefaultOrg(true);
			}
			
			voList.add(vo);
		}
		
		return voList;
	}

	@Override
	public OrgVo genVo(SysOrg po) {
		if (po == null) {
			return null;
		}
		OrgVo vo = new OrgVo();
		BeanUtilEx.copyProperties(vo, po);

		return vo;
	}

}
