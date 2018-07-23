package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.OrgRoleUserDao;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.OrgRoleUserService;
import com.ngtesting.platform.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;

@Service
public class OrgRoleUserServiceImpl extends BaseServiceImpl implements OrgRoleUserService {
	@Autowired
	UserService userService;
	@Autowired
	OrgRoleUserDao orgRoleUserDao;

	@Override
	public List<TstUser> listUserByOrgRole(Integer orgId, Integer orgRoleId) {

        List<TstUser> allUsers = userService.listAllOrgUsers(orgId);

        List<TstUser> orgRoleUsers;
        if (orgRoleId == null) {
        	orgRoleUsers = new LinkedList<>();
        } else {
        	orgRoleUsers = listOrgRoleUsers(orgRoleId);
        }

        List<TstUser> vos = new LinkedList<TstUser>();
        for (TstUser po1 : allUsers) {
        	TstUser vo = genVo(po1);

        	vo.setSelected(false);
        	vo.setSelecting(false);
        	for (TstUser po2 : orgRoleUsers) {
        		if (po1.getId().longValue() == po2.getId().longValue()) {
            		vo.setSelected(true);
            		vo.setSelecting(true);
            	}
        	}
        	vos.add(vo);
        }

		return vos;
	}

	private TstUser genVo(TstUser po1) {
//		TstUser vo = new TstUser(po1.getId(), po1.getName());
//
//		return vo;

		return null;
	}

	private List<TstUser> listOrgRoleUsers(Integer orgRoleId) {

//		DetachedCriteria dc = DetachedCriteria.forClass(TstUser.class);
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

	@Override
	public boolean saveOrgRoleUsers(Integer roleId, List<TstUser> orgUsers) {
//		if (orgUsers == null) {
//			return false;
//		}
//
//		TestOrgRole orgRole = (TestOrgRole) get(TestOrgRole.class, roleId);
//		Set<TstUser> userSet = orgRole.getUserSet();
//
//		for (Object obj: orgUsers) {
//			TstUser vo = JSON.parseObject(JSON.toJSONString(obj), TstUser.class);
//			if (vo.getSelecting() != vo.getSelected()) { // 变化了
//				TstUser orgUser = (TstUser) get(TstUser.class, vo.getId());
//
//    			if (vo.getSelecting() && !userSet.contains(orgUser)) { // 勾选
//					userSet.add(orgUser);
//    			} else if (orgUser != null) { // 取消
//					userSet.remove(orgUser);
//    			}
//			}
//		}
//		saveOrUpdate(orgRole);

		return true;
	}

}
