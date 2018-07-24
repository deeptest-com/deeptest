package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.OrgRolePrivilegeRelationDao;
import com.ngtesting.platform.model.TstOrgPrivilegeDefine;
import com.ngtesting.platform.service.OrgPrivilegeService;
import com.ngtesting.platform.service.OrgRoleService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Map;

@Service
public class OrgPrivilegeServiceImpl extends BaseServiceImpl implements OrgPrivilegeService {
	@Autowired
	private OrgRolePrivilegeRelationDao orgRolePrivilegeRelationDao;

    @Autowired
    OrgRoleService orgRoleService;

    @Override
    public Map<String, Boolean> listByUser(Integer userId, Integer orgId) {
//        String hql = "select role from TestOrgRole role" +
//                " join role.userSet users " +
//                " where users.id = ?" +
//                " and role.orgId = ?" +
//
//                " and role.deleted != true and role.disabled!= true " +
//                " order by role.id asc";
//
//        List<TestOrgRole> ls = getDao().getListByHQL(hql, userId, orgId);
//
//        Map<String, Boolean> map = new HashMap();
//        for (TestOrgRole role: ls) {
//            for (TestOrgPrivilegeDefine priv: role.getOrgPrivilegeSet()) {
//                map.put(priv.getCode().toString(), true);
//            }
//        }
//
//        return map;

        return null;
    }

    @Override
    public List<TstOrgPrivilegeDefine> listAllOrgPrivileges(Integer orgId) {
        return null;
    }

}
