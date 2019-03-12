package com.ngtesting.platform.shiro;

import com.ngtesting.platform.utils.PasswordEncoder;
import com.ngtesting.platform.utils.StringUtil;
import org.apache.shiro.authc.*;
import org.apache.shiro.authc.credential.SimpleCredentialsMatcher;

public class CredentialsMatcher extends SimpleCredentialsMatcher {

	@Override
	public boolean doCredentialsMatch(AuthenticationToken token, AuthenticationInfo info) {
		UsernamePasswordToken utoken = (UsernamePasswordToken) token;

		SimpleAuthenticationInfo simpleAuthenticationInfo = (SimpleAuthenticationInfo) info;

		// 获得用户输入的密码
		String inPassword = new String(utoken.getPassword());
        if (StringUtil.isEmpty(inPassword)) { // loginWithVerifyCode的时候password为空，其他时候为空不会到到这里
            return true;
        }

		// 获得数据库中的密码
		String dbPassword = (String) info.getCredentials();

		// 进行密码的比对
		PasswordEncoder passwordEncoder = new PasswordEncoder(simpleAuthenticationInfo.getCredentialsSalt());
		Boolean pass = passwordEncoder.checkPassword(dbPassword, inPassword);

		return pass;
	}
}
