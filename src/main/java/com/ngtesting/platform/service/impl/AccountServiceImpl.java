package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.entity.TestUser;
import com.ngtesting.platform.entity.TestVerifyCode;
import com.ngtesting.platform.service.AccountService;
import com.ngtesting.platform.vo.UserVo;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import java.util.Date;
import java.util.List;
import java.util.UUID;

@Service
public class AccountServiceImpl extends BaseServiceImpl implements AccountService {

	@Override
	public TestUser loginPers(String email, String password, Boolean rememberMe) {
		String newToken = null;
		DetachedCriteria dc = DetachedCriteria.forClass(TestUser.class);
		dc.add(Restrictions.eq("email", email));
		dc.add(Restrictions.eq("password", password));
		dc.add(Restrictions.ne("deleted", true));
		List<TestUser> ls = (List<TestUser>) findAllByCriteria(dc);

		TestUser user = null;
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
	public TestUser registerPers(String name, String email, String phone, String password) {
		String newToken = null;
		DetachedCriteria dc = DetachedCriteria.forClass(TestUser.class);
		
		dc.add(Restrictions.eq("email", email));
		dc.add(Restrictions.ne("deleted", true));
		dc.add(Restrictions.ne("disabled", true));
		List<TestUser> ls = (List<TestUser>) findAllByCriteria(dc);

		if (ls.size() > 0) {
			return null;
		}

		TestUser user = new TestUser();
		newToken = UUID.randomUUID().toString();
		user.setName(name);
		user.setToken(newToken);
		user.setEmail(email);
		user.setPhone(phone);
		user.setPassword(password);
		user.setAvatar("upload/sample/user/avatar.png");
		
		user.setLastLoginTime(new Date());
		saveOrUpdate(user);

		return user;
	}

	@Override
	public TestVerifyCode forgotPasswordPers(Long userId) {
		TestUser user = (TestUser) get(TestUser.class, userId);
		if (user == null) {
			return null;
		}

		TestVerifyCode po = new TestVerifyCode();
		String code = UUID.randomUUID().toString().replaceAll("-", "");
		Date now = new Date();
		po.setRefId(user.getId());
		po.setCode(code);
		po.setCreateTime(now);
		po.setExpireTime(new Date(now.getTime() + 10 * 60 * 1000));
		saveOrUpdate(po);

		return po;
	}

	@Override
	public TestUser getByPhone(String phone) {
		DetachedCriteria dc = DetachedCriteria.forClass(TestUser.class);
		dc.add(Restrictions.eq("phone", phone));
		dc.add(Restrictions.ne("deleted", true));
		dc.add(Restrictions.ne("disabled", true));

		List ls = findAllByCriteria(dc);
		if (ls.size() > 0) {
			return (TestUser) ls.get(0);
		} else {
			return null;
		}
	}
	@Override
	public TestUser getByEmail(String email) {
		DetachedCriteria dc = DetachedCriteria.forClass(TestUser.class);
		dc.add(Restrictions.eq("email", email));
		dc.add(Restrictions.ne("deleted", true));
		dc.add(Restrictions.ne("disabled", true));

		List ls = findAllByCriteria(dc);
		if (ls.size() > 0) {
			return (TestUser) ls.get(0);
		} else {
			return null;
		}
	}

	@Override
	public TestUser resetPasswordPers(String verifyCode, String password) {

		String newToken = null;
		DetachedCriteria dc = DetachedCriteria.forClass(TestVerifyCode.class);
		dc.add(Restrictions.eq("code", verifyCode));
		dc.add(Restrictions.ge("expireTime", new Date()));
		dc.add(Restrictions.ne("deleted", true));
		dc.add(Restrictions.ne("disabled", true));
		dc.addOrder(Order.desc("id"));
		List<TestVerifyCode> ls = (List<TestVerifyCode>) findAllByCriteria(dc);

		if (ls.size() < 1) {
			return null;
		}
		
		TestVerifyCode code = ls.get(0);
		code.setDeleted(true);
		saveOrUpdate(code);
		
		TestUser user = (TestUser) get(TestUser.class, code.getRefId());
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
	public TestUser logoutPers(String email) {
		
		DetachedCriteria dc = DetachedCriteria.forClass(TestUser.class);
		dc.add(Restrictions.eq("email", email));
		dc.add(Restrictions.ne("deleted", true));
		dc.add(Restrictions.ne("disabled", true));
		List<TestUser> ls = (List<TestUser>) findAllByCriteria(dc);

		TestUser user = null;
		if (ls.size() > 0) {
			user = ls.get(0);
			user.setToken("");
			saveOrUpdate(user);
		}
		return user;
	}
	

	@Override
	public TestUser getByToken(String token) {
		DetachedCriteria dc = DetachedCriteria.forClass(TestUser.class);
		dc.add(Restrictions.eq("token", token));
		
		dc.add(Restrictions.ne("deleted", true));
		dc.add(Restrictions.ne("disabled", true));

		List ls = findAllByCriteria(dc);
		if (ls.size() > 0) {
			return (TestUser) ls.get(0);
		} else {
			return null;
		}
	}

	@Override
	public boolean changePasswordPers(Long userId, String oldPassword, String password) {
		TestUser po = (TestUser) get(TestUser.class, userId);
		if (po == null || !po.getPassword().equals(oldPassword)) {
			return false;
		}
		
		po.setPassword(password);
		saveOrUpdate(po);
		return true;
	}
	
	@Override
	public TestUser saveProfile(UserVo vo) {
		TestUser po = (TestUser) get(TestUser.class, vo.getId());

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
