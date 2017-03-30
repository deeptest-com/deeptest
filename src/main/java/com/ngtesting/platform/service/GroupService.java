package com.ngtesting.platform.service;

import java.util.List;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.SysGroup;
import com.ngtesting.platform.vo.GroupVo;
import com.ngtesting.platform.vo.Page;

public interface GroupService extends BaseService {

	Page listByPage(Long companyId, String keywords, String disabled, Integer currentPage, Integer itemsPerPage);
	
	SysGroup save(GroupVo vo, Long companyId);
	boolean delete(Long id);
	boolean disable(Long id);
	
	List<SysGroup> listByUser(Long companyId, Long userId);
	boolean saveGroupsByUser(JSONObject to);

	List<GroupVo> genVos(List<SysGroup> pos);
	GroupVo genVo(SysGroup user);

}
