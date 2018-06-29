package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.service.SysPrivilegeService;
import org.springframework.stereotype.Service;

import java.util.Map;

@Service
public class SysPrivilegeServiceImpl extends BaseServiceImpl implements SysPrivilegeService {

    @Override
    public Map<String, Boolean> listByUser(Long userId) {

//        String hql = "select priv from SysPrivilege priv" +
//                " join priv.sysRoleSet roles " +
//                " join roles.userSet users " +
//                " where users.id = ?" +
//                " and priv.deleted != true and priv.disabled!= true " +
//                " order by priv.id asc";
//
//        List<SysPrivilege> ls = getDao().getListByHQL(hql, userId);
//
//        Map<String, Boolean> map = new HashMap();
//        for (SysPrivilege po: ls) {
//            map.put(po.getCode().toString(), true);
//        }
//
//		return map;

        return null;
	}

}
