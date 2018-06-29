package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.model.TstOrgPrivilege;
import com.ngtesting.platform.service.OrgRolePrivilegeService;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;
import java.util.Map;

@Service
public class OrgRolePrivilegeServiceImpl extends BaseServiceImpl implements OrgRolePrivilegeService {

	@Override
	public List<TstOrgPrivilege> listPrivilegesByOrgRole(Integer orgId, Integer orgRoleId) {

        List<TstOrgPrivilege> allPrivileges = listAllOrgPrivileges();

        List<TstOrgPrivilege> orgRolePrivileges;
        if (orgRoleId == null) {
        	orgRolePrivileges = new LinkedList<>();
        } else {
        	orgRolePrivileges = listOrgRolePrivileges(orgId, orgRoleId);
        }

        List<TstOrgPrivilege> vos = new LinkedList<TstOrgPrivilege>();
        for (TstOrgPrivilege po1 : allPrivileges) {
        	TstOrgPrivilege vo = genVo(orgId, po1);

        	vo.setSelected(false);
        	vo.setSelecting(false);
        	for (TstOrgPrivilege po2 : orgRolePrivileges) {
        		if (po1.getId().longValue() == po2.getId().longValue()) {
            		vo.setSelected(true);
            		vo.setSelecting(true);
            	}
        	}
        	vos.add(vo);
        }

		return vos;
	}

	private TstOrgPrivilege genVo(Integer orgId, TstOrgPrivilege po1) {
		TstOrgPrivilege vo = new TstOrgPrivilege(po1.getId(), po1.getName(), po1.getDescr(), orgId);

		return vo;
	}

	private List<TstOrgPrivilege> listOrgRolePrivileges(Integer orgId, Integer orgRoleId) {

//		DetachedCriteria dc = DetachedCriteria.forClass(TstOrgPrivilege.class);
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

	private List<TstOrgPrivilege> listAllOrgPrivileges() {
//		DetachedCriteria dc = DetachedCriteria.forClass(TstOrgPrivilege.class);
//
//        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//        dc.addOrder(Order.asc("id"));
//        List<TstOrgPrivilege> ls = findAllByCriteria(dc);
//
//		return ls;

		return null;
	}

	@Override
	public boolean saveOrgRolePrivileges(Integer roleId, List<TstOrgPrivilege> orgPrivileges) {
//		if (orgPrivileges == null) {
//			return false;
//		}
//
//		TestOrgRole orgRole = (TestOrgRole) get(TestOrgRole.class, roleId);
//		Set<TstOrgPrivilege> privilegeSet = orgRole.getOrgPrivilegeSet();
//
//		for (Object obj: orgPrivileges) {
//			TstOrgPrivilege vo = JSON.parseObject(JSON.toJSONString(obj), TstOrgPrivilege.class);
//			if (vo.getSelecting() != vo.getSelected()) { // 变化了
//				TstOrgPrivilege orgPrivilege = (TstOrgPrivilege) get(TstOrgPrivilege.class, vo.getId());
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
//
//		List<TestOrgRole> ls = getDao().getListByHQL(hql, userId, orgId);
//
//		Map<String, Boolean> map = new HashMap();
//		for (TestOrgRole role: ls) {
//            for (TstOrgPrivilege priv: role.getOrgPrivilegeSet()) {
//                map.put(priv.getCode().toString(), true);
//            }
//		}
//
//		return map;

		return null;
	}

}
