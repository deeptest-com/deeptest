package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.SysPrivilegeDao;
import com.ngtesting.platform.model.SysPrivilege;
import com.ngtesting.platform.service.intf.SysPrivilegeService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

@Service
public class SysPrivilegeServiceImpl extends BaseServiceImpl implements SysPrivilegeService {

    @Autowired
    SysPrivilegeDao sysPrivilegeDao;

    @Override
    public Map<String, Boolean> listByUser(Integer userId) {

        List<SysPrivilege> privs = sysPrivilegeDao.queryByUser(userId);

        Map<String, Boolean> map = new HashMap();
        for (SysPrivilege po: privs) {
            map.put(po.getCode().toString(), true);
        }

		return map;
	}

}
