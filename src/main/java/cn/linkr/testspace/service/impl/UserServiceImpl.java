package cn.linkr.testspace.service.impl;

import java.util.Date;
import java.util.LinkedList;
import java.util.List;
import java.util.UUID;

import org.apache.commons.lang.StringUtils;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import cn.linkr.testspace.entity.EvtGuest;
import cn.linkr.testspace.entity.SysCompany;
import cn.linkr.testspace.entity.SysUser;
import cn.linkr.testspace.entity.SysUser.AgentType;
import cn.linkr.testspace.entity.SysVerifyCode;
import cn.linkr.testspace.service.UserService;
import cn.linkr.testspace.util.BeanUtilEx;
import cn.linkr.testspace.util.DateUtils;
import cn.linkr.testspace.util.StringUtil;
import cn.linkr.testspace.vo.CompanyVo;
import cn.linkr.testspace.vo.GuestVo;
import cn.linkr.testspace.vo.Page;
import cn.linkr.testspace.vo.UserVo;

@Service
public class UserServiceImpl extends BaseServiceImpl implements UserService {

	@Override
	public Page listByPage(Long companyId, String keywords, String disabled, Integer currentPage, Integer itemsPerPage) {
        DetachedCriteria dc = DetachedCriteria.forClass(SysUser.class);
        dc.add(Restrictions.eq("companyId", companyId));
        
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        
        if (StringUtil.isNotEmpty(disabled)) {
        	dc.add(Restrictions.eq("disabled", Boolean.valueOf(disabled)));
        }
        
        dc.addOrder(Order.asc("id"));
        Page page = findPage(dc, currentPage * itemsPerPage, itemsPerPage);
		
		return page;
	}

	@Override
	public SysUser save(UserVo vo) {
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
		
		saveOrUpdate(po);
		return po;
	}

	@Override
	public boolean remove(Long id) {
		SysUser po = (SysUser) get(SysUser.class, id);
		po.setDeleted(true);
		saveOrUpdate(po);
		
		return true;
	}

	@Override
	public boolean disablePers(Long id) {
		SysUser po = (SysUser) get(SysUser.class, id);
		po.setDisabled(!po.getDisabled());
		saveOrUpdate(po);
		
		return true;
	}
    
	@Override
	public UserVo genVo(SysUser user) {
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
