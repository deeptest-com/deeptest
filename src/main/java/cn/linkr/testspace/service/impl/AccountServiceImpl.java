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
import cn.linkr.testspace.service.AccountService;
import cn.linkr.testspace.service.UserService;
import cn.linkr.testspace.util.BeanUtilEx;
import cn.linkr.testspace.util.DateUtils;
import cn.linkr.testspace.util.StringUtil;
import cn.linkr.testspace.vo.CompanyVo;
import cn.linkr.testspace.vo.GuestVo;
import cn.linkr.testspace.vo.Page;
import cn.linkr.testspace.vo.UserVo;

@Service
public class AccountServiceImpl extends BaseServiceImpl implements AccountService {

	@Override
	public SysUser loginPers(String email, String password, Boolean rememberMe) {
		String newToken = null;
		DetachedCriteria dc = DetachedCriteria.forClass(SysUser.class);
		dc.add(Restrictions.eq("email", email));
		dc.add(Restrictions.eq("password", password));
		dc.add(Restrictions.ne("deleted", true));
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
	public SysVerifyCode forgotPasswordPers(Long userId) {
		SysUser user = (SysUser) get(SysUser.class, userId);
		if (user == null) {
			return null;
		}

		SysVerifyCode po = new SysVerifyCode();
		String code = UUID.randomUUID().toString().replaceAll("-", "");
		Date now = new Date();
		po.setRefId(user.getId());
		po.setCode(code);
		po.setCreateTime(now);
		po.setExpireTime(new Date(now.getTime() + 60 * 60 * 1000));
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
	public SysUser getByEmail(String email) {
		DetachedCriteria dc = DetachedCriteria.forClass(SysUser.class);
		dc.add(Restrictions.eq("email", email));
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
	public SysUser resetPasswordPers(String verifyCode, String password) {

		String newToken = null;
		DetachedCriteria dc = DetachedCriteria.forClass(SysVerifyCode.class);
		dc.add(Restrictions.eq("code", verifyCode));
		dc.add(Restrictions.ge("expireTime", new Date()));
		dc.add(Restrictions.ne("deleted", true));
		dc.add(Restrictions.ne("disabled", true));
		dc.addOrder(Order.desc("id"));
		List<SysVerifyCode> ls = (List<SysVerifyCode>) findAllByCriteria(dc);

		if (ls.size() < 1) {
			return null;
		}
		
		SysVerifyCode code = ls.get(0);
		code.setDeleted(true);
		saveOrUpdate(code);
		
		SysUser user = (SysUser) get(SysUser.class, code.getRefId());
		if (user == null) {
			return null;
		}

		newToken = UUID.randomUUID().toString();
		user.setToken(newToken);
		user.setPassword(password);

		user.setLastLoginTime(new Date());
		saveOrUpdate(user);
		
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
	public boolean changePasswordPers(Long userId, String oldPassword, String password) {
		SysUser po = (SysUser) get(SysUser.class, userId);
		if (po == null || !po.getPassword().equals(oldPassword)) {
			return false;
		}
		
		po.setPassword(password);
		saveOrUpdate(po);
		return true;
	}
	
	@Override
	public SysUser saveProfile(UserVo vo) {
		SysUser po = (SysUser) get(SysUser.class, vo.getId());

		String name = vo.getName();
		String email = vo.getEmail(); 
		String phone = vo.getPhone(); 
		String avatar = vo.getAvatar();
		po.setPhone(phone);
		po.setName(name);
		po.setEmail(email);
		po.setAvatar(avatar);
		saveOrUpdate(po);
		return po;
	}

}
