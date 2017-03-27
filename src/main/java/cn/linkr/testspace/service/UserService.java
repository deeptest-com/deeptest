package cn.linkr.testspace.service;

import java.util.List;
import java.util.Map;

import cn.linkr.testspace.entity.SysUser;
import cn.linkr.testspace.entity.SysVerifyCode;
import cn.linkr.testspace.vo.Page;
import cn.linkr.testspace.vo.UserVo;

public interface UserService extends BaseService {

	Page listByPage(Long companyId, String keywords, String disabled, Integer currentPage, Integer itemsPerPage);
	
	SysUser save(UserVo vo);
	boolean remove(Long id);
	boolean disablePers(Long id);

	List<UserVo> genVos(List<SysUser> pos);
	UserVo genVo(SysUser user);

}
