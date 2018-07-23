package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.OrgRoleDao;
import com.ngtesting.platform.model.TstOrgPrivilegeDefine;
import com.ngtesting.platform.model.TstOrgRole;
import com.ngtesting.platform.service.OrgRoleService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class OrgRoleServiceImpl extends BaseServiceImpl implements OrgRoleService {
	@Autowired
    OrgRoleDao orgRoleDao;

	@Override
	public List<TstOrgRole> list(Integer orgId, String keywords, String disabled) {
        List<TstOrgRole> ls = orgRoleDao.query(orgId, keywords, disabled);

		return ls;
	}

    @Override
    public TstOrgRole get(Integer orgRoleId) {
        return orgRoleDao.get(orgRoleId);
    }

    @Override
	public TstOrgRole save(TstOrgRole vo, Integer orgId) {
//		if (vo == null) {
//			return null;
//		}
//
//		TestOrgRole po = new TestOrgRole();
//		if (vo.getId() != null) {
//			po = (TestOrgRole) get(TestOrgRole.class, vo.getId());
//		}
//
//		po.setName(vo.getName());
//		po.setDescr(vo.getDescr());
//		po.setOrgId(orgId);
//		po.setDisabled(vo.getDisabled());
//
//		saveOrUpdate(po);
//		return po;

		return null;
	}

	@Override
	public boolean delete(Integer id) {
//		TestOrgRole po = (TestOrgRole) get(TestOrgRole.class, id);
//		po.setDeleted(true);
//		saveOrUpdate(po);

		return true;
	}

//	@Override
//	public void initOrgRoleBasicDataPers(Integer orgId) {
//		for (TestOrgRole.OrgRoleCode e : TestOrgRole.OrgRoleCode.values()) {
//            TestOrgRole po = new TestOrgRole();
//
//            po.setName(e.name);
//            po.setCode(e);
//            po.setDescr("");
//            po.setOrgId(orgId);
//            po.getOrgPrivilegeSet().addAll(getDefaultPrivByRoleCode(e));
//
//            saveOrUpdate(po);
//		}
//	}

    @Override
    public List<TstOrgPrivilegeDefine> getDefaultPrivByRoleCode(TstOrgRole.OrgRoleCode e) {
//        TestOrgPrivilegeDefine.OrgPrivilegeCode code = TestOrgPrivilegeDefine.OrgPrivilegeCode.valueOf(e.code);
//        DetachedCriteria dc = DetachedCriteria.forClass(TestOrgPrivilegeDefine.class);
//        dc.add(Restrictions.eq("code", code));
//        dc.add(Restrictions.ne("deleted", true));
//        dc.add(Restrictions.ne("disabled", true));
//
//        dc.addOrder(Order.asc("id"));
//        List<TestOrgPrivilegeDefine> ls = findAllByCriteria(dc);
//
//	    return ls;

		return null;
    }

//    @Override
//	public void addUserToOrgRolePers(TestUser user, Integer orgId, TestOrgRole.OrgRoleCode code) {
//		DetachedCriteria dc = DetachedCriteria.forClass(TestOrgRole.class);
//		dc.add(Restrictions.eq("orgId", orgId));
//		dc.add(Restrictions.eq("code", code));
//		dc.add(Restrictions.ne("deleted", true));
//		dc.add(Restrictions.ne("disabled", true));
//
//		dc.addOrder(Order.asc("id"));
//		List<TestOrgRole> ls = findAllByCriteria(dc);
//		TestOrgRole role = ls.get(0);
//		role.getUserSet().add(user);
//		saveOrUpdate(role);
//	}

}
