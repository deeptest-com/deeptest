package cn.linkr.events.service.impl;

import java.util.Date;
import java.util.List;
import java.util.UUID;

import org.apache.commons.lang.StringUtils;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import cn.linkr.events.entity.SysUser;
import cn.linkr.events.entity.SysUser.AgentType;
import cn.linkr.events.entity.SysVerifyCode;
import cn.linkr.events.service.UserService;
import cn.linkr.events.util.StringUtil;

@Service
public class UserServiceImpl extends BaseServiceImpl implements UserService {

	@Override
	public SysUser getUserByToken(String token) {
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
	public SysUser loginPers(String phone, String password, String platform, String isWebview, String deviceToken) {
		String newToken = null;
		DetachedCriteria dc = DetachedCriteria.forClass(SysUser.class);
		dc.add(Restrictions.eq("phone", phone));
		dc.add(Restrictions.eq("password", password));
		dc.add(Restrictions.ne("deleted", true));
		dc.add(Restrictions.ne("disabled", true));
		List<SysUser> ls = (List<SysUser>) findAllByCriteria(dc);

		SysUser user = null;
		if (ls.size() > 0) {
			user = ls.get(0);
			newToken = UUID.randomUUID().toString();
			user.setToken(newToken);

			if (StringUtils.isNotEmpty(platform)) {
				user.setPlatform(SysUser.PlatformType.StringToEnum(platform.trim().toUpperCase()));
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
		}
		return user;
	}

	@Override
	public SysUser registerPers(String phone, String password, String platform, String isWebview, String deviceToken) {
		String newToken = null;
		DetachedCriteria dc = DetachedCriteria.forClass(SysUser.class);
		dc.add(Restrictions.eq("phone", phone));
		dc.add(Restrictions.ne("deleted", true));
		dc.add(Restrictions.ne("disabled", true));
		List<SysUser> ls = (List<SysUser>) findAllByCriteria(dc);

		if (ls.size() > 0) {
			return null;
		}

		SysUser user = new SysUser();
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

		return user;
	}

	@Override
	public SysVerifyCode forgetPaswordPers(String phone) {
		SysUser user = getUserByPhone(phone);
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
	public SysUser getUserByPhone(String phone) {
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

		SysUser user = getUserByPhone(phone);
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
    
}
