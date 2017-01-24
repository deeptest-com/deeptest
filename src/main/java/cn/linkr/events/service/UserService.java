package cn.linkr.events.service;

import java.util.List;
import java.util.Map;

import cn.linkr.events.entity.SysUser;
import cn.linkr.events.entity.SysVerifyCode;
import cn.linkr.events.vo.Page;
import cn.linkr.events.vo.UserVo;

public interface UserService extends BaseService {


	SysUser loginPers(String mobile, String password, Boolean rememberMe);
	SysUser logoutPers(String email);

	SysUser registerPers(String name, String email, String phone, String password);
	
	boolean changePasswordPers(Long userId, String oldPassword, String password);

	SysVerifyCode forgotPasswordPers(Long userId);
	SysUser resetPasswordPers(String verifyCode, String password);

	Page listByPage(long companyId, int currentPage, int itemsPerPage);
	SysUser saveProfile(UserVo vo);
	SysUser save(UserVo vo);
	boolean remove(Long id);
	boolean disablePers(Long id);

	SysUser getByToken(String token);
	SysUser getByPhone(String token);
	SysUser getByEmail(String email);

	List<UserVo> genVos(List<SysUser> pos);
	UserVo genVo(SysUser user);

}
