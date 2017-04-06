package com.ngtesting.platform.service.impl;

import java.util.LinkedList;
import java.util.List;

import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.ngtesting.platform.entity.SysCustomField;
import com.ngtesting.platform.entity.SysOrg;
import com.ngtesting.platform.entity.SysUser;
import com.ngtesting.platform.service.AccountService;
import com.ngtesting.platform.service.CustomFieldService;
import com.ngtesting.platform.service.RelationOrgGroupUserService;
import com.ngtesting.platform.service.RelationProjectRoleUserService;
import com.ngtesting.platform.service.UserService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.util.StringUtil;
import com.ngtesting.platform.vo.CustomFieldVo;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.UserVo;

@Service
public class CustomFieldServiceImpl extends BaseServiceImpl implements CustomFieldService {
	
	@Autowired
	AccountService accountService;

	@Override
	public Page listByPage(Long orgId, String keywords, String disabled, Integer currentPage, Integer itemsPerPage) {
        DetachedCriteria dc = DetachedCriteria.forClass(SysCustomField.class);
        
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
	public SysCustomField save(CustomFieldVo userVo, Long orgId) {
		if (userVo == null) {
			return null;
		}
		
		SysCustomField po;
		if (userVo.getId() != null) {
			po = (SysCustomField) get(SysUser.class, userVo.getId());
		} else {
			po = new SysCustomField();
		}
		
		po.setName(userVo.getName());
		po.setDisabled(userVo.getDisabled());
		
		SysOrg org = (SysOrg)get(SysOrg.class, orgId);
		if (!contains(org.getUserSet(), userVo.getId())) {
//			org.getUserSet().add(po);
			saveOrUpdate(org);
		}
		
		saveOrUpdate(po);
		return po;
	}

	@Override
	public boolean delete(Long id) {
		SysCustomField po = (SysCustomField) get(SysCustomField.class, id);
		po.setDeleted(true);
		saveOrUpdate(po);
		
		return true;
	}
    
	@Override
	public CustomFieldVo genVo(SysCustomField user) {
		if (user == null) {
			return null;
		}
		CustomFieldVo vo = new CustomFieldVo();
		BeanUtilEx.copyProperties(vo, user);
		
		return vo;
	}
	@Override
	public List<CustomFieldVo> genVos(List<SysCustomField> pos) {
        List<CustomFieldVo> vos = new LinkedList<CustomFieldVo>();

        for (SysCustomField po: pos) {
        	CustomFieldVo vo = genVo(po);
        	vos.add(vo);
        }
		return vos;
	}

}
