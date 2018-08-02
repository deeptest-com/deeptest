package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.OrgRoleDao;
import com.ngtesting.platform.model.TstOrgPrivilegeDefine;
import com.ngtesting.platform.model.TstOrgRole;
import com.ngtesting.platform.service.OrgRoleService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.UUID;

@Service
public class OrgRoleServiceImpl extends BaseServiceImpl implements OrgRoleService {
	@Autowired
    OrgRoleDao orgRoleDao;

	@Override
	public List<TstOrgRole> list(Integer orgId, String keywords, Boolean disabled) {
        List<TstOrgRole> ls = orgRoleDao.query(orgId, keywords, disabled);

		return ls;
	}

    @Override
    public List<TstOrgRole> listAllOrgRoles(Integer orgId) {
        List<TstOrgRole> ls = orgRoleDao.query(orgId, null, false);

        return ls;
    }

    @Override
    public TstOrgRole get(Integer orgRoleId) {
        return orgRoleDao.get(orgRoleId);
    }

    @Override
	public TstOrgRole save(TstOrgRole vo, Integer orgId) {
        vo.setOrgId(orgId);

        if (vo.getId() == null) {
            vo.setCode(UUID.randomUUID().toString());
            orgRoleDao.save(vo);
        } else {
            orgRoleDao.update(vo);
        }

        return vo;
	}

	@Override
	public boolean delete(Integer id) {
//		TestOrgRole po = (TestOrgRole) getDetail(TestOrgRole.class, id);
//		po.setDeleted(true);
//		saveOrUpdate(po);

		return true;
	}

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

}
