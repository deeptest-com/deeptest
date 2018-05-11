package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.ngtesting.platform.entity.TestOrgRole;
import com.ngtesting.platform.entity.TestUser;
import com.ngtesting.platform.service.OrgRoleUserService;
import com.ngtesting.platform.service.UserService;
import com.ngtesting.platform.vo.UserVo;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;
import java.util.Set;

@Service
public class OrgRoleUserServiceImpl extends BaseServiceImpl implements OrgRoleUserService {
	@Autowired
	UserService userService;

	@Override
	public List<UserVo> listUserByOrgRole(Long orgId, Long orgRoleId) {

        List<TestUser> allUsers = userService.listAllOrgUsers(orgId);

        List<TestUser> orgRoleUsers;
        if (orgRoleId == null) {
        	orgRoleUsers = new LinkedList<>();
        } else {
        	orgRoleUsers = listOrgRoleUsers(orgRoleId);
        }

        List<UserVo> vos = new LinkedList<UserVo>();
        for (TestUser po1 : allUsers) {
        	UserVo vo = genVo(po1);

        	vo.setSelected(false);
        	vo.setSelecting(false);
        	for (TestUser po2 : orgRoleUsers) {
        		if (po1.getId().longValue() == po2.getId().longValue()) {
            		vo.setSelected(true);
            		vo.setSelecting(true);
            	}
        	}
        	vos.add(vo);
        }

		return vos;
	}

	private UserVo genVo(TestUser po1) {
		UserVo vo = new UserVo(po1.getId(), po1.getName());

		return vo;
	}

	private List<TestUser> listOrgRoleUsers(Long orgRoleId) {

		DetachedCriteria dc = DetachedCriteria.forClass(TestUser.class);

        dc.createAlias("orgRoleSet", "roles");
        dc.add(Restrictions.eq("roles.id", orgRoleId));

        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));

        dc.addOrder(Order.asc("id"));
        List ls = findAllByCriteria(dc);

		return ls;
	}

	@Override
	public boolean saveOrgRoleUsers(Long roleId, List<UserVo> orgUsers) {
		if (orgUsers == null) {
			return false;
		}

		TestOrgRole orgRole = (TestOrgRole) get(TestOrgRole.class, roleId);
		Set<TestUser> userSet = orgRole.getUserSet();

		for (Object obj: orgUsers) {
			UserVo vo = JSON.parseObject(JSON.toJSONString(obj), UserVo.class);
			if (vo.getSelecting() != vo.getSelected()) { // 变化了
				TestUser orgUser = (TestUser) get(TestUser.class, vo.getId());

    			if (vo.getSelecting() && !userSet.contains(orgUser)) { // 勾选
					userSet.add(orgUser);
    			} else if (orgUser != null) { // 取消
					userSet.remove(orgUser);
    			}
			}
		}
		saveOrUpdate(orgRole);

		return true;
	}

}
