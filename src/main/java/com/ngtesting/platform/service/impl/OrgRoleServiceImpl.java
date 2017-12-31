package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.entity.TestOrgPrivilege;
import com.ngtesting.platform.entity.TestOrgRole;
import com.ngtesting.platform.entity.TestUser;
import com.ngtesting.platform.service.OrgRoleService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.util.StringUtil;
import com.ngtesting.platform.vo.OrgRoleVo;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;

@Service
public class OrgRoleServiceImpl extends BaseServiceImpl implements OrgRoleService {

	@Override
	public List list(Long orgId, String keywords, String disabled) {
        DetachedCriteria dc = DetachedCriteria.forClass(TestOrgRole.class);
        dc.add(Restrictions.eq("orgId", orgId));

        dc.add(Restrictions.eq("deleted", Boolean.FALSE));

        if (StringUtil.isNotEmpty(keywords)) {
        	dc.add(Restrictions.like("name", "%" + keywords + "%"));
        }
        if (StringUtil.isNotEmpty(disabled)) {
        	dc.add(Restrictions.eq("disabled", Boolean.valueOf(disabled)));
        }

        dc.addOrder(Order.asc("id"));
        List ls = findAllByCriteria(dc);

		return ls;
	}

	@Override
	public TestOrgRole save(OrgRoleVo vo, Long orgId) {
		if (vo == null) {
			return null;
		}

		TestOrgRole po = new TestOrgRole();
		if (vo.getId() != null) {
			po = (TestOrgRole) get(TestOrgRole.class, vo.getId());
		}

		po.setName(vo.getName());
		po.setDescr(vo.getDescr());
		po.setOrgId(orgId);
		po.setDisabled(vo.getDisabled());

		saveOrUpdate(po);
		return po;
	}

	@Override
	public boolean delete(Long id) {
		TestOrgRole po = (TestOrgRole) get(TestOrgRole.class, id);
		po.setDeleted(true);
		saveOrUpdate(po);

		return true;
	}

//	@Override
//	public void initOrgRoleBasicDataPers(Long orgId) {
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
    public List<TestOrgPrivilege> getDefaultPrivByRoleCode(TestOrgRole.OrgRoleCode e) {
        TestOrgPrivilege.OrgPrivilegeCode code = TestOrgPrivilege.OrgPrivilegeCode.valueOf(e.code);
        DetachedCriteria dc = DetachedCriteria.forClass(TestOrgPrivilege.class);
        dc.add(Restrictions.eq("code", code));
        dc.add(Restrictions.ne("deleted", true));
        dc.add(Restrictions.ne("disabled", true));

        dc.addOrder(Order.asc("id"));
        List<TestOrgPrivilege> ls = findAllByCriteria(dc);

	    return ls;
    }

//    @Override
//	public void addUserToOrgRolePers(TestUser user, Long orgId, TestOrgRole.OrgRoleCode code) {
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

	@Override
	public OrgRoleVo genVo(TestOrgRole role) {
		OrgRoleVo vo = new OrgRoleVo();
		BeanUtilEx.copyProperties(vo, role);

		return vo;
	}
	@Override
	public List<OrgRoleVo> genVos(List<TestOrgRole> pos) {
        List<OrgRoleVo> vos = new LinkedList<OrgRoleVo>();

        for (TestOrgRole po: pos) {
        	OrgRoleVo vo = genVo(po);
        	vos.add(vo);
        }
		return vos;
	}
}
