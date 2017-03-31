package com.ngtesting.platform.service;

import java.util.List;
import java.util.Set;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.SysGroup;
import com.ngtesting.platform.entity.SysGroupUser;
import com.ngtesting.platform.vo.GroupVo;
import com.ngtesting.platform.vo.Page;

public interface GroupService extends BaseService {

	Page listByPage(Long companyId, String keywords, String disabled, Integer currentPage, Integer itemsPerPage);
	
	SysGroup save(GroupVo vo, Long companyId);
	boolean delete(Long id);
	boolean disable(Long id);
	
	List<GroupVo> listByUser(Long companyId, Long userId);
	boolean saveGroupsByUser(JSONObject to, Long companyId, Long userId);

	List<GroupVo> genVos(List<SysGroup> pos);
	GroupVo genVo(SysGroup user);

	List<SysGroupUser> listUserGroups(Long companyId, Long userId);

	SysGroupUser getGroupUser(Long companyId, Long userId, Long groupId);

}
