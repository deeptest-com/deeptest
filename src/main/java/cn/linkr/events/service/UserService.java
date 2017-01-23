package cn.linkr.events.service;

import java.util.List;
import java.util.Map;

import cn.linkr.events.entity.SysUser;
import cn.linkr.events.entity.SysVerifyCode;
import cn.linkr.events.vo.Page;
import cn.linkr.events.vo.UserVo;

public interface UserService extends BaseService {

	SysUser getByToken(String token);

	SysUser getByPhone(String token);

	SysUser loginPers(String mobile, String password, Boolean rememberMe);
	SysUser logoutPers(String email);

	SysUser registerPers(String name, String email, String phone, String password);

	SysVerifyCode forgetPaswordPers(String mobile);

	SysUser resetPasswordPers(String verifyCode, String mobile,
			String password, String platform, String isWebView,
			String deviceToken);


	Page listByPage(long companyId, int currentPage, int itemsPerPage);
	SysUser save(UserVo vo);
	boolean remove(Long id);
	boolean disable(Long id);

	List<UserVo> genVos(List<SysUser> pos);
	UserVo genVo(SysUser user);

	SysUser saveProfile(UserVo vo);

}
