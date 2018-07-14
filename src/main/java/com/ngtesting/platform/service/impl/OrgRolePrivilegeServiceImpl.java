package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.OrgRolePrivilegeDao;
import com.ngtesting.platform.model.TstOrgPrivilegeDefine;
import com.ngtesting.platform.service.OrgRolePrivilegeService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.HashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;

@Service
public class OrgRolePrivilegeServiceImpl extends BaseServiceImpl implements OrgRolePrivilegeService {

	@Autowired
	private OrgRolePrivilegeDao orgRolePrivilegeDao;

	@Override
	public List<TstOrgPrivilegeDefine> listPrivilegesByOrgRole(Integer orgId, Integer orgRoleId) {

        List<TstOrgPrivilegeDefine> allPrivileges = listAllOrgPrivileges();

        List<TstOrgPrivilegeDefine> orgRolePrivileges;
        if (orgRoleId == null) {
        	orgRolePrivileges = new LinkedList<>();
        } else {
        	orgRolePrivileges = listOrgRolePrivileges(orgId, orgRoleId);
        }

        List<TstOrgPrivilegeDefine> vos = new LinkedList<TstOrgPrivilegeDefine>();
        for (TstOrgPrivilegeDefine po1 : allPrivileges) {

			po1.setSelected(false);
			po1.setSelecting(false);
        	for (TstOrgPrivilegeDefine po2 : orgRolePrivileges) {
        		if (po1.getId().longValue() == po2.getId().longValue()) {
					po1.setSelected(true);
					po1.setSelecting(true);
            	}
        	}
        	vos.add(po1);
        }

		return vos;
	}

	private List<TstOrgPrivilegeDefine> listOrgRolePrivileges(Integer orgId, Integer orgRoleId) {

//		DetachedCriteria dc = DetachedCriteria.forClass(TstOrgPrivilegeDefine.class);
//
//        dc.createAlias("orgRoleSet", "roles");
//        dc.add(Restrictions.eq("roles.id", orgRoleId));
//
//        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//
//        dc.addOrder(Order.asc("id"));
//        List ls = findAllByCriteria(dc);
//
//		return ls;

		return null;
	}

	private List<TstOrgPrivilegeDefine> listAllOrgPrivileges() {
//		DetachedCriteria dc = DetachedCriteria.forClass(TstOrgPrivilegeDefine.class);
//
//        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//        dc.addOrder(Order.asc("id"));
//        List<TstOrgPrivilegeDefine> ls = findAllByCriteria(dc);
//
//		return ls;

		return null;
	}

	@Override
	public boolean saveOrgRolePrivileges(Integer roleId, List<TstOrgPrivilegeDefine> orgPrivileges) {
//		if (orgPrivileges == null) {
//			return false;
//		}
//
//		TestOrgRole orgRole = (TestOrgRole) get(TestOrgRole.class, roleId);
//		Set<TstOrgPrivilegeDefine> privilegeSet = orgRole.getOrgPrivilegeSet();
//
//		for (Object obj: orgPrivileges) {
//			TstOrgPrivilegeDefine vo = JSON.parseObject(JSON.toJSONString(obj), TstOrgPrivilegeDefine.class);
//			if (vo.getSelecting() != vo.getSelected()) { // 变化了
//				TstOrgPrivilegeDefine orgPrivilege = (TstOrgPrivilegeDefine) get(TstOrgPrivilegeDefine.class, vo.getId());
//
//    			if (vo.getSelecting() && !privilegeSet.contains(orgPrivilege)) { // 勾选
//    				privilegeSet.add(orgPrivilege);
//    			} else if (orgPrivilege != null) { // 取消
//    				privilegeSet.remove(orgPrivilege);
//    			}
//			}
//		}
//		saveOrUpdate(orgRole);

		return true;
	}

	@Override
	public Map<String, Boolean> listByUser(Integer userId, Integer orgId) {
//	    String hql = "select role from TestOrgRole role" +
//                " join role.userSet users " +
//                " where users.id = ?" +
//                " and role.orgId = ?" +
//
//                " and role.deleted != true and role.disabled!= true " +
//                " order by role.id asc";

		List<TstOrgPrivilegeDefine> ls = orgRolePrivilegeDao.queryByUser(userId, orgId);

		Map<String, Boolean> map = new HashMap();
		for (TstOrgPrivilegeDefine priv: ls) {
			map.put(priv.getCode().toString(), true);
		}

		return map;
	}

}
