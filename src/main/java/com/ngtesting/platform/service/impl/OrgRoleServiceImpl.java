package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.OrgRoleDao;
import com.ngtesting.platform.model.TstOrgRole;
import com.ngtesting.platform.service.intf.OrgRoleService;
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
    public TstOrgRole get(Integer orgRoleId, Integer orgId) {
        return orgRoleDao.get(orgRoleId, orgId);
    }

    @Override
	public TstOrgRole save(TstOrgRole vo, Integer orgId) {
        vo.setOrgId(orgId);

        if (vo.getId() == null) {
            vo.setCode(UUID.randomUUID().toString());
            orgRoleDao.save(vo);
        } else {
            Integer count = orgRoleDao.update(vo);
            if (count == 0) {
                return null;
            }
        }

        return vo;
	}

	@Override
	public boolean delete(Integer id, Integer orgId) {
        Integer count = orgRoleDao.delete(id, orgId);

		return count > 0;
	}

}
