package cn.linkr.events.service.impl;

import java.util.Date;
import java.util.LinkedList;
import java.util.List;
import java.util.UUID;

import org.apache.commons.lang.StringUtils;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import cn.linkr.events.entity.EvtGuest;
import cn.linkr.events.entity.SysCompany;
import cn.linkr.events.entity.SysUser;
import cn.linkr.events.entity.SysUser.AgentType;
import cn.linkr.events.entity.SysVerifyCode;
import cn.linkr.events.service.UserService;
import cn.linkr.events.util.BeanUtilEx;
import cn.linkr.events.util.DateUtils;
import cn.linkr.events.util.StringUtil;
import cn.linkr.events.vo.CompanyVo;
import cn.linkr.events.vo.GuestVo;
import cn.linkr.events.vo.Page;
import cn.linkr.events.vo.UserVo;

@Service
public class UserServiceImpl extends BaseServiceImpl implements UserService {

	@Override
	public SysUser loginPers(String email, String password, Boolean rememberMe) {
		String newToken = null;
		DetachedCriteria dc = DetachedCriteria.forClass(SysUser.class);
		dc.add(Restrictions.eq("email", email));
		dc.add(Restrictions.eq("password", password));
		dc.add(Restrictions.ne("deleted", true));
		dc.add(Restrictions.ne("disabled", true));
		List<SysUser> ls = (List<SysUser>) findAllByCriteria(dc);

		SysUser user = null;
		if (ls.size() > 0) {
			user = ls.get(0);
			newToken = UUID.randomUUID().toString();
			user.setToken(newToken);

			user.setLastLoginTime(new Date());
			saveOrUpdate(user);
		}
		return user;
	}

	@Override
	public SysUser registerPers(String name, String email, String phone, String password) {
		String newToken = null;
		DetachedCriteria dc = DetachedCriteria.forClass(SysUser.class);
		
		dc.add(Restrictions.eq("email", email));
		dc.add(Restrictions.ne("deleted", true));
		dc.add(Restrictions.ne("disabled", true));
		List<SysUser> ls = (List<SysUser>) findAllByCriteria(dc);

		if (ls.size() > 0) {
			return null;
		}

		SysUser user = new SysUser();
		newToken = UUID.randomUUID().toString();
		user.setName(name);
		user.setToken(newToken);
		user.setEmail(email);
		user.setPhone(phone);
		user.setPassword(password);
		
		user.setLastLoginTime(new Date());
		saveOrUpdate(user);

		return user;
	}

	@Override
	public SysVerifyCode forgetPaswordPers(String phone) {
		SysUser user = getByPhone(phone);
		if (user == null) {
			return null;
		}

		SysVerifyCode po = new SysVerifyCode();
		String code = StringUtil.RandomNumbString(4);
		Date now = new Date();
		po.setUserId(user.getId());
		po.setCode(code);
		po.setCreateTime(now);
		po.setExpireTime(new Date(now.getTime() + 300000));
		saveOrUpdate(po);

		return po;
	}

	@Override
	public SysUser getByPhone(String phone) {
		DetachedCriteria dc = DetachedCriteria.forClass(SysUser.class);
		dc.add(Restrictions.eq("phone", phone));
		dc.add(Restrictions.ne("deleted", true));
		dc.add(Restrictions.ne("disabled", true));

		List ls = findAllByCriteria(dc);
		if (ls.size() > 0) {
			return (SysUser) ls.get(0);
		} else {
			return null;
		}
	}

	@Override
	public SysUser resetPasswordPers(String verifyCode, String phone,
			String password, String platform, String isWebview,
			String deviceToken) {

		SysUser user = getByPhone(phone);
		if (user == null) {
			return null;
		}

		String newToken = null;
		DetachedCriteria dc = DetachedCriteria.forClass(SysVerifyCode.class);
		dc.add(Restrictions.eq("userId", user.getId()));
		dc.add(Restrictions.eq("code", verifyCode));
		dc.add(Restrictions.ne("deleted", true));
		dc.add(Restrictions.ne("disabled", true));
		List<SysVerifyCode> ls = (List<SysVerifyCode>) findAllByCriteria(dc);

		if (ls.size() < 1) {
			return null;
		}

		newToken = UUID.randomUUID().toString();
		user.setToken(newToken);
		user.setPhone(phone);
		user.setPassword(password);

		if (StringUtils.isNotEmpty(platform)) {
			user.setPlatform(SysUser.PlatformType.valueOf(platform.trim().toUpperCase()));
		}

		if (StringUtils.isNotEmpty(isWebview)) {
			AgentType agent = Boolean.valueOf(isWebview)? AgentType.WEBVIEW: AgentType.BROWSER;
			user.setAgent(agent);
		}

		if (StringUtils.isNotEmpty(deviceToken)) {
			user.setToken(deviceToken);
		}
		user.setLastLoginTime(new Date());
		saveOrUpdate(user);

		SysVerifyCode code = ls.get(0);
		code.setDeleted(true);
		saveOrUpdate(code);

		return user;
	}

	@Override
	public SysUser logoutPers(String email) {
		
		DetachedCriteria dc = DetachedCriteria.forClass(SysUser.class);
		dc.add(Restrictions.eq("email", email));
		dc.add(Restrictions.ne("deleted", true));
		dc.add(Restrictions.ne("disabled", true));
		List<SysUser> ls = (List<SysUser>) findAllByCriteria(dc);

		SysUser user = null;
		if (ls.size() > 0) {
			user = ls.get(0);
			user.setToken("");
			saveOrUpdate(user);
		}
		return user;
	}
	

	@Override
	public SysUser getByToken(String token) {
		DetachedCriteria dc = DetachedCriteria.forClass(SysUser.class);
		dc.add(Restrictions.eq("token", token));
		
		dc.add(Restrictions.ne("deleted", true));
		dc.add(Restrictions.ne("disabled", true));

		List ls = findAllByCriteria(dc);
		if (ls.size() > 0) {
			return (SysUser) ls.get(0);
		} else {
			return null;
		}
	}

	@Override
	public UserVo genVo(SysUser user) {
		UserVo vo = new UserVo();
		BeanUtilEx.copyProperties(vo, user);
		
		return vo;
	}

	@Override
	public Page listByPage(long companyId, int currentPage, int itemsPerPage) {
        DetachedCriteria dc = DetachedCriteria.forClass(SysUser.class);
        dc.add(Restrictions.eq("companyId", companyId));
        
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
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
	public boolean disable(Long id) {
		SysUser po = (SysUser) get(SysUser.class, id);
		po.setDisabled(true);
		saveOrUpdate(po);
		
		return true;
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

	@Override
	public SysUser saveProfile(UserVo vo) {
		SysUser po = (SysUser) get(SysUser.class, vo.getId());

		String name = vo.getName();
		String email = vo.getEmail(); 
		String phone = vo.getPhone(); 

		po.setPhone(phone);
		po.setName(name);
		po.setEmail(email);
		saveOrUpdate(po);
		return po;
	}
    
}
